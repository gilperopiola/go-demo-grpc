## Installing stuff

**protoc:** Part of the protocol-buffers repo from Google, on Windows you can download the compiled binary from https://github.com/protocolbuffers/protobuf/releases

-> On MacOS, instead you should just

 > brew install protobuf

---

**buf:** CLI tool that makes everything easier, install from https://buf.build/docs/cli/installation/

---

> go get google.golang.org/grpc
> go get buf.build/go/protovalidate

---

**protoc-gen-go:** Autogenerates **.pb.go** files based on proto messages

 > go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

**protoc-gen-go-grpc:** Autogenerates **_grpc.pb.go** files based on proto services

 > go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

---
To run the generation, just do
 > buf generate