package main

import (
	"fmt"

	"stm32/hal/raw/pwr"
	"stm32/hal/raw/rcc"
	"stm32/hal/raw/rtc"
)

func initRTC() {
	const (
		mask = rcc.LSEON | rcc.RTCSEL | rcc.RTCEN
		cfg  = rcc.LSEON | rcc.RTCSEL_LSE | rcc.RTCEN
	)
	RCC := rcc.RCC
	PWR := pwr.PWR
	if RCC.CSR.Bits(mask) != cfg {
		fmt.Println("Configuring backup domain...")
		RCC.PWREN().Set()
		RCC.PWREN().Load()
		PWR.DBP().Set()
		RCC.RTCRST().Set()
		RCC.RTCRST().Clear()
		RCC.CSR.StoreBits(mask, cfg)
		for RCC.LSERDY().Load() == 0 {
		}
		PWR.DBP().Clear()
		RCC.PWREN().Clear()
		fmt.Println("Done.")
	}

}

func checkRTC() bool {
	return rtc.RTC.INITS().Load() != 0
}

func setRTC(t DateTime) {
	RCC := rcc.RCC
	RTC := rtc.RTC
	PWR := pwr.PWR
	RCC.PWREN().Set()
	RCC.PWREN().Load()
	PWR.DBP().Set()
	RTC.WPR.Store(0xca)
	RTC.WPR.Store(0x53)
	RTC.INIT().Set()
	for RTC.INITF().Load() == 0 {
	}
	//const prer = (2-1)<<16 + (1 - 1)
	//RTC.PRER.Store(prer)
	//RTC.PRER.Store(prer)
	RTC.DR.U32.Store(t.dr)
	RTC.TR.U32.Store(t.tr)
	RTC.INIT().Clear()
	RTC.WPR.Store(0xff)
	PWR.DBP().Clear()
	RCC.PWREN().Clear()
}

type Weekday byte

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

//emgo:const
var wdayStr = [7]string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

func (wd Weekday) String() string {
	return wdayStr[wd-1]
}

type DateTime struct {
	tr uint32
	dr uint32
}

func readRTC() DateTime {
	RTC := rtc.RTC
	var t DateTime
	t.dr = RTC.DR.U32.Load()
	for {
		t.tr = RTC.TR.U32.Load()
		dr = RTC.DR.U32.Load()
		if dr == t.dr {
			return t
		}
		t.dr = dr
	}
}

func (t DateTime) Year() int {
	return 2000 + int(t.dr>>20&0xf*10+t.dr>>16&0xf)
}

func (t DateTime) Month() int {
	return int(t.dr>>12&1*10 + t.dr>>8&0xf)
}

func (t DateTime) Day() int {
	return int(t.dr>>4&3*10 + t.dr&0xf)
}

func (t DateTime) Weekday() Weekday {
	return Weekday(t.dr >> 13 & 7)
}

func (t DateTime) Hour() int {
	return int(t.tr>>20&3*10 + t.tr>>16&0xf)
}

func (t DateTime) Minute() int {
	return int(t.tr>>12&7*10 + t.tr>>8&0xf)
}

func (t DateTime) Second() int {
	return int(t.tr>>4&7*10 + t.tr&0xf)
}

func makeDateTime(Y, M, D, h, m, s int, wd Weekday) (t DateTime) {
	Y -= 2000
	t.dr = uint32(Y/10<<20 + Y%10<<16 + M/10<<12 + M%10<<8 + D/10<<4 + D%10 + int(wd)<<13)
	t.tr = uint32(h/10<<20 + h%10<<16 + m/10<<12 + m%10<<8 + s/10<<4 + s%10)
	return
}