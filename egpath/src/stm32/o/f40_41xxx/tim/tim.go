// +build f40_41xxx

// Peripheral: TIM_Periph  TIM.
// Instances:
//  TIM2   mmap.TIM2_BASE
//  TIM3   mmap.TIM3_BASE
//  TIM4   mmap.TIM4_BASE
//  TIM5   mmap.TIM5_BASE
//  TIM6   mmap.TIM6_BASE
//  TIM7   mmap.TIM7_BASE
//  TIM12  mmap.TIM12_BASE
//  TIM13  mmap.TIM13_BASE
//  TIM14  mmap.TIM14_BASE
//  TIM1   mmap.TIM1_BASE
//  TIM8   mmap.TIM8_BASE
//  TIM9   mmap.TIM9_BASE
//  TIM10  mmap.TIM10_BASE
//  TIM11  mmap.TIM11_BASE
// Registers:
//  0x00 16  CR1   Control register 1.
//  0x04 16  CR2   Control register 2.
//  0x08 16  SMCR  Slave mode control register.
//  0x0C 16  DIER  DMA/interrupt enable register.
//  0x10 16  SR    Status register.
//  0x14 16  EGR   Event generation register.
//  0x18 16  CCMR1 Capture/compare mode register 1.
//  0x1C 16  CCMR2 Capture/compare mode register 2.
//  0x20 16  CCER  Capture/compare enable register.
//  0x24 32  CNT   Counter register.
//  0x28 16  PSC   Prescaler.
//  0x2C 32  ARR   Auto-reload register.
//  0x30 16  RCR   Repetition counter register.
//  0x34 32  CCR1  Capture/compare register 1.
//  0x38 32  CCR2  Capture/compare register 2.
//  0x3C 32  CCR3  Capture/compare register 3.
//  0x40 32  CCR4  Capture/compare register 4.
//  0x44 16  BDTR  Break and dead-time register.
//  0x48 16  DCR   DMA control register.
//  0x4C 16  DMAR  DMA address for full transfer.
//  0x50 16  OR    Option register.
// Import:
//  stm32/o/f40_41xxx/mmap
package tim

const (
	CEN   CR1_Bits = 0x01 << 0 //+ Counter enable.
	UDIS  CR1_Bits = 0x01 << 1 //+ Update disable.
	URS   CR1_Bits = 0x01 << 2 //+ Update request source.
	OPM   CR1_Bits = 0x01 << 3 //+ One pulse mode.
	DIR   CR1_Bits = 0x01 << 4 //+ Direction.
	CMS   CR1_Bits = 0x03 << 5 //+ CMS[1:0] bits (Center-aligned mode selection).
	CMS_0 CR1_Bits = 0x01 << 5 //  Bit 0.
	CMS_1 CR1_Bits = 0x02 << 5 //  Bit 1.
	ARPE  CR1_Bits = 0x01 << 7 //+ Auto-reload preload enable.
	CKD   CR1_Bits = 0x03 << 8 //+ CKD[1:0] bits (clock division).
	CKD_0 CR1_Bits = 0x01 << 8 //  Bit 0.
	CKD_1 CR1_Bits = 0x02 << 8 //  Bit 1.
)

const (
	CCPC  CR2_Bits = 0x01 << 0  //+ Capture/Compare Preloaded Control.
	CCUS  CR2_Bits = 0x01 << 2  //+ Capture/Compare Control Update Selection.
	CCDS  CR2_Bits = 0x01 << 3  //+ Capture/Compare DMA Selection.
	MMS   CR2_Bits = 0x07 << 4  //+ MMS[2:0] bits (Master Mode Selection).
	MMS_0 CR2_Bits = 0x01 << 4  //  Bit 0.
	MMS_1 CR2_Bits = 0x02 << 4  //  Bit 1.
	MMS_2 CR2_Bits = 0x04 << 4  //  Bit 2.
	TI1S  CR2_Bits = 0x01 << 7  //+ TI1 Selection.
	OIS1  CR2_Bits = 0x01 << 8  //+ Output Idle state 1 (OC1 output).
	OIS1N CR2_Bits = 0x01 << 9  //+ Output Idle state 1 (OC1N output).
	OIS2  CR2_Bits = 0x01 << 10 //+ Output Idle state 2 (OC2 output).
	OIS2N CR2_Bits = 0x01 << 11 //+ Output Idle state 2 (OC2N output).
	OIS3  CR2_Bits = 0x01 << 12 //+ Output Idle state 3 (OC3 output).
	OIS3N CR2_Bits = 0x01 << 13 //+ Output Idle state 3 (OC3N output).
	OIS4  CR2_Bits = 0x01 << 14 //+ Output Idle state 4 (OC4 output).
)

const (
	SMS    SMCR_Bits = 0x07 << 0  //+ SMS[2:0] bits (Slave mode selection).
	SMS_0  SMCR_Bits = 0x01 << 0  //  Bit 0.
	SMS_1  SMCR_Bits = 0x02 << 0  //  Bit 1.
	SMS_2  SMCR_Bits = 0x04 << 0  //  Bit 2.
	TS     SMCR_Bits = 0x07 << 4  //+ TS[2:0] bits (Trigger selection).
	TS_0   SMCR_Bits = 0x01 << 4  //  Bit 0.
	TS_1   SMCR_Bits = 0x02 << 4  //  Bit 1.
	TS_2   SMCR_Bits = 0x04 << 4  //  Bit 2.
	MSM    SMCR_Bits = 0x01 << 7  //+ Master/slave mode.
	ETF    SMCR_Bits = 0x0F << 8  //+ ETF[3:0] bits (External trigger filter).
	ETF_0  SMCR_Bits = 0x01 << 8  //  Bit 0.
	ETF_1  SMCR_Bits = 0x02 << 8  //  Bit 1.
	ETF_2  SMCR_Bits = 0x04 << 8  //  Bit 2.
	ETF_3  SMCR_Bits = 0x08 << 8  //  Bit 3.
	ETPS   SMCR_Bits = 0x03 << 12 //+ ETPS[1:0] bits (External trigger prescaler).
	ETPS_0 SMCR_Bits = 0x01 << 12 //  Bit 0.
	ETPS_1 SMCR_Bits = 0x02 << 12 //  Bit 1.
	ECE    SMCR_Bits = 0x01 << 14 //+ External clock enable.
	ETP    SMCR_Bits = 0x01 << 15 //+ External trigger polarity.
)

const (
	UIE   DIER_Bits = 0x01 << 0  //+ Update interrupt enable.
	CC1IE DIER_Bits = 0x01 << 1  //+ Capture/Compare 1 interrupt enable.
	CC2IE DIER_Bits = 0x01 << 2  //+ Capture/Compare 2 interrupt enable.
	CC3IE DIER_Bits = 0x01 << 3  //+ Capture/Compare 3 interrupt enable.
	CC4IE DIER_Bits = 0x01 << 4  //+ Capture/Compare 4 interrupt enable.
	COMIE DIER_Bits = 0x01 << 5  //+ COM interrupt enable.
	TIE   DIER_Bits = 0x01 << 6  //+ Trigger interrupt enable.
	BIE   DIER_Bits = 0x01 << 7  //+ Break interrupt enable.
	UDE   DIER_Bits = 0x01 << 8  //+ Update DMA request enable.
	CC1DE DIER_Bits = 0x01 << 9  //+ Capture/Compare 1 DMA request enable.
	CC2DE DIER_Bits = 0x01 << 10 //+ Capture/Compare 2 DMA request enable.
	CC3DE DIER_Bits = 0x01 << 11 //+ Capture/Compare 3 DMA request enable.
	CC4DE DIER_Bits = 0x01 << 12 //+ Capture/Compare 4 DMA request enable.
	COMDE DIER_Bits = 0x01 << 13 //+ COM DMA request enable.
	TDE   DIER_Bits = 0x01 << 14 //+ Trigger DMA request enable.
)

const (
	UIF   SR_Bits = 0x01 << 0  //+ Update interrupt Flag.
	CC1IF SR_Bits = 0x01 << 1  //+ Capture/Compare 1 interrupt Flag.
	CC2IF SR_Bits = 0x01 << 2  //+ Capture/Compare 2 interrupt Flag.
	CC3IF SR_Bits = 0x01 << 3  //+ Capture/Compare 3 interrupt Flag.
	CC4IF SR_Bits = 0x01 << 4  //+ Capture/Compare 4 interrupt Flag.
	COMIF SR_Bits = 0x01 << 5  //+ COM interrupt Flag.
	TIF   SR_Bits = 0x01 << 6  //+ Trigger interrupt Flag.
	BIF   SR_Bits = 0x01 << 7  //+ Break interrupt Flag.
	CC1OF SR_Bits = 0x01 << 9  //+ Capture/Compare 1 Overcapture Flag.
	CC2OF SR_Bits = 0x01 << 10 //+ Capture/Compare 2 Overcapture Flag.
	CC3OF SR_Bits = 0x01 << 11 //+ Capture/Compare 3 Overcapture Flag.
	CC4OF SR_Bits = 0x01 << 12 //+ Capture/Compare 4 Overcapture Flag.
)

const (
	UG   EGR_Bits = 0x01 << 0 //+ Update Generation.
	CC1G EGR_Bits = 0x01 << 1 //+ Capture/Compare 1 Generation.
	CC2G EGR_Bits = 0x01 << 2 //+ Capture/Compare 2 Generation.
	CC3G EGR_Bits = 0x01 << 3 //+ Capture/Compare 3 Generation.
	CC4G EGR_Bits = 0x01 << 4 //+ Capture/Compare 4 Generation.
	COMG EGR_Bits = 0x01 << 5 //+ Capture/Compare Control Update Generation.
	TG   EGR_Bits = 0x01 << 6 //+ Trigger Generation.
	BG   EGR_Bits = 0x01 << 7 //+ Break Generation.
)

const (
	CC1S     CCMR1_Bits = 0x03 << 0  //+ CC1S[1:0] bits (Capture/Compare 1 Selection).
	CC1S_0   CCMR1_Bits = 0x01 << 0  //  Bit 0.
	CC1S_1   CCMR1_Bits = 0x02 << 0  //  Bit 1.
	OC1FE    CCMR1_Bits = 0x01 << 2  //+ Output Compare 1 Fast enable.
	OC1PE    CCMR1_Bits = 0x01 << 3  //+ Output Compare 1 Preload enable.
	OC1M     CCMR1_Bits = 0x07 << 4  //+ OC1M[2:0] bits (Output Compare 1 Mode).
	OC1M_0   CCMR1_Bits = 0x01 << 4  //  Bit 0.
	OC1M_1   CCMR1_Bits = 0x02 << 4  //  Bit 1.
	OC1M_2   CCMR1_Bits = 0x04 << 4  //  Bit 2.
	OC1CE    CCMR1_Bits = 0x01 << 7  //+ Output Compare 1Clear Enable.
	CC2S     CCMR1_Bits = 0x03 << 8  //+ CC2S[1:0] bits (Capture/Compare 2 Selection).
	CC2S_0   CCMR1_Bits = 0x01 << 8  //  Bit 0.
	CC2S_1   CCMR1_Bits = 0x02 << 8  //  Bit 1.
	OC2FE    CCMR1_Bits = 0x01 << 10 //+ Output Compare 2 Fast enable.
	OC2PE    CCMR1_Bits = 0x01 << 11 //+ Output Compare 2 Preload enable.
	OC2M     CCMR1_Bits = 0x07 << 12 //+ OC2M[2:0] bits (Output Compare 2 Mode).
	OC2M_0   CCMR1_Bits = 0x01 << 12 //  Bit 0.
	OC2M_1   CCMR1_Bits = 0x02 << 12 //  Bit 1.
	OC2M_2   CCMR1_Bits = 0x04 << 12 //  Bit 2.
	OC2CE    CCMR1_Bits = 0x01 << 15 //+ Output Compare 2 Clear Enable.
	IC1PSC   CCMR1_Bits = 0x03 << 2  //  IC1PSC[1:0] bits (Input Capture 1 Prescaler).
	IC1PSC_0 CCMR1_Bits = 0x01 << 2  //  Bit 0.
	IC1PSC_1 CCMR1_Bits = 0x01 << 3  //  Bit 1.
	IC1F     CCMR1_Bits = 0x0F << 4  //  IC1F[3:0] bits (Input Capture 1 Filter).
	IC1F_0   CCMR1_Bits = 0x01 << 4  //  Bit 0.
	IC1F_1   CCMR1_Bits = 0x02 << 4  //  Bit 1.
	IC1F_2   CCMR1_Bits = 0x04 << 4  //  Bit 2.
	IC1F_3   CCMR1_Bits = 0x01 << 7  //  Bit 3.
	IC2PSC   CCMR1_Bits = 0x03 << 10 //  IC2PSC[1:0] bits (Input Capture 2 Prescaler).
	IC2PSC_0 CCMR1_Bits = 0x01 << 10 //  Bit 0.
	IC2PSC_1 CCMR1_Bits = 0x01 << 11 //  Bit 1.
	IC2F     CCMR1_Bits = 0x0F << 12 //  IC2F[3:0] bits (Input Capture 2 Filter).
	IC2F_0   CCMR1_Bits = 0x01 << 12 //  Bit 0.
	IC2F_1   CCMR1_Bits = 0x02 << 12 //  Bit 1.
	IC2F_2   CCMR1_Bits = 0x04 << 12 //  Bit 2.
	IC2F_3   CCMR1_Bits = 0x01 << 15 //  Bit 3.
)

const (
	CC3S     CCMR2_Bits = 0x03 << 0  //+ CC3S[1:0] bits (Capture/Compare 3 Selection).
	CC3S_0   CCMR2_Bits = 0x01 << 0  //  Bit 0.
	CC3S_1   CCMR2_Bits = 0x02 << 0  //  Bit 1.
	OC3FE    CCMR2_Bits = 0x01 << 2  //+ Output Compare 3 Fast enable.
	OC3PE    CCMR2_Bits = 0x01 << 3  //+ Output Compare 3 Preload enable.
	OC3M     CCMR2_Bits = 0x07 << 4  //+ OC3M[2:0] bits (Output Compare 3 Mode).
	OC3M_0   CCMR2_Bits = 0x01 << 4  //  Bit 0.
	OC3M_1   CCMR2_Bits = 0x02 << 4  //  Bit 1.
	OC3M_2   CCMR2_Bits = 0x04 << 4  //  Bit 2.
	OC3CE    CCMR2_Bits = 0x01 << 7  //+ Output Compare 3 Clear Enable.
	CC4S     CCMR2_Bits = 0x03 << 8  //+ CC4S[1:0] bits (Capture/Compare 4 Selection).
	CC4S_0   CCMR2_Bits = 0x01 << 8  //  Bit 0.
	CC4S_1   CCMR2_Bits = 0x02 << 8  //  Bit 1.
	OC4FE    CCMR2_Bits = 0x01 << 10 //+ Output Compare 4 Fast enable.
	OC4PE    CCMR2_Bits = 0x01 << 11 //+ Output Compare 4 Preload enable.
	OC4M     CCMR2_Bits = 0x07 << 12 //+ OC4M[2:0] bits (Output Compare 4 Mode).
	OC4M_0   CCMR2_Bits = 0x01 << 12 //  Bit 0.
	OC4M_1   CCMR2_Bits = 0x02 << 12 //  Bit 1.
	OC4M_2   CCMR2_Bits = 0x04 << 12 //  Bit 2.
	OC4CE    CCMR2_Bits = 0x01 << 15 //+ Output Compare 4 Clear Enable.
	IC3PSC   CCMR2_Bits = 0x03 << 2  //  IC3PSC[1:0] bits (Input Capture 3 Prescaler).
	IC3PSC_0 CCMR2_Bits = 0x01 << 2  //  Bit 0.
	IC3PSC_1 CCMR2_Bits = 0x01 << 3  //  Bit 1.
	IC3F     CCMR2_Bits = 0x0F << 4  //  IC3F[3:0] bits (Input Capture 3 Filter).
	IC3F_0   CCMR2_Bits = 0x01 << 4  //  Bit 0.
	IC3F_1   CCMR2_Bits = 0x02 << 4  //  Bit 1.
	IC3F_2   CCMR2_Bits = 0x04 << 4  //  Bit 2.
	IC3F_3   CCMR2_Bits = 0x01 << 7  //  Bit 3.
	IC4PSC   CCMR2_Bits = 0x03 << 10 //  IC4PSC[1:0] bits (Input Capture 4 Prescaler).
	IC4PSC_0 CCMR2_Bits = 0x01 << 10 //  Bit 0.
	IC4PSC_1 CCMR2_Bits = 0x01 << 11 //  Bit 1.
	IC4F     CCMR2_Bits = 0x0F << 12 //  IC4F[3:0] bits (Input Capture 4 Filter).
	IC4F_0   CCMR2_Bits = 0x01 << 12 //  Bit 0.
	IC4F_1   CCMR2_Bits = 0x02 << 12 //  Bit 1.
	IC4F_2   CCMR2_Bits = 0x04 << 12 //  Bit 2.
	IC4F_3   CCMR2_Bits = 0x01 << 15 //  Bit 3.
)

const (
	CC1E  CCER_Bits = 0x01 << 0  //+ Capture/Compare 1 output enable.
	CC1P  CCER_Bits = 0x01 << 1  //+ Capture/Compare 1 output Polarity.
	CC1NE CCER_Bits = 0x01 << 2  //+ Capture/Compare 1 Complementary output enable.
	CC1NP CCER_Bits = 0x01 << 3  //+ Capture/Compare 1 Complementary output Polarity.
	CC2E  CCER_Bits = 0x01 << 4  //+ Capture/Compare 2 output enable.
	CC2P  CCER_Bits = 0x01 << 5  //+ Capture/Compare 2 output Polarity.
	CC2NE CCER_Bits = 0x01 << 6  //+ Capture/Compare 2 Complementary output enable.
	CC2NP CCER_Bits = 0x01 << 7  //+ Capture/Compare 2 Complementary output Polarity.
	CC3E  CCER_Bits = 0x01 << 8  //+ Capture/Compare 3 output enable.
	CC3P  CCER_Bits = 0x01 << 9  //+ Capture/Compare 3 output Polarity.
	CC3NE CCER_Bits = 0x01 << 10 //+ Capture/Compare 3 Complementary output enable.
	CC3NP CCER_Bits = 0x01 << 11 //+ Capture/Compare 3 Complementary output Polarity.
	CC4E  CCER_Bits = 0x01 << 12 //+ Capture/Compare 4 output enable.
	CC4P  CCER_Bits = 0x01 << 13 //+ Capture/Compare 4 output Polarity.
	CC4NP CCER_Bits = 0x01 << 15 //+ Capture/Compare 4 Complementary output Polarity.
)

const ()

const ()

const ()

const (
	REP RCR_Bits = 0xFF << 0 //+ Repetition Counter Value.
)

const ()

const ()

const ()

const ()

const (
	DTG    BDTR_Bits = 0xFF << 0  //+ DTG[0:7] bits (Dead-Time Generator set-up).
	DTG_0  BDTR_Bits = 0x01 << 0  //  Bit 0.
	DTG_1  BDTR_Bits = 0x02 << 0  //  Bit 1.
	DTG_2  BDTR_Bits = 0x04 << 0  //  Bit 2.
	DTG_3  BDTR_Bits = 0x08 << 0  //  Bit 3.
	DTG_4  BDTR_Bits = 0x10 << 0  //  Bit 4.
	DTG_5  BDTR_Bits = 0x20 << 0  //  Bit 5.
	DTG_6  BDTR_Bits = 0x40 << 0  //  Bit 6.
	DTG_7  BDTR_Bits = 0x80 << 0  //  Bit 7.
	LOCK   BDTR_Bits = 0x03 << 8  //+ LOCK[1:0] bits (Lock Configuration).
	LOCK_0 BDTR_Bits = 0x01 << 8  //  Bit 0.
	LOCK_1 BDTR_Bits = 0x02 << 8  //  Bit 1.
	OSSI   BDTR_Bits = 0x01 << 10 //+ Off-State Selection for Idle mode.
	OSSR   BDTR_Bits = 0x01 << 11 //+ Off-State Selection for Run mode.
	BKE    BDTR_Bits = 0x01 << 12 //+ Break enable.
	BKP    BDTR_Bits = 0x01 << 13 //+ Break Polarity.
	AOE    BDTR_Bits = 0x01 << 14 //+ Automatic Output enable.
	MOE    BDTR_Bits = 0x01 << 15 //+ Main Output enable.
)

const (
	DBA   DCR_Bits = 0x1F << 0 //+ DBA[4:0] bits (DMA Base Address).
	DBA_0 DCR_Bits = 0x01 << 0 //  Bit 0.
	DBA_1 DCR_Bits = 0x02 << 0 //  Bit 1.
	DBA_2 DCR_Bits = 0x04 << 0 //  Bit 2.
	DBA_3 DCR_Bits = 0x08 << 0 //  Bit 3.
	DBA_4 DCR_Bits = 0x10 << 0 //  Bit 4.
	DBL   DCR_Bits = 0x1F << 8 //+ DBL[4:0] bits (DMA Burst Length).
	DBL_0 DCR_Bits = 0x01 << 8 //  Bit 0.
	DBL_1 DCR_Bits = 0x02 << 8 //  Bit 1.
	DBL_2 DCR_Bits = 0x04 << 8 //  Bit 2.
	DBL_3 DCR_Bits = 0x08 << 8 //  Bit 3.
	DBL_4 DCR_Bits = 0x10 << 8 //  Bit 4.
)

const (
	DMAB DMAR_Bits = 0xFFFF << 0 //+ DMA register for burst accesses.
)

const (
	TI4_RMP    OR_Bits = 0x03 << 6  //+ TI4_RMP[1:0] bits (TIM5 Input 4 remap).
	TI4_RMP_0  OR_Bits = 0x01 << 6  //  Bit 0.
	TI4_RMP_1  OR_Bits = 0x02 << 6  //  Bit 1.
	ITR1_RMP   OR_Bits = 0x03 << 10 //+ ITR1_RMP[1:0] bits (TIM2 Internal trigger 1 remap).
	ITR1_RMP_0 OR_Bits = 0x01 << 10 //  Bit 0.
	ITR1_RMP_1 OR_Bits = 0x02 << 10 //  Bit 1.
)
