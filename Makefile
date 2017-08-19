REPOSITORY=github.com/kazegusuri/go-proto-zap-marshaler

regenerate:
	protoc -I. --go_out=. examples/*.proto
	protoc -I. --zap-marshaler_out=. examples/*.proto

install:
	go install $(REPOSITORY)/protoc-gen-zap-marshaler

test:
	go test -v $(REPOSITORY)/examples
