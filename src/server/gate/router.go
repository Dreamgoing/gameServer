package gate

import (
	"server/msg"
	"server/game"

	"server/login"
)

///gate模块决定了,某一个消息具体交给内部某一个模块来处理
func init() {

	///指定消息到game模块,模块间使用ChanRPC通信
	msg.Processor.SetRouter(&msg.Ok{},game.ChanRPC)

	msg.Processor.SetRouter(&msg.SignUp{},login.ChanRPC)

	msg.Processor.SetRouter(&msg.SignIn{},login.ChanRPC)

	msg.Processor.SetRouter(&msg.Up{},game.ChanRPC)

	msg.Processor.SetRouter(&msg.Left{},game.ChanRPC)

	msg.Processor.SetRouter(&msg.Right{},game.ChanRPC)

	///同步用户个人信息在login的个人模块中
	msg.Processor.SetRouter(&msg.UpLoad{},login.ChanRPC)


}
