package plugin

import (
	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type plugin struct {
	*generator.Generator
	generator.PluginImports
	zapcore generator.Single
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

	for _, msg := range file.Messages() {
		if msg.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}

		if gogoproto.IsProto3(file.FileDescriptorProto) {
			p.generateProto3Message(file, msg)
		}
	}
}

func (p *plugin) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {
	typeName := generator.CamelCaseSlice(message.TypeName())
	p.P(`func (m *`, typeName, `) MarshalLogObject(enc `, p.zapcore.Use(), `.ObjectEncoder) error {`)
	p.In()
	for _, field := range message.Field {
		fieldName := p.GetOneOfFieldName(message, field)
		jsonName := field.GetJsonName()
		variableName := "m." + fieldName

		repeated := field.IsRepeated() && !p.isMapType(field)
		if repeated {
			p.P("// repeated")
			p.P(`enc.AddArray("`, jsonName, `", `, p.zapcore.Use(), `.ArrayMarshalerFunc(func(aenc `, p.zapcore.Use(), `.ArrayEncoder) error {`)
			p.In()

			p.P(`for _, v := range `, variableName, `{`)
			p.In()
			p.P(`_ = v`) // suppress unused error
		}

		p.generateForField(file, message, field, repeated)

		if repeated {
			p.Out()
			p.P(`}`)

			p.P(`return nil`)
			p.Out()
			p.P(`}))`)
		}
	}
	p.P(`return nil`)
	p.Out()
	p.P(`}`)
	p.P()
}

func (p *plugin) isMapType(field *descriptor.FieldDescriptorProto) bool {
	if field.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE && field.IsRepeated() {
		return true
	}
	return false
}

func (p *plugin) isOneofType(field *descriptor.FieldDescriptorProto) bool {
	return field.OneofIndex != nil
}

func (p *plugin) generateForField(file *generator.FileDescriptor, message *generator.Descriptor, field *descriptor.FieldDescriptorProto, repeated bool) {
	fieldName := p.GetOneOfFieldName(message, field)
	jsonName := field.GetJsonName()
	variableName := "m." + fieldName

	if p.isMapType(field) {
		// TODO: support map
		return
	}
	if p.isOneofType(field) {
		// TODO: support oneof
		return
	}
	// TODO: support well known type

	switch *(field.Type) {
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		p.generateAdder("Int32", jsonName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		p.generateAdder("Int64", jsonName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		p.generateAdder("Uint32", jsonName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		p.generateAdder("Uint64", jsonName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_SINT32:
		p.generateAdder("Int32", jsonName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_SINT64:
		p.generateAdder("Int64", jsonName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_FIXED32:
		p.generateAdder("Uint32", jsonName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_FIXED64:
		p.generateAdder("Uint64", jsonName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		p.generateAdder("Int32", jsonName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		p.generateAdder("Int64", jsonName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		p.generateAdder("Float32", jsonName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		p.generateAdder("Float64", jsonName, variableName, repeated)

	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		p.generateAdder("Bool", jsonName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		p.generateAdder("String", jsonName, variableName, repeated)
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		p.generateAdder("ByteString", jsonName, variableName, repeated)

	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		if repeated {
			p.P(`aenc.AppendString(v.String())`)
		} else {
			p.P(`enc.AddString("`, jsonName, `", `, variableName, `.String())`)
		}

	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		p.P(`if `, variableName, ` != nil {`)
		p.In()

		p.P(`var vv interface{} = `, variableName)
		p.P(`if _, ok := `, `vv.(`, p.zapcore.Use(), `.ObjectMarshaler); ok {`)
		p.In()
		p.generateAdder("Object", jsonName, variableName, repeated)
		p.Out()
		p.P(`}`)

		p.Out()
		p.P(`}`)

	case descriptor.FieldDescriptorProto_TYPE_GROUP:
		// ?
	}
}

func (p *plugin) generateAdder(ftype, name, variable string, repeated bool) {
	if repeated {
		p.P("aenc.Append", ftype, "(v)")
	} else {
		p.P("enc.Add", ftype, `("`, name, `", `, variable, `)`)
	}
}
