package login

import (
	"server/login/internal"
)

var (
	Module  = new(internal.Module)

	UserDB = internal.Userdb
	ChanRPC = internal.ChanRPC
	UserAgent = internal.UserAgent

)

