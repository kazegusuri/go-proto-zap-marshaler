# go-proto-zap-marshaler

A protoc plugin whitch generates `MarshalLogObject()` functions for each generated go structs to implement `zapcore.ObjectMarshaler` interface for [uber-go/zap](https://github.com/uber-go/zap). So you can use `zap.Object("proto", someProtoMessage)` to log contents of the proto message.

## Install and usage

First you need `protoc-gen-zap-marshaler` of protoc plugin. You can get it by doing:

```bash
$ go get github.com/kazegusuri/go-proto-zap-marshaler/protoc-gen-zap-marshaler
```

To generate marshaler functions from proto, use `--zap-marshaler_out` with `protoc` command. That runs `protoc-gen-zap-marshaler` internally and then generates `*.zap.go` files.

```
$ protoc --zap-marshaler_out=. path/to/example.proto
```
