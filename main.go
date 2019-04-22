package main

import (
	"time"

	"github.com/tinygo-org/tinygo/src/machine"
)

func main() {

	leds := map[int]machine.GPIO{}
	leds[1] = machine.GPIO{13}
	leds[1].Configure(machine.GPIOConfig{Mode: machine.GPIO_OUTPUT})

	leds[2] = machine.GPIO{14}
	leds[2].Configure(machine.GPIOConfig{Mode: machine.GPIO_OUTPUT})

	leds[3] = machine.GPIO{15}
	leds[3].Configure(machine.GPIOConfig{Mode: machine.GPIO_OUTPUT})

	leds[4] = machine.GPIO{16}
	leds[4].Configure(machine.GPIOConfig{Mode: machine.GPIO_OUTPUT})

	button := machine.GPIO{24}
	button.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT_PULLUP})

	leds[1].Set(true)
	leds[2].Set(true)
	leds[3].Set(true)
	leds[4].Set(true)

	ledsOrder := []int{1, 2, 4, 3}
	var pressed = false
	for {
		if !button.Get() {
			if !pressed {

				println("101")

				for j := 1; j <= 4; j++ {
					for _, i := range ledsOrder {
						leds[i].Set(false)
						time.Sleep(time.Millisecond * 100)
						leds[i].Set(true)
					}
				}

			}
			pressed = true
		} else {
			pressed = false
		}

		time.Sleep(time.Microsecond * 10)
	}

}
