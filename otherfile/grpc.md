# grpc正确的安装方式

//安装gRPC Runtime
 
$ git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
 
$ git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
 
$ git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
 
// 安装 Protoc Golang 插件
 
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
 
$ git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
 
$ cd $GOPATH/src/
 
$ go install google.golang.org/grpc