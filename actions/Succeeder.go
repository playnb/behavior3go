package actions

import (
	b3 "github.com/playnb/behavior3go"
	. "github.com/playnb/behavior3go/core"
)

type Succeeder struct {
	Action
}

func (this *Succeeder) OnTick(tick *Tick) b3.Status {
	return b3.SUCCESS
}
