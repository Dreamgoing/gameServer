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
	Processor.Register(&Down{})
	Processor.Register(&Command{})
	Processor.Register(&UpLoad{})
	Processor.Register(&Match{})
	Processor.Register(&Admin{})
	Processor.Register(&UserMsg{})
	Processor.Register(&MatchMode{})
	Processor.Register(&Order{})
	Processor.Register(&Finished{})
}

///结构体定义了一个JSON消息格式

///测试消息结构
type Ok struct {
	Name string
}

type Admin struct {
	Name string `json:"name"`
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

///用来同步用户的个人资料
type UpLoad struct {
	ID int `json:"id"`
	Data UserData `json:"data"`
}

type MatchMode struct {
	Name string `json:"name"` ///参加匹配的玩家名
}



///状态消息(向客户端发送)的状态信息
const (
	Login_success = iota
	Login_mismatch
	Login_noexist
	Login_duplicate

	SignUp_success
	SignUp_duplicate
)
type State struct {
	Kind int `json:"kind"`
}


///匹配消息
type Match struct {
	///匹配名
	Name string `json:"name"`
	Car int `json:"car"`
}

type Order struct {
	Name string `json:"name"`
	Val int `json:"val"`
}

///测试的向前开车消息
type Up struct {
	Direction float32
}

///向左转
type Left struct {
	Direction float32
}

///向右转
type Right struct {
	Direction float32
}

type Down struct {
	Direction float32
}

type UserMsg struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
	Context string `json:"context"`
}

type Finished struct {
	Name string `json:"name"`
	Time int    `json:"time"`
}
///定义了具体的命令
const (
	UpCom = iota
	DownCom
	LeftCom
	RightCom
)

type Command struct {
	CarID int `json:"car_id"`
	ID int `json:"id"`
	Cmd int `json:"cmd"`
	Val float32 `json:"val"`
}