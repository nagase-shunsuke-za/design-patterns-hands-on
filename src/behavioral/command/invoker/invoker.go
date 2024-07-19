package invoker

import (
	"design-patterns-hands-on/behavioral/command/command"
)

type Button struct {
	isActive       bool
	TurnOnCommand  command.ICommand
	TurnOffCommand command.ICommand
}

func (i *Button) Press() {
	if i.isActive {
		i.TurnOffCommand.Execute()
		i.isActive = false
	} else {
		i.TurnOnCommand.Execute()
		i.isActive = true
	}
}

type RemoteController struct {
	TurnOnCommand  command.ICommand
	TurnOffCommand command.ICommand
}

func (i *RemoteController) TurnOn() {
	i.TurnOnCommand.Execute()
}

func (i *RemoteController) TurnOff() {
	i.TurnOffCommand.Execute()
}
