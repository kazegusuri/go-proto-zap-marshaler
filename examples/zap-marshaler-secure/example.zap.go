// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/example.proto

package examples

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/kazegusuri/go-proto-zap-marshaler"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	go_uber_org_zap_zapcore "go.uber.org/zap/zapcore"
	github_com_golang_protobuf_ptypes "github.com/golang/protobuf/ptypes"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (m *SimpleMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "string_value" // field string_value = 1
	enc.AddString(keyName, m.StringValue)

	keyName = "bool_value" // field bool_value = 2
	enc.AddBool(keyName, m.BoolValue)

	return nil
}

func (m *NotLoggingSimpleMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	// disabled field string_value = 1

	// disabled field bool_value = 2

	return nil
}

func (m *NumberMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "float_value" // field float_value = 1
	enc.AddFloat32(keyName, m.FloatValue)

	keyName = "double_value" // field double_value = 2
	enc.AddFloat64(keyName, m.DoubleValue)

	keyName = "int32_value" // field int32_value = 3
	enc.AddInt32(keyName, m.Int32Value)

	keyName = "int64_value" // field int64_value = 4
	enc.AddInt64(keyName, m.Int64Value)

	keyName = "uint32_value" // field uint32_value = 5
	enc.AddUint32(keyName, m.Uint32Value)

	keyName = "uint64_value" // field uint64_value = 6
	enc.AddUint64(keyName, m.Uint64Value)

	keyName = "sint32_value" // field sint32_value = 7
	enc.AddInt32(keyName, m.Sint32Value)

	keyName = "sint64_value" // field sint64_value = 8
	enc.AddInt64(keyName, m.Sint64Value)

	keyName = "fixed32_value" // field fixed32_value = 9
	enc.AddUint32(keyName, m.Fixed32Value)

	keyName = "fixed64_value" // field fixed64_value = 10
	enc.AddUint64(keyName, m.Fixed64Value)

	keyName = "sfixed32_value" // field sfixed32_value = 11
	enc.AddInt32(keyName, m.Sfixed32Value)

	keyName = "sfixed64_value" // field sfixed64_value = 12
	enc.AddInt64(keyName, m.Sfixed64Value)

	return nil
}

func (m *NotLoggingNumberMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	// disabled field float_value = 1

	// disabled field double_value = 2

	// disabled field int32_value = 3

	// disabled field int64_value = 4

	// disabled field uint32_value = 5

	// disabled field uint64_value = 6

	// disabled field sint32_value = 7

	// disabled field sint64_value = 8

	// disabled field fixed32_value = 9

	// disabled field fixed64_value = 10

	// disabled field sfixed32_value = 11

	// disabled field sfixed64_value = 12

	return nil
}

func (m *RepeatedNumberMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "float_values" // field float_values = 1
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.FloatValues {
			_ = rv
			aenc.AppendFloat32(rv)
		}
		return nil
	}))

	keyName = "double_values" // field double_values = 2
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.DoubleValues {
			_ = rv
			aenc.AppendFloat64(rv)
		}
		return nil
	}))

	keyName = "int32_values" // field int32_values = 3
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Int32Values {
			_ = rv
			aenc.AppendInt32(rv)
		}
		return nil
	}))

	keyName = "int64_values" // field int64_values = 4
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Int64Values {
			_ = rv
			aenc.AppendInt64(rv)
		}
		return nil
	}))

	keyName = "uint32_values" // field uint32_values = 5
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Uint32Values {
			_ = rv
			aenc.AppendUint32(rv)
		}
		return nil
	}))

	keyName = "uint64_values" // field uint64_values = 6
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Uint64Values {
			_ = rv
			aenc.AppendUint64(rv)
		}
		return nil
	}))

	keyName = "sint32_values" // field sint32_values = 7
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Sint32Values {
			_ = rv
			aenc.AppendInt32(rv)
		}
		return nil
	}))

	keyName = "sint64_values" // field sint64_values = 8
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Sint64Values {
			_ = rv
			aenc.AppendInt64(rv)
		}
		return nil
	}))

	keyName = "fixed32_values" // field fixed32_values = 9
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Fixed32Values {
			_ = rv
			aenc.AppendUint32(rv)
		}
		return nil
	}))

	keyName = "fixed64_values" // field fixed64_values = 10
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Fixed64Values {
			_ = rv
			aenc.AppendUint64(rv)
		}
		return nil
	}))

	keyName = "sfixed32_values" // field sfixed32_values = 11
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Sfixed32Values {
			_ = rv
			aenc.AppendInt32(rv)
		}
		return nil
	}))

	keyName = "sfixed64_values" // field sfixed64_values = 12
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Sfixed64Values {
			_ = rv
			aenc.AppendInt64(rv)
		}
		return nil
	}))

	return nil
}

func (m *NotLoggingRepeatedNumberMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	// disabled field float_values = 1

	// disabled field double_values = 2

	// disabled field int32_values = 3

	// disabled field int64_values = 4

	// disabled field uint32_values = 5

	// disabled field uint64_values = 6

	// disabled field sint32_values = 7

	// disabled field sint64_values = 8

	// disabled field fixed32_values = 9

	// disabled field fixed64_values = 10

	// disabled field sfixed32_values = 11

	// disabled field sfixed64_values = 12

	return nil
}

func (m *NestedMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "nested_value" // field nested_value = 1
	if m.NestedValue != nil {
		var vv interface{} = m.NestedValue
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	keyName = "repeated_nested_values" // field repeated_nested_values = 2
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.RepeatedNestedValues {
			_ = rv
			if rv != nil {
				var vv interface{} = rv
				if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
					aenc.AppendObject(marshaler)
				}
			}
		}
		return nil
	}))

	return nil
}

func (m *NestedMessage_Nested) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "int32_value" // field int32_value = 1
	enc.AddInt32(keyName, m.Int32Value)

	keyName = "string_value" // field string_value = 2
	enc.AddString(keyName, m.StringValue)

	return nil
}

func (m *NotLoggingNestedMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	// disabled field nested_value = 1

	// disabled field repeated_nested_values = 2

	return nil
}

func (m *NotLoggingNestedMessage_Nested) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	// disabled field int32_value = 1

	// disabled field string_value = 2

	return nil
}

func (m *EnumMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "numeric_enum_value" // field numeric_enum_value = 1
	enc.AddString(keyName, m.NumericEnumValue.String())

	keyName = "repeated_numeric_enum_values" // field repeated_numeric_enum_values = 2
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.RepeatedNumericEnumValues {
			_ = rv
			aenc.AppendString(rv.String())
		}
		return nil
	}))

	keyName = "aliased_enum_value" // field aliased_enum_value = 3
	enc.AddString(keyName, m.AliasedEnumValue.String())

	keyName = "nested_enum_value" // field nested_enum_value = 4
	enc.AddString(keyName, m.NestedEnumValue.String())

	keyName = "repeated_nested_enum_values" // field repeated_nested_enum_values = 5
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.RepeatedNestedEnumValues {
			_ = rv
			aenc.AppendString(rv.String())
		}
		return nil
	}))

	return nil
}

func (m *NotLoggingEnumMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	// disabled field numeric_enum_value = 1

	// disabled field repeated_numeric_enum_values = 2

	// disabled field aliased_enum_value = 3

	// disabled field nested_enum_value = 4

	// disabled field repeated_nested_enum_values = 5

	return nil
}

func (m *Oneof) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "int32_value" // field int32_value = 1
	if ov, ok := m.GetOneofValue().(*Oneof_Int32Value); ok {
		_ = ov
		enc.AddInt32(keyName, ov.Int32Value)
	}

	keyName = "string_value" // field string_value = 2
	if ov, ok := m.GetOneofValue().(*Oneof_StringValue); ok {
		_ = ov
		enc.AddString(keyName, ov.StringValue)
	}

	return nil
}

func (m *OneofMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "int32_value" // field int32_value = 1
	if ov, ok := m.GetOneofValue().(*OneofMessage_Int32Value); ok {
		_ = ov
		enc.AddInt32(keyName, ov.Int32Value)
	}

	keyName = "string_value" // field string_value = 2
	if ov, ok := m.GetOneofValue().(*OneofMessage_StringValue); ok {
		_ = ov
		enc.AddString(keyName, ov.StringValue)
	}

	keyName = "repeated_oneof_values" // field repeated_oneof_values = 3
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.RepeatedOneofValues {
			_ = rv
			if rv != nil {
				var vv interface{} = rv
				if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
					aenc.AppendObject(marshaler)
				}
			}
		}
		return nil
	}))

	return nil
}

func (m *NotLoggingOneofMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	// disabled field int32_value = 1

	// disabled field string_value = 2

	// disabled field repeated_oneof_values = 3

	return nil
}

func (m *MapMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "mapped_value" // field mapped_value = 1
	enc.AddObject(keyName, go_uber_org_zap_zapcore.ObjectMarshalerFunc(func(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
		for mk, mv := range m.MappedValue {
			key := fmt.Sprint(mk)
			_ = key
			enc.AddString(key, mv)
		}
		return nil
	}))

	keyName = "mapped_enum_value" // field mapped_enum_value = 2
	enc.AddObject(keyName, go_uber_org_zap_zapcore.ObjectMarshalerFunc(func(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
		for mk, mv := range m.MappedEnumValue {
			key := mk
			_ = key
			enc.AddString(key, mv.String())
		}
		return nil
	}))

	keyName = "mapped_nested_value" // field mapped_nested_value = 3
	enc.AddObject(keyName, go_uber_org_zap_zapcore.ObjectMarshalerFunc(func(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
		for mk, mv := range m.MappedNestedValue {
			key := mk
			_ = key
			if mv != nil {
				var vv interface{} = mv
				if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
					enc.AddObject(key, marshaler)
				}
			}
		}
		return nil
	}))

	return nil
}

func (m *NotLoggingMapMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	// disabled field mapped_value = 1

	// disabled field mapped_enum_value = 2

	// disabled field mapped_nested_value = 3

	return nil
}

func (m *JsonNameOptionMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "string_value" // field string_value = 1
	enc.AddString(keyName, m.StringValue)

	return nil
}

func (m *NotLoggingJsonNameOptionMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	// disabled field string_value = 1

	return nil
}

func (m *WellKnownTypeMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "duration" // field duration = 1
	if d, err := github_com_golang_protobuf_ptypes.Duration(m.Duration); err == nil {
		enc.AddDuration(keyName, d)
	}

	keyName = "timestamp" // field timestamp = 2
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.Timestamp); err == nil {
		enc.AddTime(keyName, t)
	}

	return nil
}

func (m *NotLoggingWellKnownTypeMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	// disabled field duration = 1

	// disabled field timestamp = 2

	return nil
}

func (m *MixedLoggingMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "string_value" // field string_value = 1
	enc.AddString(keyName, m.StringValue)

	// disabled field bool_value = 2

	// disabled field int32_value = 3

	return nil
}

func (m *SkipFieldMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	// disabled field a = 1

	// disabled field b = 2

	return nil
}
