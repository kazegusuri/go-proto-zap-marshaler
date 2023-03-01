REPOSITORY=github.com/kazegusuri/go-proto-zap-marshaler

regenerate:
	$(MAKE) regenerate/proto
	$(MAKE) regenerate/examples

regenerate/proto:
	rm -rf gen && mkdir -p gen
	protoc -I. --gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:gen zap_marshaler.proto
	mv gen/$(REPOSITORY)/zap_marshaler.pb.go ./

regenerate/examples:
	protoc -I. --go_out=. ./examples/*.proto
	protoc -I. --zap-marshaler_out=. ./examples/*.proto
	mv example.*.go ./examples/zap-marshaler

	protoc -I. --go_out=. ./examples/*.proto
	protoc -I. --zap-marshaler-secure_out=. ./examples/*.proto
	mv example.*.go ./examples/zap-marshaler-secure

install:
	go install ./protoc-gen-zap-marshaler
	go install ./protoc-gen-zap-marshaler-secure

test:
	go test -v ./examples/...
