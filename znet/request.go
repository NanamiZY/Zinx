package znet

import "github.com/NanamiZY/Zinx/ziface"

type Request struct {
	conn ziface.IConnection
	data []byte
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
