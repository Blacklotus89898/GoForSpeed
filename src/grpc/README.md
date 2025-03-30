# gRPC 

## To generate from proto file

```bash
# Install go dependencies
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate go files from .proto files
protoc --go_out=. --go-grpc_out=. greeter.proto    

# Dependencies
go mod tidy

# Running
go run server.go
go run client.go
```