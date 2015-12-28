// +build f10x_hd

// Peripheral: DBGMCU_Periph  Debug MCU.
// Instances:
//  DBGMCU  mmap.DBGMCU_BASE
// Registers:
//  0x00 32  IDCODE
//  0x04 32  CR
// Import:
//  stm32/o/f10x_hd/mmap
package dbgmcu

const (
	DEV_ID    IDCODE_Bits = 0xFFF << 0   //+ Device Identifier.
	REV_ID    IDCODE_Bits = 0xFFFF << 16 //+ REV_ID[15:0] bits (Revision Identifier).
	REV_ID_0  IDCODE_Bits = 0x01 << 16   //  Bit 0.
	REV_ID_1  IDCODE_Bits = 0x02 << 16   //  Bit 1.
	REV_ID_2  IDCODE_Bits = 0x04 << 16   //  Bit 2.
	REV_ID_3  IDCODE_Bits = 0x08 << 16   //  Bit 3.
	REV_ID_4  IDCODE_Bits = 0x10 << 16   //  Bit 4.
	REV_ID_5  IDCODE_Bits = 0x20 << 16   //  Bit 5.
	REV_ID_6  IDCODE_Bits = 0x40 << 16   //  Bit 6.
	REV_ID_7  IDCODE_Bits = 0x80 << 16   //  Bit 7.
	REV_ID_8  IDCODE_Bits = 0x100 << 16  //  Bit 8.
	REV_ID_9  IDCODE_Bits = 0x200 << 16  //  Bit 9.
	REV_ID_10 IDCODE_Bits = 0x400 << 16  //  Bit 10.
	REV_ID_11 IDCODE_Bits = 0x800 << 16  //  Bit 11.
	REV_ID_12 IDCODE_Bits = 0x1000 << 16 //  Bit 12.
	REV_ID_13 IDCODE_Bits = 0x2000 << 16 //  Bit 13.
	REV_ID_14 IDCODE_Bits = 0x4000 << 16 //  Bit 14.
	REV_ID_15 IDCODE_Bits = 0x8000 << 16 //  Bit 15.
)

const (
	DBG_SLEEP              CR_Bits = 0x01 << 0  //+ Debug Sleep Mode.
	DBG_STOP               CR_Bits = 0x01 << 1  //+ Debug Stop Mode.
	DBG_STANDBY            CR_Bits = 0x01 << 2  //+ Debug Standby mode.
	TRACE_IOEN             CR_Bits = 0x01 << 5  //+ Trace Pin Assignment Control.
	TRACE_MODE             CR_Bits = 0x03 << 6  //+ TRACE_MODE[1:0] bits (Trace Pin Assignment Control).
	TRACE_MODE_0           CR_Bits = 0x01 << 6  //  Bit 0.
	TRACE_MODE_1           CR_Bits = 0x02 << 6  //  Bit 1.
	DBG_IWDG_STOP          CR_Bits = 0x01 << 8  //+ Debug Independent Watchdog stopped when Core is halted.
	DBG_WWDG_STOP          CR_Bits = 0x01 << 9  //+ Debug Window Watchdog stopped when Core is halted.
	DBG_TIM1_STOP          CR_Bits = 0x01 << 10 //+ TIM1 counter stopped when core is halted.
	DBG_TIM2_STOP          CR_Bits = 0x01 << 11 //+ TIM2 counter stopped when core is halted.
	DBG_TIM3_STOP          CR_Bits = 0x01 << 12 //+ TIM3 counter stopped when core is halted.
	DBG_TIM4_STOP          CR_Bits = 0x01 << 13 //+ TIM4 counter stopped when core is halted.
	DBG_CAN1_STOP          CR_Bits = 0x01 << 14 //+ Debug CAN1 stopped when Core is halted.
	DBG_I2C1_SMBUS_TIMEOUT CR_Bits = 0x01 << 15 //+ SMBUS timeout mode stopped when Core is halted.
	DBG_I2C2_SMBUS_TIMEOUT CR_Bits = 0x01 << 16 //+ SMBUS timeout mode stopped when Core is halted.
	DBG_TIM8_STOP          CR_Bits = 0x01 << 17 //+ TIM8 counter stopped when core is halted.
	DBG_TIM5_STOP          CR_Bits = 0x01 << 18 //+ TIM5 counter stopped when core is halted.
	DBG_TIM6_STOP          CR_Bits = 0x01 << 19 //+ TIM6 counter stopped when core is halted.
	DBG_TIM7_STOP          CR_Bits = 0x01 << 20 //+ TIM7 counter stopped when core is halted.
	DBG_CAN2_STOP          CR_Bits = 0x01 << 21 //+ Debug CAN2 stopped when Core is halted.
	DBG_TIM15_STOP         CR_Bits = 0x01 << 22 //+ Debug TIM15 stopped when Core is halted.
	DBG_TIM16_STOP         CR_Bits = 0x01 << 23 //+ Debug TIM16 stopped when Core is halted.
	DBG_TIM17_STOP         CR_Bits = 0x01 << 24 //+ Debug TIM17 stopped when Core is halted.
	DBG_TIM12_STOP         CR_Bits = 0x01 << 25 //+ Debug TIM12 stopped when Core is halted.
	DBG_TIM13_STOP         CR_Bits = 0x01 << 26 //+ Debug TIM13 stopped when Core is halted.
	DBG_TIM14_STOP         CR_Bits = 0x01 << 27 //+ Debug TIM14 stopped when Core is halted.
	DBG_TIM9_STOP          CR_Bits = 0x01 << 28 //+ Debug TIM9 stopped when Core is halted.
	DBG_TIM10_STOP         CR_Bits = 0x01 << 29 //+ Debug TIM10 stopped when Core is halted.
	DBG_TIM11_STOP         CR_Bits = 0x01 << 30 //+ Debug TIM11 stopped when Core is halted.
)
