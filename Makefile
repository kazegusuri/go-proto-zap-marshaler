REPOSITORY=github.com/kazegusuri/go-proto-zap-marshaler

regenerate:
	$(MAKE) regenerate/proto
	$(MAKE) regenerate/examples

regenerate/proto:
	rm -rf gen && mkdir -p gen
	# protoc -I. -I${GOPATH}/src --gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/gogoproto/descriptor.proto:gen zap_marshaler.proto
	protoc -I. -I${GOPATH}/src --gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:gen zap_marshaler.proto
	mv gen/$(REPOSITORY)/zap_marshaler.pb.go ./

regenerate/examples:
	cp examples/*.proto examples/zap-marshaler/
	protoc -I. --go_out=. examples/zap-marshaler/*.proto
	protoc -I. --zap-marshaler_out=. examples/zap-marshaler/*.proto

	cp examples/*.proto examples/zap-marshaler-secure/
	protoc -I. --go_out=. examples/zap-marshaler-secure/*.proto
	protoc -I. --zap-marshaler-secure_out=. examples/zap-marshaler-secure/*.proto

install:
	go install $(REPOSITORY)/protoc-gen-zap-marshaler
	go install $(REPOSITORY)/protoc-gen-zap-marshaler-secure

test:
	go test -v $(REPOSITORY)/examples/...
