```shell script
# 生成代码
protoc --go_out=/Users/qms/go/src helloworld.proto
protoc --go-grpc_out=/Users/qms/go/src helloworld.proto
```


```shell script
# mac 安装grpc

brew install protobuf
brew install protoc-gen-go
brew install protoc-gen-go-grpc
```