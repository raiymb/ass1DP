package command_pattern

type TurnOnCommand struct {
	Device Device
}

func (c *TurnOnCommand) execute() {
	c.Device.turnOn()
}
