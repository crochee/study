# grpc正确的安装方式

// 安装 Protoc Golang 插件

$ 首先，官网安装proto插件 https://github.com/protocolbuffers/protobuf/releases

$ 其次，安装proto的库文件 go get -u github.com/golang/protobuf/proto

$ 再者，安装生成go文件的插件 go get -u github.com/golang/protobuf/protoc-gen-go

$ 其后,gogoprotobuf有两个插件可以使用,protoc-gen-gogo：和protoc-gen-go生成的文件差不多，性能也几乎一样(稍微快一点点),protoc-gen-gofast：生成的文件更复杂，性能也更高(快5-7倍)

$ gogo go get github.com/gogo/protobuf/protoc-gen-gogo

$ gofast go get github.com/gogo/protobuf/protoc-gen-gofast

$ 安装gogoprotobuf库文件 go get github.com/gogo/protobuf/proto

$ //gogo protoc --gogo_out=. *.proto //gofast protoc --gofast_out=. *.proto

$ 修改json tag 的插件 go get -u github.com/favadi/protoc-go-inject-tag   在proto文件加上// @inject_tag: json:"is_auth"   执行protoc-go-inject-tag -input=./user.pb.go


$ go-micro插件 go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master

$ 参考https://segmentfault.com/a/1190000009277748

## 整体流程
大致就是以 .proto 文件为基础，编写插件对 protoc 进行扩展，编译出不同语言不同模块的源文件。

1）首先定义 .proto 文件；
2）然后由 protoc 将 .proto 文件编译成 protobuf 格式的数据；
3）将 2 中编译后的数据传递到各个插件，生成对应语言、对应模块的源代码。
Go Plugins 用于生成 .pb.go 文件
gRPC Plugins 用于生成 _grpc.pb.go
gRPC-Gateway 则是 pb.gw.go
其中步骤2和3是一起的，只需要在 protoc 编译时传递不同参数即可。

比如以下命令会同时生成 Go、gRPC 、gRPC-Gateway 需要的 3 个文件。

protoc --go_out . --go-grpc_out . --grpc-gateway_out . hello_world.proto
