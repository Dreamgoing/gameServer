package internal

import (
	"reflect"
	"server/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"

	"server/login"
)

///@todo 考虑如何简化,服务器端的游戏逻辑设计

///@todo 客户端和服务器端,设置计时器,和超时等功能

///当前广播,适用于弱实时性项目,

///完成匹配记时

///

func init() {
	//向当前模块(game 模块)注册Ok 消息的消息处理函数 handleOk
	handler(&msg.Ok{},handleOk)
	handler(&msg.Up{},handleUp)
	handler(&msg.Left{},handleLeft)
	handler(&msg.Right{},handleRight)
	handler(&msg.Match{},handleMatch)
	handler(&msg.Admin{},handleAdmin)
}

func handler(m interface{},h interface{})  {

	skeleton.RegisterChanRPC(reflect.TypeOf(m),h)
}

///Ok 的消息处理函数
func handleOk(args []interface{}) {
	///收到的 Ok 消息
	m := args[0].(*msg.Ok)

	//log.Debug("%v",m)
	///消息的发送者
	a := args[1].(gate.Agent)

	log.Debug("Ok %v", m.Name)

	///给发送者回应一个Ok消息
	a.WriteMsg(&msg.Ok{
		Name: "client",
	})
}

func handleUp(args []interface{})  {
	///收到的 up消息
	 m:=args[0].(*msg.Up)

	//获取网关,即客户端地址,消息的发送着

	a:=args[1].(gate.Agent)


	///此处可以处理相应得逻辑

	///将命令消息广播转发到其他的分组,设置相应的速度
	log.Debug("%v %v",m.Direction,a)

	///向客户端发送,确认接收的消息,注意此处的Car为指针类型
	tmp:=a.UserData().(*msg.Car)

	tmp.Up()

	log.Debug("%v %v %v",a.RemoteAddr(),tmp.X,tmp.Y)



	///广播车的命令
	handleBroadcast(&msg.Command{CarID:tmp.CarID,ID:tmp.CarID,Cmd:msg.UpCom,Val:m.Direction})

}

func handleLeft(args []interface{})  {
	m:=args[0].(*msg.Left)

	a:=args[1].(gate.Agent)

	tmp:=a.UserData().(*msg.Car)
	tmp.Left()

	log.Debug("%v %v",m.Direction,a)
	log.Debug("%v %v %v",a.RemoteAddr(),tmp.X,tmp.Y)


	///广播向左转的命令
	handleBroadcast(&msg.Command{tmp.CarID,tmp.CarID,msg.LeftCom,m.Direction})

}
func handleRight(args []interface{})  {
	m:=args[0].(*msg.Right)

	a:=args[1].(gate.Agent)

	tmp:=a.UserData().(*msg.Car)
	tmp.Right()

	log.Debug("%v %v",m.Direction,a)
	log.Debug("%v %v %v",a.RemoteAddr(),tmp.X,tmp.Y)


	///广播向左转的命令
	handleBroadcast(&msg.Command{tmp.CarID,tmp.CarID,msg.RightCom,m.Direction})

}

func handleMatch(args []interface{})  {
	///处理当前匹配多人模式的消息

	m,ok:=args[0].(msg.Match)
	if ok {
		log.Debug("multi-player match game %v",m)
	}
	///从当前匹配队列中,找到满足条件的匹配

}

///广播
func handleBroadcast(cmd *msg.Command){
	log.Debug("%v",len(agents))
	for a:=range agents{
		log.Debug("%v\n",a.RemoteAddr().String())
		a.WriteMsg(cmd)
	}
}

func handleAdmin(args []interface{})  {
	if val,ok:=args[0].(*msg.Admin);ok{
		log.Debug("Admin login: %s",val.Name)
	}
	if a,ok:=args[1].(gate.Agent);ok{
		log.Debug("Admin ip: %v\n",a.RemoteAddr().String())
		for it:=range agents {
			log.Debug("%v\n",it.RemoteAddr().String())
		}
	}

}

///用来处理不同用户之间的聊天功能
func handleUserMsg(args []interface{})  {

	m,ok:=args[0].(*msg.UserMsg)
	if ok{
		log.Debug("user message %v",m)
	}


	a, ok := args[1].(gate.Agent)
	if ok{
		log.Debug("send from: %v",a.RemoteAddr().String())
	}

	///a发送消息到b
	bname:=m.Dst

	///@todo 对于这里的判断要进行用户当前是否在线处理
	bgate:=login.UserAgent[bname]

	if bgate==nil {
		///没找到对应用户的gate,进行错误处理

	}

	///转发给用户b
	bgate.WriteMsg(&msg.UserMsg{m.Src,m.Dst,m.Context})

}