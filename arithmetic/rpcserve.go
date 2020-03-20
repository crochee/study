package arithmetic

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

/**
* @ Description:
* @Author:
* @Date: 2020/3/10 17:41
 */
func servermain() {
	RegisterHelloService(new(HelloService))
	l, err := net.Listen("tcp", ":13149")
	if err != nil {
		log.Fatal("listen err:", err)
	}
	cnn, err := l.Accept()
	if err != nil {
		log.Fatal("accept err:", err)
	}
	rpc.ServeConn(cnn)
	fmt.Println("finis!")
}

const HelloServiceName = "HelloService" // 结构体名需要告诉对方

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
	SayBye(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func (h *HelloService) SayBye(request string, reply *string) error {
	*reply = "goodbye:" + request
	return nil
}
