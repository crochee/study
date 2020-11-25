### delve 工具
go get -u github.com/go-delve/delve/cmd/dlv

常用两种模式dlv debug和 dlv attach
### pprof工具
参考标准库pprof写法
对生产的prof文件使用go tool pprof +指定文件.prof即可进入分析模式