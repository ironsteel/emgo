// +build l1xx_md

// Peripheral: COMP_Periph  Comparator.
// Instances:
//  COMP  mmap.COMP_BASE
// Registers:
//  0x00 32  CSR Comparator control and status register.
// Import:
//  stm32/o/l1xx_md/mmap
package comp

const (
	V10KPU    CSR_Bits = 0x01 << 0  //+ 10K pull-up resistor.
	V400KPU   CSR_Bits = 0x01 << 1  //+ 400K pull-up resistor.
	V10KPD    CSR_Bits = 0x01 << 2  //+ 10K pull-down resistor.
	V400KPD   CSR_Bits = 0x01 << 3  //+ 400K pull-down resistor.
	CMP1EN    CSR_Bits = 0x01 << 4  //+ Comparator 1 enable.
	SW1       CSR_Bits = 0x01 << 5  //+ SW1 analog switch enable.
	CMP1OUT   CSR_Bits = 0x01 << 7  //+ Comparator 1 output.
	SPEED     CSR_Bits = 0x01 << 12 //+ Comparator 2 speed.
	CMP2OUT   CSR_Bits = 0x01 << 13 //+ Comparator 2 ouput.
	VREFOUTEN CSR_Bits = 0x01 << 16 //+ Comparator Vref Enable.
	WNDWE     CSR_Bits = 0x01 << 17 //+ Window mode enable.
	INSEL     CSR_Bits = 0x07 << 18 //+ INSEL[2:0] Inversion input Selection.
	INSEL_0   CSR_Bits = 0x01 << 18 //  Bit 0.
	INSEL_1   CSR_Bits = 0x02 << 18 //  Bit 1.
	INSEL_2   CSR_Bits = 0x04 << 18 //  Bit 2.
	OUTSEL    CSR_Bits = 0x07 << 21 //+ OUTSEL[2:0] comparator 2 output redirection.
	OUTSEL_0  CSR_Bits = 0x01 << 21 //  Bit 0.
	OUTSEL_1  CSR_Bits = 0x02 << 21 //  Bit 1.
	OUTSEL_2  CSR_Bits = 0x04 << 21 //  Bit 2.
	FCH3      CSR_Bits = 0x01 << 26 //+ Bit 26.
	FCH8      CSR_Bits = 0x01 << 27 //+ Bit 27.
	RCH13     CSR_Bits = 0x01 << 28 //+ Bit 28.
	CAIE      CSR_Bits = 0x01 << 29 //+ Bit 29.
	CAIF      CSR_Bits = 0x01 << 30 //+ Bit 30.
	TSUSP     CSR_Bits = 0x01 << 31 //+ Bit 31.
)
