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
