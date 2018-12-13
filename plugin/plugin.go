package plugin

import (
	"fmt"

	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/kazegusuri/go-proto-zap-marshaler"
)

type plugin struct {
	*generator.Generator
	generator.PluginImports
	zapcore generator.Single
	ptypes  generator.Single
}

func NewPlugin() generator.Plugin {
	return &plugin{}
}

func (p *plugin) Name() string {
	return "zap-marshaler"
}

func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.zapcore = p.NewImport("go.uber.org/zap/zapcore")
	p.ptypes = p.NewImport("github.com/golang/protobuf/ptypes")

	for _, msg := range file.Messages() {
		if msg.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}

		if gogoproto.IsProto3(file.FileDescriptorProto) {
			p.generateProto3Message(file, msg)
		}
	}
}

func isTarget(d *descriptor.FieldDescriptorProto) bool {
	if d.GetOptions() == nil {
		return false
	}
	ext, err := proto.GetExtension(d.Options, zap_marshaler.E_Field)
	if err != nil {
		return false
	}

	rule, ok := ext.(*zap_marshaler.ZapMarshalerRule)
	return ok && rule.Enabled
}

func (p *plugin) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {
	typeName := generator.CamelCaseSlice(message.TypeName())
	p.P(`func (m *`, typeName, `) MarshalLogObject(enc `, p.zapcore.Use(), `.ObjectEncoder) error {`)
	p.In()
	p.P("var keyName string")
	p.P("_ = keyName")
	p.P("")

	for _, field := range message.Field {
		if !isTarget(field) {
			continue
		}
		fieldName := p.GetOneOfFieldName(message, field)
		jsonName := field.GetName()
		variableName := "m." + fieldName
		p.P(`keyName = "`, jsonName, `" // field `, field.GetName(), " = ", fmt.Sprint(field.GetNumber()))

		repeated := field.IsRepeated() && !p.IsMap(field)
		if repeated {
			p.P(`enc.AddArray(keyName, `, p.zapcore.Use(), `.ArrayMarshalerFunc(func(aenc `, p.zapcore.Use(), `.ArrayEncoder) error {`)
			p.In()

			p.P(`for _, rv := range `, variableName, `{`)
			variableName = "rv"
			p.In()
			p.P(`_ = `, variableName) // suppress unused error
		}

		if p.isOneofType(field) {
			oneofName := p.GetFieldName(message, field)
			oneofType := p.OneOfTypeName(message, field)
			p.P(`if ov, ok := m`, `.Get`, oneofName+`().(* `+oneofType+`); ok {`)
			variableName = "ov." + fieldName
			p.In()
			p.P(` _ = ov`)
		}

		p.generateForField(file, message, field, "keyName", variableName, repeated)

		if p.isOneofType(field) {
			p.Out()
			p.P(`}`)
		}

		if repeated {
			p.Out()
			p.P(`}`)

			p.P(`return nil`)
			p.Out()
			p.P(`}))`)
		}

		p.P()
	}
	p.P(`return nil`)
	p.Out()
	p.P(`}`)
	p.P()
}

func (p *plugin) isOneofType(field *descriptor.FieldDescriptorProto) bool {
	return field.OneofIndex != nil
}

func (p *plugin) generateForField(file *generator.FileDescriptor, message *generator.Descriptor, field *descriptor.FieldDescriptorProto, keyName, variableName string, repeated bool) {
	// TODO: support well known type to log pretty message
	// TODO: marshal unknown type

	switch field.GetTypeName() {
	case ".google.protobuf.Duration":
		p.P(`if d, err := `, p.ptypes.Use(), ".Duration(", variableName, "); err == nil {")
		p.In()
		p.generateAdder("Duration", keyName, "d", repeated)
		p.Out()
		p.P(`}`)
		return

	case ".google.protobuf.Timestamp":
		p.P(`if t, err := `, p.ptypes.Use(), ".Timestamp(", variableName, "); err == nil {")
		p.In()
		p.generateAdder("Time", keyName, "t", repeated)
		p.Out()
		p.P(`}`)
		return
	}

	if p.IsMap(field) {
		mapDesc := p.GoMapType(nil, field)
		if mapDesc == nil {
			p.P("// unavaiable map type")
			return
		}
		keyField := mapDesc.KeyField
		valField := mapDesc.ValueField

		// sanity check to avoid unexpected loop
		if p.IsMap(valField) {
			p.P("// unavaible map type: nested map")
			return
		}

		p.P(`enc.AddObject(keyName, `, p.zapcore.Use(), `.ObjectMarshalerFunc(func(enc `, p.zapcore.Use(), `.ObjectEncoder) error {`)
		p.In()

		p.P("for mk, mv := range ", variableName, " {")
		p.In()
		if keyField.IsString() {
			p.P("key := mk")
		} else {
			p.P("key := fmt.Sprint(mk)")
		}
		p.P(" _ = key")
		p.generateForField(file, message, valField, "key", "mv", false)

		p.Out()
		p.P("}")

		p.P(`return nil`)
		p.Out()
		p.P(`}))`)

		return
	}

	switch *(field.Type) {
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		p.generateAdder("Int32", keyName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		p.generateAdder("Int64", keyName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		p.generateAdder("Uint32", keyName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		p.generateAdder("Uint64", keyName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_SINT32:
		p.generateAdder("Int32", keyName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_SINT64:
		p.generateAdder("Int64", keyName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_FIXED32:
		p.generateAdder("Uint32", keyName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_FIXED64:
		p.generateAdder("Uint64", keyName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		p.generateAdder("Int32", keyName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		p.generateAdder("Int64", keyName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		p.generateAdder("Float32", keyName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		p.generateAdder("Float64", keyName, variableName, repeated)

	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		p.generateAdder("Bool", keyName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		p.generateAdder("String", keyName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		p.generateAdder("ByteString", keyName, variableName, repeated)

	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		if repeated {
			p.P(`aenc.AppendString(`, variableName, `.String())`)
		} else {
			p.P(`enc.AddString(`, keyName, `, `, variableName, `.String())`)
		}

	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		p.P(`if `, variableName, ` != nil {`)
		p.In()

		p.P(`var vv interface{} = `, variableName)
		p.P(`if marshaler, ok := `, `vv.(`, p.zapcore.Use(), `.ObjectMarshaler); ok {`)
		p.In()
		p.generateAdder("Object", keyName, "marshaler", repeated)
		p.Out()
		p.P(`}`)

		p.Out()
		p.P(`}`)

	case descriptor.FieldDescriptorProto_TYPE_GROUP:
		// ?
		p.P("// group type not supported")
	}
}

func (p *plugin) generateAdder(ftype, keyName, variable string, repeated bool) {
	if repeated {
		p.P("aenc.Append", ftype, "(", variable, ")")
	} else {
		p.P("enc.Add", ftype, `(`, keyName, `, `, variable, `)`)
	}
}
