REPOSITORY=github.com/kazegusuri/go-proto-zap-marshaler

regenerate:
	protoc -I. -I${GOPATH}/src --gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/gogoproto/descriptor.proto:. zap_marshaler.proto
	protoc -I. --go_out=. examples/*.proto
	protoc -I. --zap-marshaler_out=. examples/*.proto

install:
	go install $(REPOSITORY)/protoc-gen-zap-marshaler

test:
	go test -v $(REPOSITORY)/examples/...
