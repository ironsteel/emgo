// +build f10x_hd

// Peripheral: CEC_Periph  Consumer Electronics Control (CEC).
// Instances:
//  CEC  mmap.CEC_BASE
// Registers:
//  0x00 32  CFGR
//  0x04 32  OAR
//  0x08 32  PRES
//  0x0C 32  ESR
//  0x10 32  CSR
//  0x14 32  TXD
//  0x18 32  RXD
// Import:
//  stm32/o/f10x_hd/mmap
package cec

const (
	PE   CFGR_Bits = 0x01 << 0 //+ Peripheral Enable.
	IE   CFGR_Bits = 0x01 << 1 //+ Interrupt Enable.
	BTEM CFGR_Bits = 0x01 << 2 //+ Bit Timing Error Mode.
	BPEM CFGR_Bits = 0x01 << 3 //+ Bit Period Error Mode.
)

const (
	OA   OAR_Bits = 0x0F << 0 //+ OA[3:0]: Own Address.
	OA_0 OAR_Bits = 0x01 << 0 //  Bit 0.
	OA_1 OAR_Bits = 0x02 << 0 //  Bit 1.
	OA_2 OAR_Bits = 0x04 << 0 //  Bit 2.
	OA_3 OAR_Bits = 0x08 << 0 //  Bit 3.
)

const ()

const (
	BTE   ESR_Bits = 0x01 << 0 //+ Bit Timing Error.
	BPE   ESR_Bits = 0x01 << 1 //+ Bit Period Error.
	RBTFE ESR_Bits = 0x01 << 2 //+ Rx Block Transfer Finished Error.
	SBE   ESR_Bits = 0x01 << 3 //+ Start Bit Error.
	ACKE  ESR_Bits = 0x01 << 4 //+ Block Acknowledge Error.
	LINE  ESR_Bits = 0x01 << 5 //+ Line Error.
	TBTFE ESR_Bits = 0x01 << 6 //+ Tx Block Transfer Finished Error.
)

const (
	TSOM  CSR_Bits = 0x01 << 0 //+ Tx Start Of Message.
	TEOM  CSR_Bits = 0x01 << 1 //+ Tx End Of Message.
	TERR  CSR_Bits = 0x01 << 2 //+ Tx Error.
	TBTRF CSR_Bits = 0x01 << 3 //+ Tx Byte Transfer Request or Block Transfer Finished.
	RSOM  CSR_Bits = 0x01 << 4 //+ Rx Start Of Message.
	REOM  CSR_Bits = 0x01 << 5 //+ Rx End Of Message.
	RERR  CSR_Bits = 0x01 << 6 //+ Rx Error.
	RBTF  CSR_Bits = 0x01 << 7 //+ Rx Block Transfer Finished.
)

const ()

const ()
