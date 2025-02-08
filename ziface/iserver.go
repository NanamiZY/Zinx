package ziface

// 服务器接口
type IServer interface {
	//启动服务方法
	Start()
	//停止服务方法
	Stop()
	//运行服务方法
	Serve()
	//路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
	AddRouter(msgId uint32, router IRouter)
}
