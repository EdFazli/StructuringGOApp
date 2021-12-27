# GO Application Structure using HEX Architecture
How to structure your GO application.
  
![Domain Driven Design Hexagonal Architecture](Images/DDD_Hex_Arch.png)
  
## Directory Setup  
1. `go env | grep GOPATH`  
> GOPATH = D:/GO  
2. `mkdir -p D:/GO/{bin, src, pkg}`  
3. `cd src && mkdir hex`   
4. `cd hex && go mod init`  
5. `mkdir cmd internal`  
6. `cd cmd` and create main.go  
7. 