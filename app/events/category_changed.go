package events

import "github.com/goravel/framework/contracts/event"

type CategoryChanged struct {
}

func (receiver *CategoryChanged) Handle(args []event.Arg) ([]event.Arg, error) {
	return args, nil
}
