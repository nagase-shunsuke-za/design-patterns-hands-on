package command

import (
	"design-patterns-hands-on/behavioral/command/receiver"
)

type LightOnCommand struct {
	Light *receiver.Light
}

func (c *LightOnCommand) Execute() {
	c.Light.TurnOn()
}

type LightOffCommand struct {
	Light *receiver.Light
}

func (c *LightOffCommand) Execute() {
	c.Light.TurnOff()
}
