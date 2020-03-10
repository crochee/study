package arithmetic

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

/**
* @ Description:
* @Author:
* @Date: 2019/10/10 17:30
 */

// 服务器
type multiEchoServer struct {
	lis          *net.TCPListener // 服务端监听链接
	curClientId  int              // 客户端id
	clients      map[int]*Client  // 存储客户端
	broadcastMsg chan []byte      // 广播消息的通道
}

// 客户端
type Client struct {
	id          int              // 客户端id
	conn        *net.TCPConn     // 客户端链接
	recvMsg     chan []byte      // 接受消息的通道
	sendMsg     chan []byte      // 发送消息的通道
	isRecvClose chan bool        // 关闭接受消息的通道
	isSendClose chan bool        // 关闭发送消息的通道
	mes         *multiEchoServer // 关联的服务器
}

func NewServer() *multiEchoServer { // 初始化服务器
	return &multiEchoServer{
		lis:          nil,
		curClientId:  0,
		clients:      make(map[int]*Client),
		broadcastMsg: make(chan []byte, 10),
	}
}

func (m *multiEchoServer) Start(port string) error { // 启动服务器
	// 获取TCP地址
	addr, err := net.ResolveTCPAddr("tcp", ":"+port)
	if err != nil {
		return err
	}
	m.lis, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	// 监听端口
	// 启动一个goroutine处理广播
	go m.BroadCast()
	// 启动一个goroutine处理客户端的链接
	go func() {
		for {
			conn, err := m.lis.AcceptTCP() // 确认客户端的链接
			if err != nil {
				log.Println(err)
				continue
			}
			cli := &Client{ // 初始化客户端
				id:          m.curClientId,
				conn:        conn,
				recvMsg:     make(chan []byte, 10),
				sendMsg:     make(chan []byte, 10),
				isRecvClose: make(chan bool, 1),
				isSendClose: make(chan bool, 1),
				mes:         m,
			}
			m.clients[m.curClientId] = cli
			m.curClientId++
			// 收发消息处理
			go cli.Recv()
			go cli.Send()
		}
	}()
	return nil
}

func (m *multiEchoServer) Close() {
	m.lis.Close() //
	for _, client := range m.clients {
		client.conn.Close()
		client.isRecvClose <- true
		client.isSendClose <- true
	}
}

func (m *multiEchoServer) Count() int {
	return len(m.clients)
}

// 广播的方法 监听通道是否有消息
func (m *multiEchoServer) BroadCast() {
	for data := range m.broadcastMsg { // 发现广播消息的通道有数据  就分发到各个服务器
		for _, client := range m.clients {
			client.recvMsg <- data
		}
	}
}

// 删除客户端的方法  先将其链接关闭后删除服务器的客户端列表
func (m *multiEchoServer) DelClient(client *Client) error {
	client.conn.Close()
	delete(m.clients, client.id)
	return nil
}

// 客户端发送消息的方法
func (c *Client) Send() {
	defer func() {
		fmt.Println(c.conn.RemoteAddr().String() + "  send exit!")
	}()
	for {
		select {
		case <-c.isSendClose: // 当有消息通知发送的退出的时候就退出
			return
		default:
			reader := bufio.NewReader(c.conn)
			data, err := reader.ReadBytes('\n')
			if err != nil {
				log.Println(err)
				return
			}
			c.mes.broadcastMsg <- data // 读取到消息时，发送到服务器的广播通道
		}
	}
}

func (c *Client) Recv() {
	defer func() {
		fmt.Println(c.conn.RemoteAddr().String() + "  rev exit!")
	}()
	for {
		select {
		case <-c.isRecvClose: // 当接受到要求接受函数退出的时候退出
			return
		case data := <-c.recvMsg: // 当有消息从广播通道到发送通道的时候，将消息写入网络请求
			if _, err := c.conn.Write(data); err != nil {
				return
			}
		}
	}
}