package msg

import (
	//"github.com/name5566/leaf/network"
	"github.com/name5566/leaf/network/json"
)

//var Processor network.Processor
var Processor = json.NewProcessor()

//var SignUp_processor = json.NewProcessor()

func init() {

	///注册json消息
	Processor.Register(&Ok{})
	Processor.Register(&SignUp{})
	Processor.Register(&SignIn{})
	Processor.Register(&State{})
	Processor.Register(&Up{})
	Processor.Register(&Right{})
	Processor.Register(&Left{})
}

///结构体定义了一个JSON消息格式

///测试消息结构
type Ok struct {
	Name string
}




///注册消息结构
type SignUp struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

///登录消息结构
type SignIn struct {
	Name string `json:"name"`
	Password string `json:"password"`
}



///状态消息(向服务器发送)
type State struct {
	Name string `json:"name"`
}

///测试的向前开车消息
type Up struct {
	Direction int
}

///向左转
type Left struct {
	Direction int
}

///向右转
type Right struct {
	Direction int
}