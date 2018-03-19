// This example shows how to use GPIO to as digital output (to blink connected
// LEDs) and input (to read state of KEY1).
package main

import (
	"delay"
	"fmt"

	"nrf5/hal/clock"
	"nrf5/hal/gpio"
	"nrf5/hal/irq"
	"nrf5/hal/rtc"
	"nrf5/hal/system"
	"nrf5/hal/system/timer/rtcst"
)

var leds [4]gpio.Pin

func init() {
	// Initialize system and runtime.
	system.Setup(clock.XTAL, clock.XTAL, true)
	rtcst.Setup(rtc.RTC0, 1)

	// Allocate pins (always do it in one place to avoid conflicts).
	p0 := gpio.P0
	leds[0] = p0.Pin(24)
	leds[1] = p0.Pin(20)
	leds[2] = p0.Pin(15)
	leds[3] = p0.Pin(11)

	// Configure pins.
	for _, led := range leds {
		led.Setup(gpio.ModeOut | gpio.DriveD0H1)
	}
}

func main() {
	for {
		fmt.Printf("bla\n")
		for _, led := range leds {
			led.Set()
			delay.Millisec(200)
			led.Clear()
		}
	}
}

//emgo:const
//c:__attribute__((section(".ISRs")))
var ISRs = [...]func(){
	irq.RTC0: rtcst.ISR,
}
