package examples

import (
	"go.uber.org/zap"
)

func ExampleTest() {
	l := zap.NewExample()
	l.Info("test", zap.Object("simple", &SimpleMessage{
		StringValue: "xyz",
		BoolValue:   true,
	}))
	// output: {"level":"info","msg":"test","simple":{"stringValue":"xyz","boolValue":true}}
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
	// output: {"level":"info","msg":"test","number":{"floatValue":0.5,"doubleValue":2.2,"int32Value":-3,"int64Value":-4,"uint32Value":5,"uint64Value":6,"sint32Value":-7,"sint64Value":-8,"fixed32Value":9,"fixed64Value":10,"sfixed32Value":-11,"sfixed64Value":-12}}
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
	// output: {"level":"info","msg":"test","number":{"floatValues":[0.5,1],"doubleValues":[2.2,1],"int32Values":[-3,3],"int64Values":[-4,4],"uint32Values":[5,55],"uint64Values":[6,66],"sint32Values":[-7,7],"sint64Values":[-8,8],"fixed32Values":[9,99],"fixed64Values":[10,100],"sfixed32Values":[-11,11],"sfixed64Values":[-12,12]}}
}

func ExampleNestedMessage() {
	// TODO
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
	// output: {"level":"info","msg":"test","nested":{"nestedValue":{"int32Value":100,"stringValue":"xxx"}}}
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
	// output: {"level":"info","msg":"test","enum":{"numericEnumValue":"ONE","repeatedNumericEnumValues":["ONE","TWO"],"aliasedEnumValue":"STARTED","nestedEnumValue":"PENDING","repeatedNestedEnumValues":["PENDING","COMPLETED"]}}
}

func ExampleOneofMessage() {
	// TODO
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
	// output: {"level":"info","msg":"test","oneof":{}}
}

func ExampleMapMessage() {
	// TODO
	l := zap.NewExample()
	l.Info("test", zap.Object("map", &MapMessage{
		MappedValue: map[string]string{
			"foo": "foo1",
			"bar": "bar1",
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
	// output: {"level":"info","msg":"test","map":{}}
}
