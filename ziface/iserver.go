package ziface

// 服务器接口
type IServer interface {
	Start()
	Stop()
	Serve()
}
