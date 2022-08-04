### Getting started with Go development
  - Initing go module:
  `go mod init github.com/robertokbr/payment-service`

  - Adding some deps:
    `go get golang.org/x/crypto/bcrypt`
    `go get github.com/satori/go.uuid`
    `github.com/jinzhu/gorm`

### Creating protobufer services
- Instaling go plugin for proto compiler 
  ```bash
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
  ```

- Compiling proto to go_package
  - `protoc --proto_path=protos protos/*.proto --go_out=src/infra/pb  --go_opt=paths=source_relative --go-grpc_out=src/infra/pb --go-grpc_opt=paths=source_relative` 

  

  protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto