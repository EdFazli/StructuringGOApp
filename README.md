# GO Application Structure using DDD HEX Architecture
How to structure your GO application.
  
![Domain Driven Design Hexagonal Architecture](Images/DDD_Hex_Arch.png)
  
## Directory Setup  
1. `go env | grep GOPATH`  
> GOPATH = D:/GO  
2. `mkdir -p D:/GO/{bin, src, pkg}`  
3. `cd src && mkdir hex`   
4. `cd hex && go mod init`  
...  
...  
...  
  
## gRPC - Google Remote Procedure Call  
Documentation
> gRPC : https://grpc.io/docs/languages/go/  
> Download Protocol Buffer : https://github.com/protocolbuffers/protobuf/releases/tag/v3.19.1  
  
#### Prerequisites  
- **Go**
- **Protocol Buffer** compiler, `protoc`, version 3  
> https://github.com/protocolbuffers/protobuf/releases  
>> download protoc-3.19.1-win64.zip  
- **Go Plugins**  
> Install the protocol compiler plugins for Go using the following commands:  
>> $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
>> $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1  

> Update your PATH so that the protoc compiler can find the plugins:  
>> $ export PATH="$PATH:$(go env GOPATH)/bin"  
