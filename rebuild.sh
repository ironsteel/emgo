#!/bin/bash

set -e

rm -rf egroot/pkg/* egpath/pkg/*

list="
builtin

sync/atomic
sync/barrier
delay
mmio
bits

cortexm
cortexm/fpu
cortexm/exce
cortexm/systick
cortexm/sleep

runtime/noos
sync
runtime


strconv
math/matrix32

stm32/stlink
stm32/f4/clock
stm32/f4/flash
stm32/f4/gpio
stm32/f4/periph
stm32/f4/setup
stm32/f4/exti
stm32/f4/irq

stm32/l1/clock
stm32/l1/flash
stm32/l1/gpio
stm32/l1/periph
stm32/l1/setup
stm32/l1/exti
stm32/l1/irq
"
for p in $list; do 
	echo $p
	egc $p
done
