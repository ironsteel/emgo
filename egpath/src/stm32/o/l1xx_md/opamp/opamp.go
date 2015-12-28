// +build l1xx_md

// Peripheral: OPAMP_Periph  Operational Amplifier (OPAMP).
// Instances:
//  OPAMP  mmap.OPAMP_BASE
// Registers:
//  0x00 32  CSR   Control/status register.
//  0x04 32  OTR   Offset trimming register for normal mode.
//  0x08 32  LPOTR Offset trimming register for low power mode.
// Import:
//  stm32/o/l1xx_md/mmap
package opamp

const (
	OPA1PD     CSR_Bits = 0x01 << 0  //+ OPAMP1 disable.
	S3SEL1     CSR_Bits = 0x01 << 1  //+ Switch 3 for OPAMP1 Enable.
	S4SEL1     CSR_Bits = 0x01 << 2  //+ Switch 4 for OPAMP1 Enable.
	S5SEL1     CSR_Bits = 0x01 << 3  //+ Switch 5 for OPAMP1 Enable.
	S6SEL1     CSR_Bits = 0x01 << 4  //+ Switch 6 for OPAMP1 Enable.
	OPA1CAL_L  CSR_Bits = 0x01 << 5  //+ OPAMP1 Offset calibration for P differential pair.
	OPA1CAL_H  CSR_Bits = 0x01 << 6  //+ OPAMP1 Offset calibration for N differential pair.
	OPA1LPM    CSR_Bits = 0x01 << 7  //+ OPAMP1 Low power enable.
	OPA2PD     CSR_Bits = 0x01 << 8  //+ OPAMP2 disable.
	S3SEL2     CSR_Bits = 0x01 << 9  //+ Switch 3 for OPAMP2 Enable.
	S4SEL2     CSR_Bits = 0x01 << 10 //+ Switch 4 for OPAMP2 Enable.
	S5SEL2     CSR_Bits = 0x01 << 11 //+ Switch 5 for OPAMP2 Enable.
	S6SEL2     CSR_Bits = 0x01 << 12 //+ Switch 6 for OPAMP2 Enable.
	OPA2CAL_L  CSR_Bits = 0x01 << 13 //+ OPAMP2 Offset calibration for P differential pair.
	OPA2CAL_H  CSR_Bits = 0x01 << 14 //+ OPAMP2 Offset calibration for N differential pair.
	OPA2LPM    CSR_Bits = 0x01 << 15 //+ OPAMP2 Low power enable.
	OPA3PD     CSR_Bits = 0x01 << 16 //+ OPAMP3 disable.
	S3SEL3     CSR_Bits = 0x01 << 17 //+ Switch 3 for OPAMP3 Enable.
	S4SEL3     CSR_Bits = 0x01 << 18 //+ Switch 4 for OPAMP3 Enable.
	S5SEL3     CSR_Bits = 0x01 << 19 //+ Switch 5 for OPAMP3 Enable.
	S6SEL3     CSR_Bits = 0x01 << 20 //+ Switch 6 for OPAMP3 Enable.
	OPA3CAL_L  CSR_Bits = 0x01 << 21 //+ OPAMP3 Offset calibration for P differential pair.
	OPA3CAL_H  CSR_Bits = 0x01 << 22 //+ OPAMP3 Offset calibration for N differential pair.
	OPA3LPM    CSR_Bits = 0x01 << 23 //+ OPAMP3 Low power enable.
	ANAWSEL1   CSR_Bits = 0x01 << 24 //+ Switch ANA Enable for OPAMP1.
	ANAWSEL2   CSR_Bits = 0x01 << 25 //+ Switch ANA Enable for OPAMP2.
	ANAWSEL3   CSR_Bits = 0x01 << 26 //+ Switch ANA Enable for OPAMP3.
	S7SEL2     CSR_Bits = 0x01 << 27 //+ Switch 7 for OPAMP2 Enable.
	AOP_RANGE  CSR_Bits = 0x01 << 28 //+ Power range selection.
	OPA1CALOUT CSR_Bits = 0x01 << 29 //+ OPAMP1 calibration output.
	OPA2CALOUT CSR_Bits = 0x01 << 30 //+ OPAMP2 calibration output.
	OPA3CALOUT CSR_Bits = 0x01 << 31 //+ OPAMP3 calibration output.
)

const (
	AO1_OPT_OFFSET_TRIM OTR_Bits = 0x3FF << 0  //+ Offset trim for OPAMP1.
	AO2_OPT_OFFSET_TRIM OTR_Bits = 0x3FF << 10 //+ Offset trim for OPAMP2.
	AO3_OPT_OFFSET_TRIM OTR_Bits = 0x3FF << 20 //+ Offset trim for OPAMP2.
	OT_USER             OTR_Bits = 0x01 << 31  //+ Switch to OPAMP offset user trimmed values.
)
