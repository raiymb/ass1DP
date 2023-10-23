package command_pattern

import "fmt"

type LightBulb struct {
	isOn bool
}

func (lb *LightBulb) turnOn() {
	lb.isOn = true
	fmt.Println("Turning the light bulb on")
}

func (lb *LightBulb) turnOff() {
	lb.isOn = false
	fmt.Println("Turning the light bulb off")
}
