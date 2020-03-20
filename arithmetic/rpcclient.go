package arithmetic

import (
	"fmt"
	"log"
	"net/rpc"
)

/**
* @ Description:
* @Author:
* @Date: 2020/3/10 17:40
 */
// const HelloServiceName = "HelloService"

func clientmain() {
	c, err := DailHelloService("tcp", "10.2.4.39:13149")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = c.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

// var HelloServiceInterface = (*HelloServiceClient)(nil)

type HelloServiceClient struct {
	*rpc.Client
}

func DailHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
