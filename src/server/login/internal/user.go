package internal

import (
	"github.com/name5566/leaf/gate"
)

const (
	userLogin = iota
	userLogout
	userGame
)

type User struct {
	gate.Agent
	state int
	data *UserData ///用户的数据
	//*g.LinearContext
}
