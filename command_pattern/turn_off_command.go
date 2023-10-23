package command_pattern

type TurnOffCommand struct {
	Device Device
}

func (c *TurnOffCommand) execute() {
	c.Device.turnOff()
}
