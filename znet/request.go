package znet

import "github.com/NanamiZY/Zinx/ziface"

type Request struct {
	conn ziface.IConnection
	msg  ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// 获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// 获取请求消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
