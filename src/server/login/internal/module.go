package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
	"github.com/name5566/leaf/gate"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	UserAgent=make(map[string]gate.Agent)
}

func (m *Module) OnDestroy() {

}
