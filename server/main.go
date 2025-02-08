package main

import (
	"fmt"
	"github.com/NanamiZY/Zinx/ziface"
	"github.com/NanamiZY/Zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

// Ping Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("Hello Zinx Router V0.6"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer()

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})
	s.Serve()
}
