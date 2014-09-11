package main

import (
	"stm32/f4/gpio"
	"stm32/f4/periph"
)

var leds = gpio.D

const (
	Green = 12 + iota
	Orange
	Red
	Blue
)

func initLEDs() {
	periph.AHB1ClockEnable(periph.GPIOD)
	periph.AHB1Reset(periph.GPIOD)

	leds.SetMode(Green, gpio.Out)
	leds.SetMode(Orange, gpio.Out)
	leds.SetMode(Red, gpio.Out)
	leds.SetMode(Blue, gpio.Out)
}