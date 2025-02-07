package utils

import (
	"encoding/json"
	"github.com/NanamiZY/Zinx/ziface"
	"os"
)

type GlobalObj struct {
	TcpServer     ziface.IServer //当前github.com/NanamiZY/Zinx的全局Server对象
	Host          string         //当前服务器主机IP
	TcpPort       int            //当前服务器主机监听端口号
	Name          string         //当前服务器名称
	Version       string         //当前github.com/NanamiZY/Zinx版本号
	MaxPacketSize uint32         //都需数据包的最大值
	MaxConn       int            //当前服务器主机允许的最大链接个数
}

var GlobalObject *GlobalObj

// 读取用户配置文件
func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("conf/Zinx.json")
	if err != nil {
		panic(err)
	}
	//将json文件数据解析到struct中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name:          "github.com/NanamiZY/ZinxServerApp",
		Version:       "V0.4",
		TcpPort:       7777,
		Host:          "0.0.0.0",
		MaxConn:       12000,
		MaxPacketSize: 4096,
	}

	GlobalObject.Reload()
}
