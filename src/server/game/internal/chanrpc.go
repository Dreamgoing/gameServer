package internal

import (
	"github.com/name5566/leaf/gate"
)

///设置gate.Agent到空struct的映射
var agents=make(map[gate.Agent]struct{})
var users=make(map[string]gate.Agent)


func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	agents[a]= struct{}{}



}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	///从agents里删除
	delete(agents,a)
}
