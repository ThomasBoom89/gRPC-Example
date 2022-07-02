# gRPC example

![Go](https://img.shields.io/github/go-mod/go-version/thomasboom89/gRPC-Example/main)
![License](https://img.shields.io/badge/license-MIT-green?style=plastic)

## Requirements
```zsh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
Add GOPATH to $PATH!


## command to generate message/client/server stubs

```bash
protoc example/*.proto \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    --proto_path=.
```