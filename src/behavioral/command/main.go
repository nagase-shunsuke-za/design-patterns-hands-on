package main

import (
	"design-patterns-hands-on/behavioral/command/command"
	"design-patterns-hands-on/behavioral/command/invoker"
	"design-patterns-hands-on/behavioral/command/receiver"
	"fmt"
)

func main() {
	light := &receiver.Light{}
	lightOnCommand := &command.LightOnCommand{Light: light}
	lightOffCommand := &command.LightOffCommand{Light: light}

	button := invoker.Button{TurnOnCommand: lightOnCommand, TurnOffCommand: lightOffCommand}
	remoteController := invoker.RemoteController{TurnOnCommand: lightOnCommand, TurnOffCommand: lightOffCommand}

	button.Press()
	fmt.Printf("light.IsActive : %t\n", light.IsActive)
	button.Press()
	fmt.Printf("light.IsActive : %t\n", light.IsActive)

	remoteController.TurnOn()
	fmt.Printf("light.IsActive : %t\n", light.IsActive)
	remoteController.TurnOff()
	fmt.Printf("light.IsActive : %t\n", light.IsActive)
}
