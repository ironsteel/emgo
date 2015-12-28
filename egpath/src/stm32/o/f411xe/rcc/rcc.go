// +build f411xe

// Peripheral: RCC_Periph  Reset and Clock Control.
// Instances:
//  RCC  mmap.RCC_BASE
// Registers:
//  0x00 32  CR         Clock control register.
//  0x04 32  PLLCFGR    PLL configuration register.
//  0x08 32  CFGR       Clock configuration register.
//  0x0C 32  CIR        Clock interrupt register.
//  0x10 32  AHB1RSTR   AHB1 peripheral reset register.
//  0x14 32  AHB2RSTR   AHB2 peripheral reset register.
//  0x18 32  AHB3RSTR   AHB3 peripheral reset register.
//  0x20 32  APB1RSTR   APB1 peripheral reset register.
//  0x24 32  APB2RSTR   APB2 peripheral reset register.
//  0x30 32  AHB1ENR    AHB1 peripheral clock register.
//  0x34 32  AHB2ENR    AHB2 peripheral clock register.
//  0x38 32  AHB3ENR    AHB3 peripheral clock register.
//  0x40 32  APB1ENR    APB1 peripheral clock enable register.
//  0x44 32  APB2ENR    APB2 peripheral clock enable register.
//  0x50 32  AHB1LPENR  AHB1 peripheral clock enable in low power mode register.
//  0x54 32  AHB2LPENR  AHB2 peripheral clock enable in low power mode register.
//  0x58 32  AHB3LPENR  AHB3 peripheral clock enable in low power mode register.
//  0x60 32  APB1LPENR  APB1 peripheral clock enable in low power mode register.
//  0x64 32  APB2LPENR  APB2 peripheral clock enable in low power mode register.
//  0x70 32  BDCR       Backup domain control register.
//  0x74 32  CSR        Clock control & status register.
//  0x80 32  SSCGR      Spread spectrum clock generation register.
//  0x84 32  PLLI2SCFGR PLLI2S configuration register.
//  0x88 32  PLLSAICFGR PLLSAI configuration register.
//  0x8C 32  DCKCFGR    Dedicated Clocks configuration register.
//  0x90 32  CKGATENR   Clocks Gated Enable Register.
//  0x94 32  DCKCFGR2   Dedicated Clocks configuration register 2.
// Import:
//  stm32/o/f411xe/mmap
package rcc

const (
	HSION     CR_Bits = 0x01 << 0  //+
	HSIRDY    CR_Bits = 0x01 << 1  //+
	HSITRIM   CR_Bits = 0x1F << 3  //+
	HSITRIM_0 CR_Bits = 0x01 << 3  //  Bit 0.
	HSITRIM_1 CR_Bits = 0x02 << 3  //  Bit 1.
	HSITRIM_2 CR_Bits = 0x04 << 3  //  Bit 2.
	HSITRIM_3 CR_Bits = 0x08 << 3  //  Bit 3.
	HSITRIM_4 CR_Bits = 0x10 << 3  //  Bit 4.
	HSICAL    CR_Bits = 0xFF << 8  //+
	HSICAL_0  CR_Bits = 0x01 << 8  //  Bit 0.
	HSICAL_1  CR_Bits = 0x02 << 8  //  Bit 1.
	HSICAL_2  CR_Bits = 0x04 << 8  //  Bit 2.
	HSICAL_3  CR_Bits = 0x08 << 8  //  Bit 3.
	HSICAL_4  CR_Bits = 0x10 << 8  //  Bit 4.
	HSICAL_5  CR_Bits = 0x20 << 8  //  Bit 5.
	HSICAL_6  CR_Bits = 0x40 << 8  //  Bit 6.
	HSICAL_7  CR_Bits = 0x80 << 8  //  Bit 7.
	HSEON     CR_Bits = 0x01 << 16 //+
	HSERDY    CR_Bits = 0x01 << 17 //+
	HSEBYP    CR_Bits = 0x01 << 18 //+
	CSSON     CR_Bits = 0x01 << 19 //+
	PLLON     CR_Bits = 0x01 << 24 //+
	PLLRDY    CR_Bits = 0x01 << 25 //+
	PLLI2SON  CR_Bits = 0x01 << 26 //+
	PLLI2SRDY CR_Bits = 0x01 << 27 //+
	PLLSAION  CR_Bits = 0x01 << 28 //+
	PLLSAIRDY CR_Bits = 0x01 << 29 //+
)

const (
	PLLM       PLLCFGR_Bits = 0x3F << 0 //+
	PLLM_0     PLLCFGR_Bits = 0x01 << 0
	PLLM_1     PLLCFGR_Bits = 0x02 << 0
	PLLM_2     PLLCFGR_Bits = 0x04 << 0
	PLLM_3     PLLCFGR_Bits = 0x08 << 0
	PLLM_4     PLLCFGR_Bits = 0x10 << 0
	PLLM_5     PLLCFGR_Bits = 0x20 << 0
	PLLN       PLLCFGR_Bits = 0x1FF << 6 //+
	PLLN_0     PLLCFGR_Bits = 0x01 << 6
	PLLN_1     PLLCFGR_Bits = 0x02 << 6
	PLLN_2     PLLCFGR_Bits = 0x04 << 6
	PLLN_3     PLLCFGR_Bits = 0x08 << 6
	PLLN_4     PLLCFGR_Bits = 0x10 << 6
	PLLN_5     PLLCFGR_Bits = 0x20 << 6
	PLLN_6     PLLCFGR_Bits = 0x40 << 6
	PLLN_7     PLLCFGR_Bits = 0x80 << 6
	PLLN_8     PLLCFGR_Bits = 0x100 << 6
	PLLP       PLLCFGR_Bits = 0x03 << 16 //+
	PLLP_0     PLLCFGR_Bits = 0x01 << 16
	PLLP_1     PLLCFGR_Bits = 0x02 << 16
	PLLSRC     PLLCFGR_Bits = 0x01 << 22 //+
	PLLSRC_HSE PLLCFGR_Bits = 0x01 << 22
	PLLSRC_HSI PLLCFGR_Bits = 0x00 << 22
	PLLQ       PLLCFGR_Bits = 0x0F << 24 //+
	PLLQ_0     PLLCFGR_Bits = 0x01 << 24
	PLLQ_1     PLLCFGR_Bits = 0x02 << 24
	PLLQ_2     PLLCFGR_Bits = 0x04 << 24
	PLLQ_3     PLLCFGR_Bits = 0x08 << 24
)

const (
	SW          CFGR_Bits = 0x03 << 0  //+ SW[1:0] bits (System clock Switch).
	SW_0        CFGR_Bits = 0x01 << 0  //  Bit 0.
	SW_1        CFGR_Bits = 0x02 << 0  //  Bit 1.
	SW_HSI      CFGR_Bits = 0x00 << 0  //  HSI selected as system clock.
	SW_HSE      CFGR_Bits = 0x01 << 0  //  HSE selected as system clock.
	SW_PLL      CFGR_Bits = 0x02 << 0  //  PLL/PLLP selected as system clock.
	SWS         CFGR_Bits = 0x03 << 2  //+ SWS[1:0] bits (System Clock Switch Status).
	SWS_0       CFGR_Bits = 0x01 << 2  //  Bit 0.
	SWS_1       CFGR_Bits = 0x02 << 2  //  Bit 1.
	SWS_HSI     CFGR_Bits = 0x00 << 2  //  HSI oscillator used as system clock.
	SWS_HSE     CFGR_Bits = 0x01 << 2  //  HSE oscillator used as system clock.
	SWS_PLL     CFGR_Bits = 0x02 << 2  //  PLL/PLLP used as system clock.
	HPRE        CFGR_Bits = 0x0F << 4  //+ HPRE[3:0] bits (AHB prescaler).
	HPRE_0      CFGR_Bits = 0x01 << 4  //  Bit 0.
	HPRE_1      CFGR_Bits = 0x02 << 4  //  Bit 1.
	HPRE_2      CFGR_Bits = 0x04 << 4  //  Bit 2.
	HPRE_3      CFGR_Bits = 0x08 << 4  //  Bit 3.
	HPRE_DIV1   CFGR_Bits = 0x00 << 4  //  SYSCLK not divided.
	HPRE_DIV2   CFGR_Bits = 0x08 << 4  //  SYSCLK divided by 2.
	HPRE_DIV4   CFGR_Bits = 0x09 << 4  //  SYSCLK divided by 4.
	HPRE_DIV8   CFGR_Bits = 0x0A << 4  //  SYSCLK divided by 8.
	HPRE_DIV16  CFGR_Bits = 0x0B << 4  //  SYSCLK divided by 16.
	HPRE_DIV64  CFGR_Bits = 0x0C << 4  //  SYSCLK divided by 64.
	HPRE_DIV128 CFGR_Bits = 0x0D << 4  //  SYSCLK divided by 128.
	HPRE_DIV256 CFGR_Bits = 0x0E << 4  //  SYSCLK divided by 256.
	HPRE_DIV512 CFGR_Bits = 0x0F << 4  //  SYSCLK divided by 512.
	PPRE1       CFGR_Bits = 0x07 << 10 //+ PRE1[2:0] bits (APB1 prescaler).
	PPRE1_0     CFGR_Bits = 0x01 << 10 //  Bit 0.
	PPRE1_1     CFGR_Bits = 0x02 << 10 //  Bit 1.
	PPRE1_2     CFGR_Bits = 0x04 << 10 //  Bit 2.
	PPRE1_DIV1  CFGR_Bits = 0x00 << 10 //  HCLK not divided.
	PPRE1_DIV2  CFGR_Bits = 0x04 << 10 //  HCLK divided by 2.
	PPRE1_DIV4  CFGR_Bits = 0x05 << 10 //  HCLK divided by 4.
	PPRE1_DIV8  CFGR_Bits = 0x06 << 10 //  HCLK divided by 8.
	PPRE1_DIV16 CFGR_Bits = 0x07 << 10 //  HCLK divided by 16.
	PPRE2       CFGR_Bits = 0x07 << 13 //+ PRE2[2:0] bits (APB2 prescaler).
	PPRE2_0     CFGR_Bits = 0x01 << 13 //  Bit 0.
	PPRE2_1     CFGR_Bits = 0x02 << 13 //  Bit 1.
	PPRE2_2     CFGR_Bits = 0x04 << 13 //  Bit 2.
	PPRE2_DIV1  CFGR_Bits = 0x00 << 13 //  HCLK not divided.
	PPRE2_DIV2  CFGR_Bits = 0x04 << 13 //  HCLK divided by 2.
	PPRE2_DIV4  CFGR_Bits = 0x05 << 13 //  HCLK divided by 4.
	PPRE2_DIV8  CFGR_Bits = 0x06 << 13 //  HCLK divided by 8.
	PPRE2_DIV16 CFGR_Bits = 0x07 << 13 //  HCLK divided by 16.
	RTCPRE      CFGR_Bits = 0x1F << 16 //+
	RTCPRE_0    CFGR_Bits = 0x01 << 16
	RTCPRE_1    CFGR_Bits = 0x02 << 16
	RTCPRE_2    CFGR_Bits = 0x04 << 16
	RTCPRE_3    CFGR_Bits = 0x08 << 16
	RTCPRE_4    CFGR_Bits = 0x10 << 16
	MCO1        CFGR_Bits = 0x03 << 21 //+
	MCO1_0      CFGR_Bits = 0x01 << 21
	MCO1_1      CFGR_Bits = 0x02 << 21
	I2SSRC      CFGR_Bits = 0x01 << 23 //+
	MCO1PRE     CFGR_Bits = 0x07 << 24 //+
	MCO1PRE_0   CFGR_Bits = 0x01 << 24
	MCO1PRE_1   CFGR_Bits = 0x02 << 24
	MCO1PRE_2   CFGR_Bits = 0x04 << 24
	MCO2PRE     CFGR_Bits = 0x07 << 27 //+
	MCO2PRE_0   CFGR_Bits = 0x01 << 27
	MCO2PRE_1   CFGR_Bits = 0x02 << 27
	MCO2PRE_2   CFGR_Bits = 0x04 << 27
	MCO2        CFGR_Bits = 0x03 << 30 //+
	MCO2_0      CFGR_Bits = 0x01 << 30
	MCO2_1      CFGR_Bits = 0x02 << 30
)

const (
	LSIRDYF     CIR_Bits = 0x01 << 0  //+
	LSERDYF     CIR_Bits = 0x01 << 1  //+
	HSIRDYF     CIR_Bits = 0x01 << 2  //+
	HSERDYF     CIR_Bits = 0x01 << 3  //+
	PLLRDYF     CIR_Bits = 0x01 << 4  //+
	PLLI2SRDYF  CIR_Bits = 0x01 << 5  //+
	PLLSAIRDYF  CIR_Bits = 0x01 << 6  //+
	CSSF        CIR_Bits = 0x01 << 7  //+
	LSIRDYIE    CIR_Bits = 0x01 << 8  //+
	LSERDYIE    CIR_Bits = 0x01 << 9  //+
	HSIRDYIE    CIR_Bits = 0x01 << 10 //+
	HSERDYIE    CIR_Bits = 0x01 << 11 //+
	PLLRDYIE    CIR_Bits = 0x01 << 12 //+
	PLLI2SRDYIE CIR_Bits = 0x01 << 13 //+
	PLLSAIRDYIE CIR_Bits = 0x01 << 14 //+
	LSIRDYC     CIR_Bits = 0x01 << 16 //+
	LSERDYC     CIR_Bits = 0x01 << 17 //+
	HSIRDYC     CIR_Bits = 0x01 << 18 //+
	HSERDYC     CIR_Bits = 0x01 << 19 //+
	PLLRDYC     CIR_Bits = 0x01 << 20 //+
	PLLI2SRDYC  CIR_Bits = 0x01 << 21 //+
	PLLSAIRDYC  CIR_Bits = 0x01 << 22 //+
	CSSC        CIR_Bits = 0x01 << 23 //+
)

const (
	GPIOARST  AHB1RSTR_Bits = 0x01 << 0  //+
	GPIOBRST  AHB1RSTR_Bits = 0x01 << 1  //+
	GPIOCRST  AHB1RSTR_Bits = 0x01 << 2  //+
	GPIODRST  AHB1RSTR_Bits = 0x01 << 3  //+
	GPIOERST  AHB1RSTR_Bits = 0x01 << 4  //+
	GPIOFRST  AHB1RSTR_Bits = 0x01 << 5  //+
	GPIOGRST  AHB1RSTR_Bits = 0x01 << 6  //+
	GPIOHRST  AHB1RSTR_Bits = 0x01 << 7  //+
	GPIOIRST  AHB1RSTR_Bits = 0x01 << 8  //+
	GPIOJRST  AHB1RSTR_Bits = 0x01 << 9  //+
	GPIOKRST  AHB1RSTR_Bits = 0x01 << 10 //+
	CRCRST    AHB1RSTR_Bits = 0x01 << 12 //+
	DMA1RST   AHB1RSTR_Bits = 0x01 << 21 //+
	DMA2RST   AHB1RSTR_Bits = 0x01 << 22 //+
	DMA2DRST  AHB1RSTR_Bits = 0x01 << 23 //+
	ETHMACRST AHB1RSTR_Bits = 0x01 << 25 //+
	OTGHRST   AHB1RSTR_Bits = 0x01 << 28 //+
)

const (
	DCMIRST  AHB2RSTR_Bits = 0x01 << 0 //+
	CRYPRST  AHB2RSTR_Bits = 0x01 << 4 //+
	HASHRST  AHB2RSTR_Bits = 0x01 << 5 //+
	RNGRST   AHB2RSTR_Bits = 0x01 << 6 //+
	OTGFSRST AHB2RSTR_Bits = 0x01 << 7 //+
)

const (
	TIM2RST   APB1RSTR_Bits = 0x01 << 0  //+
	TIM3RST   APB1RSTR_Bits = 0x01 << 1  //+
	TIM4RST   APB1RSTR_Bits = 0x01 << 2  //+
	TIM5RST   APB1RSTR_Bits = 0x01 << 3  //+
	TIM6RST   APB1RSTR_Bits = 0x01 << 4  //+
	TIM7RST   APB1RSTR_Bits = 0x01 << 5  //+
	TIM12RST  APB1RSTR_Bits = 0x01 << 6  //+
	TIM13RST  APB1RSTR_Bits = 0x01 << 7  //+
	TIM14RST  APB1RSTR_Bits = 0x01 << 8  //+
	WWDGRST   APB1RSTR_Bits = 0x01 << 11 //+
	SPI2RST   APB1RSTR_Bits = 0x01 << 14 //+
	SPI3RST   APB1RSTR_Bits = 0x01 << 15 //+
	USART2RST APB1RSTR_Bits = 0x01 << 17 //+
	USART3RST APB1RSTR_Bits = 0x01 << 18 //+
	UART4RST  APB1RSTR_Bits = 0x01 << 19 //+
	UART5RST  APB1RSTR_Bits = 0x01 << 20 //+
	I2C1RST   APB1RSTR_Bits = 0x01 << 21 //+
	I2C2RST   APB1RSTR_Bits = 0x01 << 22 //+
	I2C3RST   APB1RSTR_Bits = 0x01 << 23 //+
	CAN1RST   APB1RSTR_Bits = 0x01 << 25 //+
	CAN2RST   APB1RSTR_Bits = 0x01 << 26 //+
	PWRRST    APB1RSTR_Bits = 0x01 << 28 //+
	DACRST    APB1RSTR_Bits = 0x01 << 29 //+
	UART7RST  APB1RSTR_Bits = 0x01 << 30 //+
	UART8RST  APB1RSTR_Bits = 0x01 << 31 //+
)

const (
	TIM1RST   APB2RSTR_Bits = 0x01 << 0  //+
	TIM8RST   APB2RSTR_Bits = 0x01 << 1  //+
	USART1RST APB2RSTR_Bits = 0x01 << 4  //+
	USART6RST APB2RSTR_Bits = 0x01 << 5  //+
	ADCRST    APB2RSTR_Bits = 0x01 << 8  //+
	SDIORST   APB2RSTR_Bits = 0x01 << 11 //+
	SPI1RST   APB2RSTR_Bits = 0x01 << 12 //+
	SPI4RST   APB2RSTR_Bits = 0x01 << 13 //+
	SYSCFGRST APB2RSTR_Bits = 0x01 << 14 //+
	TIM9RST   APB2RSTR_Bits = 0x01 << 16 //+
	TIM10RST  APB2RSTR_Bits = 0x01 << 17 //+
	TIM11RST  APB2RSTR_Bits = 0x01 << 18 //+
	SPI5RST   APB2RSTR_Bits = 0x01 << 20 //+
	SPI6RST   APB2RSTR_Bits = 0x01 << 21 //+
	SAI1RST   APB2RSTR_Bits = 0x01 << 22 //+
	LTDCRST   APB2RSTR_Bits = 0x01 << 26 //+
)

const (
	GPIOAEN      AHB1ENR_Bits = 0x01 << 0  //+
	GPIOBEN      AHB1ENR_Bits = 0x01 << 1  //+
	GPIOCEN      AHB1ENR_Bits = 0x01 << 2  //+
	GPIODEN      AHB1ENR_Bits = 0x01 << 3  //+
	GPIOEEN      AHB1ENR_Bits = 0x01 << 4  //+
	GPIOFEN      AHB1ENR_Bits = 0x01 << 5  //+
	GPIOGEN      AHB1ENR_Bits = 0x01 << 6  //+
	GPIOHEN      AHB1ENR_Bits = 0x01 << 7  //+
	GPIOIEN      AHB1ENR_Bits = 0x01 << 8  //+
	GPIOJEN      AHB1ENR_Bits = 0x01 << 9  //+
	GPIOKEN      AHB1ENR_Bits = 0x01 << 10 //+
	CRCEN        AHB1ENR_Bits = 0x01 << 12 //+
	BKPSRAMEN    AHB1ENR_Bits = 0x01 << 18 //+
	CCMDATARAMEN AHB1ENR_Bits = 0x01 << 20 //+
	DMA1EN       AHB1ENR_Bits = 0x01 << 21 //+
	DMA2EN       AHB1ENR_Bits = 0x01 << 22 //+
	DMA2DEN      AHB1ENR_Bits = 0x01 << 23 //+
	ETHMACEN     AHB1ENR_Bits = 0x01 << 25 //+
	ETHMACTXEN   AHB1ENR_Bits = 0x01 << 26 //+
	ETHMACRXEN   AHB1ENR_Bits = 0x01 << 27 //+
	ETHMACPTPEN  AHB1ENR_Bits = 0x01 << 28 //+
	OTGHSEN      AHB1ENR_Bits = 0x01 << 29 //+
	OTGHSULPIEN  AHB1ENR_Bits = 0x01 << 30 //+
)

const (
	DCMIEN  AHB2ENR_Bits = 0x01 << 0 //+
	CRYPEN  AHB2ENR_Bits = 0x01 << 4 //+
	HASHEN  AHB2ENR_Bits = 0x01 << 5 //+
	RNGEN   AHB2ENR_Bits = 0x01 << 6 //+
	OTGFSEN AHB2ENR_Bits = 0x01 << 7 //+
)

const (
	TIM2EN   APB1ENR_Bits = 0x01 << 0  //+
	TIM3EN   APB1ENR_Bits = 0x01 << 1  //+
	TIM4EN   APB1ENR_Bits = 0x01 << 2  //+
	TIM5EN   APB1ENR_Bits = 0x01 << 3  //+
	TIM6EN   APB1ENR_Bits = 0x01 << 4  //+
	TIM7EN   APB1ENR_Bits = 0x01 << 5  //+
	TIM12EN  APB1ENR_Bits = 0x01 << 6  //+
	TIM13EN  APB1ENR_Bits = 0x01 << 7  //+
	TIM14EN  APB1ENR_Bits = 0x01 << 8  //+
	WWDGEN   APB1ENR_Bits = 0x01 << 11 //+
	SPI2EN   APB1ENR_Bits = 0x01 << 14 //+
	SPI3EN   APB1ENR_Bits = 0x01 << 15 //+
	USART2EN APB1ENR_Bits = 0x01 << 17 //+
	USART3EN APB1ENR_Bits = 0x01 << 18 //+
	UART4EN  APB1ENR_Bits = 0x01 << 19 //+
	UART5EN  APB1ENR_Bits = 0x01 << 20 //+
	I2C1EN   APB1ENR_Bits = 0x01 << 21 //+
	I2C2EN   APB1ENR_Bits = 0x01 << 22 //+
	I2C3EN   APB1ENR_Bits = 0x01 << 23 //+
	CAN1EN   APB1ENR_Bits = 0x01 << 25 //+
	CAN2EN   APB1ENR_Bits = 0x01 << 26 //+
	PWREN    APB1ENR_Bits = 0x01 << 28 //+
	DACEN    APB1ENR_Bits = 0x01 << 29 //+
	UART7EN  APB1ENR_Bits = 0x01 << 30 //+
	UART8EN  APB1ENR_Bits = 0x01 << 31 //+
)

const (
	TIM1EN   APB2ENR_Bits = 0x01 << 0  //+
	TIM8EN   APB2ENR_Bits = 0x01 << 1  //+
	USART1EN APB2ENR_Bits = 0x01 << 4  //+
	USART6EN APB2ENR_Bits = 0x01 << 5  //+
	ADC1EN   APB2ENR_Bits = 0x01 << 8  //+
	ADC2EN   APB2ENR_Bits = 0x01 << 9  //+
	ADC3EN   APB2ENR_Bits = 0x01 << 10 //+
	SDIOEN   APB2ENR_Bits = 0x01 << 11 //+
	SPI1EN   APB2ENR_Bits = 0x01 << 12 //+
	SPI4EN   APB2ENR_Bits = 0x01 << 13 //+
	SYSCFGEN APB2ENR_Bits = 0x01 << 14 //+
	TIM9EN   APB2ENR_Bits = 0x01 << 16 //+
	TIM10EN  APB2ENR_Bits = 0x01 << 17 //+
	TIM11EN  APB2ENR_Bits = 0x01 << 18 //+
	SPI5EN   APB2ENR_Bits = 0x01 << 20 //+
	SPI6EN   APB2ENR_Bits = 0x01 << 21 //+
	SAI1EN   APB2ENR_Bits = 0x01 << 22 //+
	LTDCEN   APB2ENR_Bits = 0x01 << 26 //+
)

const (
	GPIOALPEN     AHB1LPENR_Bits = 0x01 << 0  //+
	GPIOBLPEN     AHB1LPENR_Bits = 0x01 << 1  //+
	GPIOCLPEN     AHB1LPENR_Bits = 0x01 << 2  //+
	GPIODLPEN     AHB1LPENR_Bits = 0x01 << 3  //+
	GPIOELPEN     AHB1LPENR_Bits = 0x01 << 4  //+
	GPIOFLPEN     AHB1LPENR_Bits = 0x01 << 5  //+
	GPIOGLPEN     AHB1LPENR_Bits = 0x01 << 6  //+
	GPIOHLPEN     AHB1LPENR_Bits = 0x01 << 7  //+
	GPIOILPEN     AHB1LPENR_Bits = 0x01 << 8  //+
	GPIOJLPEN     AHB1LPENR_Bits = 0x01 << 9  //+
	GPIOKLPEN     AHB1LPENR_Bits = 0x01 << 10 //+
	CRCLPEN       AHB1LPENR_Bits = 0x01 << 12 //+
	FLITFLPEN     AHB1LPENR_Bits = 0x01 << 15 //+
	SRAM1LPEN     AHB1LPENR_Bits = 0x01 << 16 //+
	SRAM2LPEN     AHB1LPENR_Bits = 0x01 << 17 //+
	BKPSRAMLPEN   AHB1LPENR_Bits = 0x01 << 18 //+
	SRAM3LPEN     AHB1LPENR_Bits = 0x01 << 19 //+
	DMA1LPEN      AHB1LPENR_Bits = 0x01 << 21 //+
	DMA2LPEN      AHB1LPENR_Bits = 0x01 << 22 //+
	DMA2DLPEN     AHB1LPENR_Bits = 0x01 << 23 //+
	ETHMACLPEN    AHB1LPENR_Bits = 0x01 << 25 //+
	ETHMACTXLPEN  AHB1LPENR_Bits = 0x01 << 26 //+
	ETHMACRXLPEN  AHB1LPENR_Bits = 0x01 << 27 //+
	ETHMACPTPLPEN AHB1LPENR_Bits = 0x01 << 28 //+
	OTGHSLPEN     AHB1LPENR_Bits = 0x01 << 29 //+
	OTGHSULPILPEN AHB1LPENR_Bits = 0x01 << 30 //+
)

const (
	DCMILPEN  AHB2LPENR_Bits = 0x01 << 0 //+
	CRYPLPEN  AHB2LPENR_Bits = 0x01 << 4 //+
	HASHLPEN  AHB2LPENR_Bits = 0x01 << 5 //+
	RNGLPEN   AHB2LPENR_Bits = 0x01 << 6 //+
	OTGFSLPEN AHB2LPENR_Bits = 0x01 << 7 //+
)

const (
	TIM2LPEN   APB1LPENR_Bits = 0x01 << 0  //+
	TIM3LPEN   APB1LPENR_Bits = 0x01 << 1  //+
	TIM4LPEN   APB1LPENR_Bits = 0x01 << 2  //+
	TIM5LPEN   APB1LPENR_Bits = 0x01 << 3  //+
	TIM6LPEN   APB1LPENR_Bits = 0x01 << 4  //+
	TIM7LPEN   APB1LPENR_Bits = 0x01 << 5  //+
	TIM12LPEN  APB1LPENR_Bits = 0x01 << 6  //+
	TIM13LPEN  APB1LPENR_Bits = 0x01 << 7  //+
	TIM14LPEN  APB1LPENR_Bits = 0x01 << 8  //+
	WWDGLPEN   APB1LPENR_Bits = 0x01 << 11 //+
	SPI2LPEN   APB1LPENR_Bits = 0x01 << 14 //+
	SPI3LPEN   APB1LPENR_Bits = 0x01 << 15 //+
	USART2LPEN APB1LPENR_Bits = 0x01 << 17 //+
	USART3LPEN APB1LPENR_Bits = 0x01 << 18 //+
	UART4LPEN  APB1LPENR_Bits = 0x01 << 19 //+
	UART5LPEN  APB1LPENR_Bits = 0x01 << 20 //+
	I2C1LPEN   APB1LPENR_Bits = 0x01 << 21 //+
	I2C2LPEN   APB1LPENR_Bits = 0x01 << 22 //+
	I2C3LPEN   APB1LPENR_Bits = 0x01 << 23 //+
	CAN1LPEN   APB1LPENR_Bits = 0x01 << 25 //+
	CAN2LPEN   APB1LPENR_Bits = 0x01 << 26 //+
	PWRLPEN    APB1LPENR_Bits = 0x01 << 28 //+
	DACLPEN    APB1LPENR_Bits = 0x01 << 29 //+
	UART7LPEN  APB1LPENR_Bits = 0x01 << 30 //+
	UART8LPEN  APB1LPENR_Bits = 0x01 << 31 //+
)

const (
	TIM1LPEN   APB2LPENR_Bits = 0x01 << 0  //+
	TIM8LPEN   APB2LPENR_Bits = 0x01 << 1  //+
	USART1LPEN APB2LPENR_Bits = 0x01 << 4  //+
	USART6LPEN APB2LPENR_Bits = 0x01 << 5  //+
	ADC1LPEN   APB2LPENR_Bits = 0x01 << 8  //+
	ADC2PEN    APB2LPENR_Bits = 0x01 << 9  //+
	ADC3LPEN   APB2LPENR_Bits = 0x01 << 10 //+
	SDIOLPEN   APB2LPENR_Bits = 0x01 << 11 //+
	SPI1LPEN   APB2LPENR_Bits = 0x01 << 12 //+
	SPI4LPEN   APB2LPENR_Bits = 0x01 << 13 //+
	SYSCFGLPEN APB2LPENR_Bits = 0x01 << 14 //+
	TIM9LPEN   APB2LPENR_Bits = 0x01 << 16 //+
	TIM10LPEN  APB2LPENR_Bits = 0x01 << 17 //+
	TIM11LPEN  APB2LPENR_Bits = 0x01 << 18 //+
	SPI5LPEN   APB2LPENR_Bits = 0x01 << 20 //+
	SPI6LPEN   APB2LPENR_Bits = 0x01 << 21 //+
	SAI1LPEN   APB2LPENR_Bits = 0x01 << 22 //+
	LTDCLPEN   APB2LPENR_Bits = 0x01 << 26 //+
)

const (
	LSEON    BDCR_Bits = 0x01 << 0 //+
	LSERDY   BDCR_Bits = 0x01 << 1 //+
	LSEBYP   BDCR_Bits = 0x01 << 2 //+
	LSEMOD   BDCR_Bits = 0x01 << 3 //+
	RTCSEL   BDCR_Bits = 0x03 << 8 //+
	RTCSEL_0 BDCR_Bits = 0x01 << 8
	RTCSEL_1 BDCR_Bits = 0x02 << 8
	RTCEN    BDCR_Bits = 0x01 << 15 //+
	BDRST    BDCR_Bits = 0x01 << 16 //+
)

const (
	LSION    CSR_Bits = 0x01 << 0  //+
	LSIRDY   CSR_Bits = 0x01 << 1  //+
	RMVF     CSR_Bits = 0x01 << 24 //+
	BORRSTF  CSR_Bits = 0x01 << 25 //+
	PADRSTF  CSR_Bits = 0x01 << 26 //+
	PORRSTF  CSR_Bits = 0x01 << 27 //+
	SFTRSTF  CSR_Bits = 0x01 << 28 //+
	WDGRSTF  CSR_Bits = 0x01 << 29 //+
	WWDGRSTF CSR_Bits = 0x01 << 30 //+
	LPWRRSTF CSR_Bits = 0x01 << 31 //+
)

const (
	MODPER    SSCGR_Bits = 0x1FFF << 0  //+
	INCSTEP   SSCGR_Bits = 0x7FFF << 13 //+
	SPREADSEL SSCGR_Bits = 0x01 << 30   //+
	SSCGEN    SSCGR_Bits = 0x01 << 31   //+
)

const (
	PLLI2SM   PLLI2SCFGR_Bits = 0x3F << 0 //+
	PLLI2SM_0 PLLI2SCFGR_Bits = 0x01 << 0
	PLLI2SM_1 PLLI2SCFGR_Bits = 0x02 << 0
	PLLI2SM_2 PLLI2SCFGR_Bits = 0x04 << 0
	PLLI2SM_3 PLLI2SCFGR_Bits = 0x08 << 0
	PLLI2SM_4 PLLI2SCFGR_Bits = 0x10 << 0
	PLLI2SM_5 PLLI2SCFGR_Bits = 0x20 << 0
	PLLI2SN   PLLI2SCFGR_Bits = 0x1FF << 6 //+
	PLLI2SN_0 PLLI2SCFGR_Bits = 0x01 << 6
	PLLI2SN_1 PLLI2SCFGR_Bits = 0x02 << 6
	PLLI2SN_2 PLLI2SCFGR_Bits = 0x04 << 6
	PLLI2SN_3 PLLI2SCFGR_Bits = 0x08 << 6
	PLLI2SN_4 PLLI2SCFGR_Bits = 0x10 << 6
	PLLI2SN_5 PLLI2SCFGR_Bits = 0x20 << 6
	PLLI2SN_6 PLLI2SCFGR_Bits = 0x40 << 6
	PLLI2SN_7 PLLI2SCFGR_Bits = 0x80 << 6
	PLLI2SN_8 PLLI2SCFGR_Bits = 0x100 << 6
	PLLI2SQ   PLLI2SCFGR_Bits = 0x0F << 24 //+
	PLLI2SQ_0 PLLI2SCFGR_Bits = 0x01 << 24
	PLLI2SQ_1 PLLI2SCFGR_Bits = 0x02 << 24
	PLLI2SQ_2 PLLI2SCFGR_Bits = 0x04 << 24
	PLLI2SQ_3 PLLI2SCFGR_Bits = 0x08 << 24
	PLLI2SR   PLLI2SCFGR_Bits = 0x07 << 28 //+
	PLLI2SR_0 PLLI2SCFGR_Bits = 0x01 << 28
	PLLI2SR_1 PLLI2SCFGR_Bits = 0x02 << 28
	PLLI2SR_2 PLLI2SCFGR_Bits = 0x04 << 28
)

const (
	PLLSAIN   PLLSAICFGR_Bits = 0x1FF << 6 //+
	PLLSAIN_0 PLLSAICFGR_Bits = 0x01 << 6
	PLLSAIN_1 PLLSAICFGR_Bits = 0x02 << 6
	PLLSAIN_2 PLLSAICFGR_Bits = 0x04 << 6
	PLLSAIN_3 PLLSAICFGR_Bits = 0x08 << 6
	PLLSAIN_4 PLLSAICFGR_Bits = 0x10 << 6
	PLLSAIN_5 PLLSAICFGR_Bits = 0x20 << 6
	PLLSAIN_6 PLLSAICFGR_Bits = 0x40 << 6
	PLLSAIN_7 PLLSAICFGR_Bits = 0x80 << 6
	PLLSAIN_8 PLLSAICFGR_Bits = 0x100 << 6
	PLLSAIQ   PLLSAICFGR_Bits = 0x0F << 24 //+
	PLLSAIQ_0 PLLSAICFGR_Bits = 0x01 << 24
	PLLSAIQ_1 PLLSAICFGR_Bits = 0x02 << 24
	PLLSAIQ_2 PLLSAICFGR_Bits = 0x04 << 24
	PLLSAIQ_3 PLLSAICFGR_Bits = 0x08 << 24
	PLLSAIR   PLLSAICFGR_Bits = 0x07 << 28 //+
	PLLSAIR_0 PLLSAICFGR_Bits = 0x01 << 28
	PLLSAIR_1 PLLSAICFGR_Bits = 0x02 << 28
	PLLSAIR_2 PLLSAICFGR_Bits = 0x04 << 28
)

const (
	PLLI2SDIVQ DCKCFGR_Bits = 0x1F << 0  //+
	PLLSAIDIVQ DCKCFGR_Bits = 0x1F << 8  //+
	PLLSAIDIVR DCKCFGR_Bits = 0x03 << 16 //+
	SAI1ASRC   DCKCFGR_Bits = 0x03 << 20 //+
	SAI1ASRC_0 DCKCFGR_Bits = 0x01 << 20
	SAI1ASRC_1 DCKCFGR_Bits = 0x02 << 20
	SAI1BSRC   DCKCFGR_Bits = 0x03 << 22 //+
	SAI1BSRC_0 DCKCFGR_Bits = 0x01 << 22
	SAI1BSRC_1 DCKCFGR_Bits = 0x02 << 22
	TIMPRE     DCKCFGR_Bits = 0x01 << 24 //+
)
