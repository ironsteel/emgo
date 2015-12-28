// +build f40_41xxx

package exti

import (
	"mmio"

	"stm32/o/f40_41xxx/exti"
	"stm32/o/f40_41xxx/rcc"
	"stm32/o/f40_41xxx/syscfg"
)

const (
	OTGFS   Lines = 1 << 18 // USB OTG FS Wakeup event.
	Ether   Lines = 1 << 19 // Ethernet Wakeup event.
	OTGHS   Lines = 1 << 20 // USB OTG HS Wakeup event.
	RTCTTS  Lines = 1 << 21 // RTC Tamper and TimeStamp events.
	RTCWKUP Lines = 1 << 22 // RTC Wakeup event.
)

func extip() *exti.EXTI_Periph { return exti.EXTI }

func exticr(n int) *mmio.U32 {
	return (*mmio.U32)(&syscfg.SYSCFG.EXTICR[n].U32)
}
func exticrEna() { rcc.RCC.SYSCFGEN().Set(); _ = rcc.RCC.APB2ENR.Load() }
func exticrDis() { rcc.RCC.SYSCFGEN().Clear() }
