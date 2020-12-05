package examples

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"go.uber.org/zap"
)

func ExampleTest() {
	l := zap.NewExample()
	l.Info("test", zap.Object("simple", &SimpleMessage{
		StringValue: "xyz",
		BoolValue:   true,
	}))
	// output: {"level":"info","msg":"test","simple":{"string_value":"xyz","bool_value":true}}
}

func ExampleNotLoggingTest() {
	l := zap.NewExample()
	l.Info("test", zap.Object("simple", &NotLoggingSimpleMessage{
		StringValue: "xyz",
		BoolValue:   true,
	}))
	// output: {"level":"info","msg":"test","simple":{"string_value":"xyz","bool_value":true}}
}

func ExampleNumberMessage() {
	l := zap.NewExample()
	l.Info("test", zap.Object("number", &NumberMessage{
		FloatValue:    0.5,
		DoubleValue:   2.2,
		Int32Value:    -3,
		Int64Value:    -4,
		Uint32Value:   5,
		Uint64Value:   6,
		Sint32Value:   -7,
		Sint64Value:   -8,
		Fixed32Value:  9,
		Fixed64Value:  10,
		Sfixed32Value: -11,
		Sfixed64Value: -12,
	}))
	// output: {"level":"info","msg":"test","number":{"float_value":0.5,"double_value":2.2,"int32_value":-3,"int64_value":-4,"uint32_value":5,"uint64_value":6,"sint32_value":-7,"sint64_value":-8,"fixed32_value":9,"fixed64_value":10,"sfixed32_value":-11,"sfixed64_value":-12}}
}

func ExampleNotLoggingNumberMessage() {
	l := zap.NewExample()
	l.Info("test", zap.Object("number", &NotLoggingNumberMessage{
		FloatValue:    0.5,
		DoubleValue:   2.2,
		Int32Value:    -3,
		Int64Value:    -4,
		Uint32Value:   5,
		Uint64Value:   6,
		Sint32Value:   -7,
		Sint64Value:   -8,
		Fixed32Value:  9,
		Fixed64Value:  10,
		Sfixed32Value: -11,
		Sfixed64Value: -12,
	}))
	// output: {"level":"info","msg":"test","number":{"float_value":0.5,"double_value":2.2,"int32_value":-3,"int64_value":-4,"uint32_value":5,"uint64_value":6,"sint32_value":-7,"sint64_value":-8,"fixed32_value":9,"fixed64_value":10,"sfixed32_value":-11,"sfixed64_value":-12}}
}

func ExampleRepeatedNumberMessage() {
	l := zap.NewExample()
	l.Info("test", zap.Object("number", &RepeatedNumberMessage{
		FloatValues:    []float32{0.5, 1},
		DoubleValues:   []float64{2.2, 1},
		Int32Values:    []int32{-3, 3},
		Int64Values:    []int64{-4, 4},
		Uint32Values:   []uint32{5, 55},
		Uint64Values:   []uint64{6, 66},
		Sint32Values:   []int32{-7, 7},
		Sint64Values:   []int64{-8, 8},
		Fixed32Values:  []uint32{9, 99},
		Fixed64Values:  []uint64{10, 100},
		Sfixed32Values: []int32{-11, 11},
		Sfixed64Values: []int64{-12, 12},
	}))
	// output: {"level":"info","msg":"test","number":{"float_values":[0.5,1],"double_values":[2.2,1],"int32_values":[-3,3],"int64_values":[-4,4],"uint32_values":[5,55],"uint64_values":[6,66],"sint32_values":[-7,7],"sint64_values":[-8,8],"fixed32_values":[9,99],"fixed64_values":[10,100],"sfixed32_values":[-11,11],"sfixed64_values":[-12,12]}}
}

func ExampleNotLoggingRepeatedNumberMessage() {
	l := zap.NewExample()
	l.Info("test", zap.Object("number", &NotLoggingRepeatedNumberMessage{
		FloatValues:    []float32{0.5, 1},
		DoubleValues:   []float64{2.2, 1},
		Int32Values:    []int32{-3, 3},
		Int64Values:    []int64{-4, 4},
		Uint32Values:   []uint32{5, 55},
		Uint64Values:   []uint64{6, 66},
		Sint32Values:   []int32{-7, 7},
		Sint64Values:   []int64{-8, 8},
		Fixed32Values:  []uint32{9, 99},
		Fixed64Values:  []uint64{10, 100},
		Sfixed32Values: []int32{-11, 11},
		Sfixed64Values: []int64{-12, 12},
	}))
	// output: {"level":"info","msg":"test","number":{"float_values":[0.5,1],"double_values":[2.2,1],"int32_values":[-3,3],"int64_values":[-4,4],"uint32_values":[5,55],"uint64_values":[6,66],"sint32_values":[-7,7],"sint64_values":[-8,8],"fixed32_values":[9,99],"fixed64_values":[10,100],"sfixed32_values":[-11,11],"sfixed64_values":[-12,12]}}
}

func ExampleNestedMessage() {
	l := zap.NewExample()
	l.Info("test", zap.Object("nested", &NestedMessage{
		NestedValue: &NestedMessage_Nested{
			Int32Value:  100,
			StringValue: "xxx",
		},
		RepeatedNestedValues: []*NestedMessage_Nested{
			{
				Int32Value:  200,
				StringValue: "yyy",
			},
			{
				Int32Value:  300,
				StringValue: "zzz",
			},
		},
	}))
	// output: {"level":"info","msg":"test","nested":{"nested_value":{"int32_value":100,"string_value":"xxx"},"repeated_nested_values":[{"int32_value":200,"string_value":"yyy"},{"int32_value":300,"string_value":"zzz"}]}}
}

func ExampleNotLoggingNestedMessage() {
	l := zap.NewExample()
	l.Info("test", zap.Object("nested", &NotLoggingNestedMessage{
		NestedValue: &NotLoggingNestedMessage_Nested{
			Int32Value:  100,
			StringValue: "xxx",
		},
		RepeatedNestedValues: []*NotLoggingNestedMessage_Nested{
			{
				Int32Value:  200,
				StringValue: "yyy",
			},
			{
				Int32Value:  300,
				StringValue: "zzz",
			},
		},
	}))
	// output: {"level":"info","msg":"test","nested":{"nested_value":{"int32_value":100,"string_value":"xxx"},"repeated_nested_values":[{"int32_value":200,"string_value":"yyy"},{"int32_value":300,"string_value":"zzz"}]}}
}

func ExampleEnumMessage() {
	l := zap.NewExample()
	l.Info("test", zap.Object("enum", &EnumMessage{
		NumericEnumValue: NumericEnum_ONE,
		RepeatedNumericEnumValues: []NumericEnum{
			NumericEnum_ONE,
			NumericEnum_TWO,
		},
		AliasedEnumValue: AliasedEnum_RUNNING,
		NestedEnumValue:  EnumMessage_PENDING,
		RepeatedNestedEnumValues: []EnumMessage_Nested{
			EnumMessage_PENDING,
			EnumMessage_COMPLETED,
		},
	}))
	// output: {"level":"info","msg":"test","enum":{"numeric_enum_value":"ONE","repeated_numeric_enum_values":["ONE","TWO"],"aliased_enum_value":"STARTED","nested_enum_value":"PENDING","repeated_nested_enum_values":["PENDING","COMPLETED"]}}
}

func ExampleNotLoggingEnumMessage() {
	l := zap.NewExample()
	l.Info("test", zap.Object("enum", &NotLoggingEnumMessage{
		NumericEnumValue: NumericEnum_ONE,
		RepeatedNumericEnumValues: []NumericEnum{
			NumericEnum_ONE,
			NumericEnum_TWO,
		},
		AliasedEnumValue: AliasedEnum_RUNNING,
		NestedEnumValue:  NotLoggingEnumMessage_PENDING,
		RepeatedNestedEnumValues: []NotLoggingEnumMessage_Nested{
			NotLoggingEnumMessage_PENDING,
			NotLoggingEnumMessage_COMPLETED,
		},
	}))
	// output: {"level":"info","msg":"test","enum":{"numeric_enum_value":"ONE","repeated_numeric_enum_values":["ONE","TWO"],"aliased_enum_value":"STARTED","nested_enum_value":"PENDING","repeated_nested_enum_values":["PENDING","COMPLETED"]}}
}

func ExampleOneofMessage() {
	l := zap.NewExample()
	l.Info("test", zap.Object("oneof", &OneofMessage{
		OneofValue: &OneofMessage_Int32Value{Int32Value: 1000},
		RepeatedOneofValues: []*Oneof{
			{
				OneofValue: &Oneof_Int32Value{Int32Value: 1000},
			},
			{
				OneofValue: &Oneof_StringValue{StringValue: "xyz"},
			},
		},
	}))
	// output: {"level":"info","msg":"test","oneof":{"int32_value":1000,"repeated_oneof_values":[{"int32_value":1000},{"string_value":"xyz"}]}}
}

func ExampleNotLoggingOneofMessage() {
	l := zap.NewExample()
	l.Info("test", zap.Object("oneof", &NotLoggingOneofMessage{
		OneofValue: &NotLoggingOneofMessage_Int32Value{Int32Value: 1000},
		RepeatedOneofValues: []*Oneof{
			{
				OneofValue: &Oneof_Int32Value{Int32Value: 1000},
			},
			{
				OneofValue: &Oneof_StringValue{StringValue: "xyz"},
			},
		},
	}))
	// output: {"level":"info","msg":"test","oneof":{"int32_value":1000,"repeated_oneof_values":[{"int32_value":1000},{"string_value":"xyz"}]}}
}

func ExampleMapMessage() {
	l := zap.NewExample()
	l.Info("test", zap.Object("map", &MapMessage{
		MappedValue: map[int32]string{
			1: "foo",
			2: "bar",
		},
		MappedEnumValue: map[string]NumericEnum{
			"one": NumericEnum_ONE,
			"two": NumericEnum_TWO,
		},
		MappedNestedValue: map[string]*NestedMessage{
			"foo": &NestedMessage{
				NestedValue: &NestedMessage_Nested{
					Int32Value:  100,
					StringValue: "xxx",
				},
				RepeatedNestedValues: []*NestedMessage_Nested{
					{
						Int32Value:  200,
						StringValue: "yyy",
					},
					{
						Int32Value:  300,
						StringValue: "zzz",
					},
				},
			},
		},
	}))
	// output: {"level":"info","msg":"test","map":{"mapped_value":{"1":"foo","2":"bar"},"mapped_enum_value":{"one":"ONE","two":"TWO"},"mapped_nested_value":{"foo":{"nested_value":{"int32_value":100,"string_value":"xxx"},"repeated_nested_values":[{"int32_value":200,"string_value":"yyy"},{"int32_value":300,"string_value":"zzz"}]}}}}
}

func ExampleNotLoggingMapMessage() {
	l := zap.NewExample()
	l.Info("test", zap.Object("map", &NotLoggingMapMessage{
		MappedValue: map[int32]string{
			1: "foo",
			2: "bar",
		},
		MappedEnumValue: map[string]NumericEnum{
			"one": NumericEnum_ONE,
			"two": NumericEnum_TWO,
		},
		MappedNestedValue: map[string]*NestedMessage{
			"foo": &NestedMessage{
				NestedValue: &NestedMessage_Nested{
					Int32Value:  100,
					StringValue: "xxx",
				},
				RepeatedNestedValues: []*NestedMessage_Nested{
					{
						Int32Value:  200,
						StringValue: "yyy",
					},
					{
						Int32Value:  300,
						StringValue: "zzz",
					},
				},
			},
		},
	}))
	// output: {"level":"info","msg":"test","map":{"mapped_value":{"1":"foo","2":"bar"},"mapped_enum_value":{"one":"ONE","two":"TWO"},"mapped_nested_value":{"foo":{"nested_value":{"int32_value":100,"string_value":"xxx"},"repeated_nested_values":[{"int32_value":200,"string_value":"yyy"},{"int32_value":300,"string_value":"zzz"}]}}}}
}

func ExampleJsonNameOptionMessage() {
	l := zap.NewExample()
	l.Info("test", zap.Object("json_name", &JsonNameOptionMessage{
		StringValue: "xxx",
	}))
	// output: {"level":"info","msg":"test","json_name":{"string_value":"xxx"}}
}

func ExampleNotLoggingJsonNameOptionMessage() {
	l := zap.NewExample()
	l.Info("test", zap.Object("json_name", &NotLoggingJsonNameOptionMessage{
		StringValue: "xxx",
	}))
	// output: {"level":"info","msg":"test","json_name":{"string_value":"xxx"}}
}

func ExampleWellKnownTypeMessage() {
	d := ptypes.DurationProto(10 * time.Minute)
	t, _ := ptypes.TimestampProto(time.Unix(1502533013, 125892275))
	l := zap.NewExample()
	l.Info("test", zap.Object("wkt", &WellKnownTypeMessage{
		Duration:  d,
		Timestamp: t,
	}))
	// output: {"level":"info","msg":"test","wkt":{"duration":"10m0s","timestamp":"2017-08-12T10:16:53.125Z"}}
}

func ExampleNotLoggingWellKnownTypeMessage() {
	d := ptypes.DurationProto(10 * time.Minute)
	t, _ := ptypes.TimestampProto(time.Unix(1502533013, 125892275))
	l := zap.NewExample()
	l.Info("test", zap.Object("wkt", &NotLoggingWellKnownTypeMessage{
		Duration:  d,
		Timestamp: t,
	}))
	// output: {"level":"info","msg":"test","wkt":{"duration":"10m0s","timestamp":"2017-08-12T10:16:53.125Z"}}
}

func ExampleMixedLoggingMessage() {
	l := zap.NewExample()
	l.Info("test", zap.Object("mlm", &MixedLoggingMessage{
		StringValue: "xxx",
		BoolValue:   true,
		Int32Value:  1,
	}))
	// {"level":"info","msg":"test","mlm":{"string_value":"xxx","bool_value":true,"int32_value":1}}
}

func ExampleTestWithNilField() {
	l := zap.NewExample()
	sl := l.Sugar()
	var sm *SimpleMessage = nil
	sl = sl.With("SimpleMessage", sm)
	sl.Infow("test")
	// output: {"level":"info","msg":"test","SimpleMessage":{}}
}
