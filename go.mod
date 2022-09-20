module github.com/kazegusuri/go-proto-zap-marshaler

go 1.18

require (
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	go.uber.org/zap v1.23.0
	google.golang.org/protobuf v1.26.0
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)

// replace github.com/gogo/protobuf => ../protobuf
replace go.uber.org/zap v1.23.0 => github.com/WqyJh/zap v1.23.1-1
