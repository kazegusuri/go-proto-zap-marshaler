regenerate:
	protoc -I. --go_out=. examples/*.proto
	protoc -I. --zap-marshaler_out=. examples/*.proto

install:
	go install github.com/kazegusuri/go-protoc-zap-marshaler/protoc-gen-zap-marshaler

test:
	go test -v github.com/kazegusuri/go-protoc-zap-marshaler/examples
