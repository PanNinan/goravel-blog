package console

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/schedule"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/carbon"
)

type Kernel struct {
}

func (kernel Kernel) Schedule() []schedule.Event {
	return []schedule.Event{
		facades.Schedule().Call(func() {
			facades.Log().Debug(carbon.Now("PRC").ToDateTimeString())
		}).EveryFifteenMinutes(),
	}
}

func (kernel Kernel) Commands() []console.Command {
	return []console.Command{}
}
