package engine

import (
	"quick_web_golang/lib"
	"quick_web_golang/model"
	"quick_web_golang/network"
	"quick_web_golang/provider"
)

func Init() {
	provider.Init()
	//network.Init()// RPC相关的，暂时关闭
}

func Start() {
	provider.Database.Start()
	provider.Cache.Start()
	provider.Sms.Start()
	//provider.SessionManager.Start()

	model.Repos = model.NewRepo(provider.Database.DB)
	//go network.RPCServer.Start()
	//go network.GatewayServer.Start()
	if lib.IsEnableNetwork() {
		network.Run()
	}

}

func Stop() {
	provider.Database.Close()
	provider.Cache.Close()
	//provider.SessionManager.Close()
	//network.GatewayServer.Close()
	//network.RPCServer.Close()
}
