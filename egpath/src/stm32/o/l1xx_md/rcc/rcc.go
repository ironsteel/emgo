// +build l1xx_md

// Peripheral: RCC_Periph  Reset and Clock Control.
// Instances:
//  RCC  mmap.RCC_BASE
// Registers:
//  0x00 32  CR        Clock control register.
//  0x04 32  ICSCR     Internal clock sources calibration register.
//  0x08 32  CFGR      Clock configuration register.
//  0x0C 32  CIR       Clock interrupt register.
//  0x10 32  AHBRSTR   AHB peripheral reset register.
//  0x14 32  APB2RSTR  APB2 peripheral reset register.
//  0x18 32  APB1RSTR  APB1 peripheral reset register.
//  0x1C 32  AHBENR    AHB peripheral clock enable register.
//  0x20 32  APB2ENR   APB2 peripheral clock enable register.
//  0x24 32  APB1ENR   APB1 peripheral clock enable register.
//  0x28 32  AHBLPENR  AHB peripheral clock enable in low power mode register.
//  0x2C 32  APB2LPENR APB2 peripheral clock enable in low power mode register.
//  0x30 32  APB1LPENR APB1 peripheral clock enable in low power mode register.
//  0x34 32  CSR       Control/status register.
// Import:
//  stm32/o/l1xx_md/mmap
package rcc

const (
	HSION    CR_Bits = 0x01 << 0  //+ Internal High Speed clock enable.
	HSIRDY   CR_Bits = 0x01 << 1  //+ Internal High Speed clock ready flag.
	MSION    CR_Bits = 0x01 << 8  //+ Internal Multi Speed clock enable.
	MSIRDY   CR_Bits = 0x01 << 9  //+ Internal Multi Speed clock ready flag.
	HSEON    CR_Bits = 0x01 << 16 //+ External High Speed clock enable.
	HSERDY   CR_Bits = 0x01 << 17 //+ External High Speed clock ready flag.
	HSEBYP   CR_Bits = 0x01 << 18 //+ External High Speed clock Bypass.
	PLLON    CR_Bits = 0x01 << 24 //+ PLL enable.
	PLLRDY   CR_Bits = 0x01 << 25 //+ PLL clock ready flag.
	CSSON    CR_Bits = 0x01 << 28 //+ Clock Security System enable.
	RTCPRE   CR_Bits = 0x03 << 29 //+ RTC/LCD Prescaler.
	RTCPRE_0 CR_Bits = 0x01 << 29 //  Bit0.
	RTCPRE_1 CR_Bits = 0x02 << 29 //  Bit1.
)

const (
	HSICAL     ICSCR_Bits = 0xFF << 0  //+ Internal High Speed clock Calibration.
	HSITRIM    ICSCR_Bits = 0x1F << 8  //+ Internal High Speed clock trimming.
	MSIRANGE   ICSCR_Bits = 0x07 << 13 //+ Internal Multi Speed clock Range.
	MSIRANGE_0 ICSCR_Bits = 0x00 << 13 //  Internal Multi Speed clock Range 65.536 KHz.
	MSIRANGE_1 ICSCR_Bits = 0x01 << 13 //  Internal Multi Speed clock Range 131.072 KHz.
	MSIRANGE_2 ICSCR_Bits = 0x02 << 13 //  Internal Multi Speed clock Range 262.144 KHz.
	MSIRANGE_3 ICSCR_Bits = 0x03 << 13 //  Internal Multi Speed clock Range 524.288 KHz.
	MSIRANGE_4 ICSCR_Bits = 0x04 << 13 //  Internal Multi Speed clock Range 1.048 MHz.
	MSIRANGE_5 ICSCR_Bits = 0x05 << 13 //  Internal Multi Speed clock Range 2.097 MHz.
	MSIRANGE_6 ICSCR_Bits = 0x06 << 13 //  Internal Multi Speed clock Range 4.194 MHz.
	MSICAL     ICSCR_Bits = 0xFF << 16 //+ Internal Multi Speed clock Calibration.
	MSITRIM    ICSCR_Bits = 0xFF << 24 //+ Internal Multi Speed clock trimming.
)

const (
	SW          CFGR_Bits = 0x03 << 0  //+ SW[1:0] bits (System clock Switch).
	SW_0        CFGR_Bits = 0x01 << 0  //  Bit 0.
	SW_1        CFGR_Bits = 0x02 << 0  //  Bit 1.
	SW_MSI      CFGR_Bits = 0x00 << 0  //  MSI selected as system clock.
	SW_HSI      CFGR_Bits = 0x01 << 0  //  HSI selected as system clock.
	SW_HSE      CFGR_Bits = 0x02 << 0  //  HSE selected as system clock.
	SW_PLL      CFGR_Bits = 0x03 << 0  //  PLL selected as system clock.
	SWS         CFGR_Bits = 0x03 << 2  //+ SWS[1:0] bits (System Clock Switch Status).
	SWS_0       CFGR_Bits = 0x01 << 2  //  Bit 0.
	SWS_1       CFGR_Bits = 0x02 << 2  //  Bit 1.
	SWS_MSI     CFGR_Bits = 0x00 << 2  //  MSI oscillator used as system clock.
	SWS_HSI     CFGR_Bits = 0x01 << 2  //  HSI oscillator used as system clock.
	SWS_HSE     CFGR_Bits = 0x02 << 2  //  HSE oscillator used as system clock.
	SWS_PLL     CFGR_Bits = 0x03 << 2  //  PLL used as system clock.
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
	PPRE1       CFGR_Bits = 0x07 << 8  //+ PRE1[2:0] bits (APB1 prescaler).
	PPRE1_0     CFGR_Bits = 0x01 << 8  //  Bit 0.
	PPRE1_1     CFGR_Bits = 0x02 << 8  //  Bit 1.
	PPRE1_2     CFGR_Bits = 0x04 << 8  //  Bit 2.
	PPRE1_DIV1  CFGR_Bits = 0x00 << 8  //  HCLK not divided.
	PPRE1_DIV2  CFGR_Bits = 0x04 << 8  //  HCLK divided by 2.
	PPRE1_DIV4  CFGR_Bits = 0x05 << 8  //  HCLK divided by 4.
	PPRE1_DIV8  CFGR_Bits = 0x06 << 8  //  HCLK divided by 8.
	PPRE1_DIV16 CFGR_Bits = 0x07 << 8  //  HCLK divided by 16.
	PPRE2       CFGR_Bits = 0x07 << 11 //+ PRE2[2:0] bits (APB2 prescaler).
	PPRE2_0     CFGR_Bits = 0x01 << 11 //  Bit 0.
	PPRE2_1     CFGR_Bits = 0x02 << 11 //  Bit 1.
	PPRE2_2     CFGR_Bits = 0x04 << 11 //  Bit 2.
	PPRE2_DIV1  CFGR_Bits = 0x00 << 11 //  HCLK not divided.
	PPRE2_DIV2  CFGR_Bits = 0x04 << 11 //  HCLK divided by 2.
	PPRE2_DIV4  CFGR_Bits = 0x05 << 11 //  HCLK divided by 4.
	PPRE2_DIV8  CFGR_Bits = 0x06 << 11 //  HCLK divided by 8.
	PPRE2_DIV16 CFGR_Bits = 0x07 << 11 //  HCLK divided by 16.
	PLLSRC      CFGR_Bits = 0x01 << 16 //+ PLL entry clock source.
	PLLSRC_HSI  CFGR_Bits = 0x00 << 16 //  HSI as PLL entry clock source.
	PLLSRC_HSE  CFGR_Bits = 0x01 << 16 //  HSE as PLL entry clock source.
	PLLMUL      CFGR_Bits = 0x0F << 18 //+ PLLMUL[3:0] bits (PLL multiplication factor).
	PLLMUL_0    CFGR_Bits = 0x01 << 18 //  Bit 0.
	PLLMUL_1    CFGR_Bits = 0x02 << 18 //  Bit 1.
	PLLMUL_2    CFGR_Bits = 0x04 << 18 //  Bit 2.
	PLLMUL_3    CFGR_Bits = 0x08 << 18 //  Bit 3.
	PLLMUL3     CFGR_Bits = 0x00 << 18 //  PLL input clock * 3.
	PLLMUL4     CFGR_Bits = 0x01 << 18 //  PLL input clock * 4.
	PLLMUL6     CFGR_Bits = 0x02 << 18 //  PLL input clock * 6.
	PLLMUL8     CFGR_Bits = 0x03 << 18 //  PLL input clock * 8.
	PLLMUL12    CFGR_Bits = 0x04 << 18 //  PLL input clock * 12.
	PLLMUL16    CFGR_Bits = 0x05 << 18 //  PLL input clock * 16.
	PLLMUL24    CFGR_Bits = 0x06 << 18 //  PLL input clock * 24.
	PLLMUL32    CFGR_Bits = 0x07 << 18 //  PLL input clock * 32.
	PLLMUL48    CFGR_Bits = 0x08 << 18 //  PLL input clock * 48.
	PLLDIV      CFGR_Bits = 0x03 << 22 //+ PLLDIV[1:0] bits (PLL Output Division).
	PLLDIV_0    CFGR_Bits = 0x01 << 22 //  Bit0.
	PLLDIV_1    CFGR_Bits = 0x02 << 22 //  Bit1.
	PLLDIV1     CFGR_Bits = 0x00 << 22 //  PLL clock output = CKVCO / 1.
	PLLDIV2     CFGR_Bits = 0x01 << 22 //  PLL clock output = CKVCO / 2.
	PLLDIV3     CFGR_Bits = 0x02 << 22 //  PLL clock output = CKVCO / 3.
	PLLDIV4     CFGR_Bits = 0x03 << 22 //  PLL clock output = CKVCO / 4.
	MCOSEL      CFGR_Bits = 0x07 << 24 //+ MCO[2:0] bits (Microcontroller Clock Output).
	MCOSEL_0    CFGR_Bits = 0x01 << 24 //  Bit 0.
	MCOSEL_1    CFGR_Bits = 0x02 << 24 //  Bit 1.
	MCOSEL_2    CFGR_Bits = 0x04 << 24 //  Bit 2.
	MCO_NOCLOCK CFGR_Bits = 0x00 << 24 //  No clock.
	MCO_SYSCLK  CFGR_Bits = 0x01 << 24 //  System clock selected.
	MCO_HSI     CFGR_Bits = 0x02 << 24 //  Internal 16 MHz RC oscillator clock selected.
	MCO_MSI     CFGR_Bits = 0x03 << 24 //  Internal Medium Speed RC oscillator clock selected.
	MCO_HSE     CFGR_Bits = 0x04 << 24 //  External 1-25 MHz oscillator clock selected.
	MCO_PLL     CFGR_Bits = 0x05 << 24 //  PLL clock divided.
	MCO_LSI     CFGR_Bits = 0x06 << 24 //  LSI selected.
	MCO_LSE     CFGR_Bits = 0x07 << 24 //  LSE selected.
	MCOPRE      CFGR_Bits = 0x07 << 28 //+ MCOPRE[2:0] bits (Microcontroller Clock Output Prescaler).
	MCOPRE_0    CFGR_Bits = 0x01 << 28 //  Bit 0.
	MCOPRE_1    CFGR_Bits = 0x02 << 28 //  Bit 1.
	MCOPRE_2    CFGR_Bits = 0x04 << 28 //  Bit 2.
	MCO_DIV1    CFGR_Bits = 0x00 << 28 //  MCO Clock divided by 1.
	MCO_DIV2    CFGR_Bits = 0x01 << 28 //  MCO Clock divided by 2.
	MCO_DIV4    CFGR_Bits = 0x02 << 28 //  MCO Clock divided by 4.
	MCO_DIV8    CFGR_Bits = 0x03 << 28 //  MCO Clock divided by 8.
	MCO_DIV16   CFGR_Bits = 0x04 << 28 //  MCO Clock divided by 16.
)

const (
	LSIRDYF  CIR_Bits = 0x01 << 0  //+ LSI Ready Interrupt flag.
	LSERDYF  CIR_Bits = 0x01 << 1  //+ LSE Ready Interrupt flag.
	HSIRDYF  CIR_Bits = 0x01 << 2  //+ HSI Ready Interrupt flag.
	HSERDYF  CIR_Bits = 0x01 << 3  //+ HSE Ready Interrupt flag.
	PLLRDYF  CIR_Bits = 0x01 << 4  //+ PLL Ready Interrupt flag.
	MSIRDYF  CIR_Bits = 0x01 << 5  //+ MSI Ready Interrupt flag.
	LSECSS   CIR_Bits = 0x01 << 6  //+ LSE CSS Interrupt flag.
	CSSF     CIR_Bits = 0x01 << 7  //+ Clock Security System Interrupt flag.
	LSIRDYIE CIR_Bits = 0x01 << 8  //+ LSI Ready Interrupt Enable.
	LSERDYIE CIR_Bits = 0x01 << 9  //+ LSE Ready Interrupt Enable.
	HSIRDYIE CIR_Bits = 0x01 << 10 //+ HSI Ready Interrupt Enable.
	HSERDYIE CIR_Bits = 0x01 << 11 //+ HSE Ready Interrupt Enable.
	PLLRDYIE CIR_Bits = 0x01 << 12 //+ PLL Ready Interrupt Enable.
	MSIRDYIE CIR_Bits = 0x01 << 13 //+ MSI Ready Interrupt Enable.
	LSECSSIE CIR_Bits = 0x01 << 14 //+ LSE CSS Interrupt Enable.
	LSIRDYC  CIR_Bits = 0x01 << 16 //+ LSI Ready Interrupt Clear.
	LSERDYC  CIR_Bits = 0x01 << 17 //+ LSE Ready Interrupt Clear.
	HSIRDYC  CIR_Bits = 0x01 << 18 //+ HSI Ready Interrupt Clear.
	HSERDYC  CIR_Bits = 0x01 << 19 //+ HSE Ready Interrupt Clear.
	PLLRDYC  CIR_Bits = 0x01 << 20 //+ PLL Ready Interrupt Clear.
	MSIRDYC  CIR_Bits = 0x01 << 21 //+ MSI Ready Interrupt Clear.
	LSECSSC  CIR_Bits = 0x01 << 22 //+ LSE CSS Interrupt Clear.
	CSSC     CIR_Bits = 0x01 << 23 //+ Clock Security System Interrupt Clear.
)

const (
	GPIOARST AHBRSTR_Bits = 0x01 << 0  //+ GPIO port A reset.
	GPIOBRST AHBRSTR_Bits = 0x01 << 1  //+ GPIO port B reset.
	GPIOCRST AHBRSTR_Bits = 0x01 << 2  //+ GPIO port C reset.
	GPIODRST AHBRSTR_Bits = 0x01 << 3  //+ GPIO port D reset.
	GPIOERST AHBRSTR_Bits = 0x01 << 4  //+ GPIO port E reset.
	GPIOHRST AHBRSTR_Bits = 0x01 << 5  //+ GPIO port H reset.
	GPIOFRST AHBRSTR_Bits = 0x01 << 6  //+ GPIO port F reset.
	GPIOGRST AHBRSTR_Bits = 0x01 << 7  //+ GPIO port G reset.
	CRCRST   AHBRSTR_Bits = 0x01 << 12 //+ CRC reset.
	FLITFRST AHBRSTR_Bits = 0x01 << 15 //+ FLITF reset.
	DMA1RST  AHBRSTR_Bits = 0x01 << 24 //+ DMA1 reset.
	DMA2RST  AHBRSTR_Bits = 0x01 << 25 //+ DMA2 reset.
	AESRST   AHBRSTR_Bits = 0x01 << 27 //+ AES reset.
	FSMCRST  AHBRSTR_Bits = 0x01 << 30 //+ FSMC reset.
)

const (
	SYSCFGRST APB2RSTR_Bits = 0x01 << 0  //+ System Configuration SYSCFG reset.
	TIM9RST   APB2RSTR_Bits = 0x01 << 2  //+ TIM9 reset.
	TIM10RST  APB2RSTR_Bits = 0x01 << 3  //+ TIM10 reset.
	TIM11RST  APB2RSTR_Bits = 0x01 << 4  //+ TIM11 reset.
	ADC1RST   APB2RSTR_Bits = 0x01 << 9  //+ ADC1 reset.
	SDIORST   APB2RSTR_Bits = 0x01 << 11 //+ SDIO reset.
	SPI1RST   APB2RSTR_Bits = 0x01 << 12 //+ SPI1 reset.
	USART1RST APB2RSTR_Bits = 0x01 << 14 //+ USART1 reset.
)

const (
	TIM2RST   APB1RSTR_Bits = 0x01 << 0  //+ Timer 2 reset.
	TIM3RST   APB1RSTR_Bits = 0x01 << 1  //+ Timer 3 reset.
	TIM4RST   APB1RSTR_Bits = 0x01 << 2  //+ Timer 4 reset.
	TIM5RST   APB1RSTR_Bits = 0x01 << 3  //+ Timer 5 reset.
	TIM6RST   APB1RSTR_Bits = 0x01 << 4  //+ Timer 6 reset.
	TIM7RST   APB1RSTR_Bits = 0x01 << 5  //+ Timer 7 reset.
	LCDRST    APB1RSTR_Bits = 0x01 << 9  //+ LCD reset.
	WWDGRST   APB1RSTR_Bits = 0x01 << 11 //+ Window Watchdog reset.
	SPI2RST   APB1RSTR_Bits = 0x01 << 14 //+ SPI 2 reset.
	SPI3RST   APB1RSTR_Bits = 0x01 << 15 //+ SPI 3 reset.
	USART2RST APB1RSTR_Bits = 0x01 << 17 //+ USART 2 reset.
	USART3RST APB1RSTR_Bits = 0x01 << 18 //+ USART 3 reset.
	UART4RST  APB1RSTR_Bits = 0x01 << 19 //+ UART 4 reset.
	UART5RST  APB1RSTR_Bits = 0x01 << 20 //+ UART 5 reset.
	I2C1RST   APB1RSTR_Bits = 0x01 << 21 //+ I2C 1 reset.
	I2C2RST   APB1RSTR_Bits = 0x01 << 22 //+ I2C 2 reset.
	USBRST    APB1RSTR_Bits = 0x01 << 23 //+ USB reset.
	PWRRST    APB1RSTR_Bits = 0x01 << 28 //+ Power interface reset.
	DACRST    APB1RSTR_Bits = 0x01 << 29 //+ DAC interface reset.
	COMPRST   APB1RSTR_Bits = 0x01 << 31 //+ Comparator interface reset.
)

const (
	GPIOAEN AHBENR_Bits = 0x01 << 0  //+ GPIO port A clock enable.
	GPIOBEN AHBENR_Bits = 0x01 << 1  //+ GPIO port B clock enable.
	GPIOCEN AHBENR_Bits = 0x01 << 2  //+ GPIO port C clock enable.
	GPIODEN AHBENR_Bits = 0x01 << 3  //+ GPIO port D clock enable.
	GPIOEEN AHBENR_Bits = 0x01 << 4  //+ GPIO port E clock enable.
	GPIOHEN AHBENR_Bits = 0x01 << 5  //+ GPIO port H clock enable.
	GPIOFEN AHBENR_Bits = 0x01 << 6  //+ GPIO port F clock enable.
	GPIOGEN AHBENR_Bits = 0x01 << 7  //+ GPIO port G clock enable.
	CRCEN   AHBENR_Bits = 0x01 << 12 //+ CRC clock enable.
	FLITFEN AHBENR_Bits = 0x01 << 15 //+ FLITF clock enable (has effect only when.
	DMA1EN  AHBENR_Bits = 0x01 << 24 //+ DMA1 clock enable.
	DMA2EN  AHBENR_Bits = 0x01 << 25 //+ DMA2 clock enable.
	AESEN   AHBENR_Bits = 0x01 << 27 //+ AES clock enable.
	FSMCEN  AHBENR_Bits = 0x01 << 30 //+ FSMC clock enable.
)

const (
	SYSCFGEN APB2ENR_Bits = 0x01 << 0  //+ System Configuration SYSCFG clock enable.
	TIM9EN   APB2ENR_Bits = 0x01 << 2  //+ TIM9 interface clock enable.
	TIM10EN  APB2ENR_Bits = 0x01 << 3  //+ TIM10 interface clock enable.
	TIM11EN  APB2ENR_Bits = 0x01 << 4  //+ TIM11 Timer clock enable.
	ADC1EN   APB2ENR_Bits = 0x01 << 9  //+ ADC1 clock enable.
	SDIOEN   APB2ENR_Bits = 0x01 << 11 //+ SDIO clock enable.
	SPI1EN   APB2ENR_Bits = 0x01 << 12 //+ SPI1 clock enable.
	USART1EN APB2ENR_Bits = 0x01 << 14 //+ USART1 clock enable.
)

const (
	TIM2EN   APB1ENR_Bits = 0x01 << 0  //+ Timer 2 clock enabled.
	TIM3EN   APB1ENR_Bits = 0x01 << 1  //+ Timer 3 clock enable.
	TIM4EN   APB1ENR_Bits = 0x01 << 2  //+ Timer 4 clock enable.
	TIM5EN   APB1ENR_Bits = 0x01 << 3  //+ Timer 5 clock enable.
	TIM6EN   APB1ENR_Bits = 0x01 << 4  //+ Timer 6 clock enable.
	TIM7EN   APB1ENR_Bits = 0x01 << 5  //+ Timer 7 clock enable.
	LCDEN    APB1ENR_Bits = 0x01 << 9  //+ LCD clock enable.
	WWDGEN   APB1ENR_Bits = 0x01 << 11 //+ Window Watchdog clock enable.
	SPI2EN   APB1ENR_Bits = 0x01 << 14 //+ SPI 2 clock enable.
	SPI3EN   APB1ENR_Bits = 0x01 << 15 //+ SPI 3 clock enable.
	USART2EN APB1ENR_Bits = 0x01 << 17 //+ USART 2 clock enable.
	USART3EN APB1ENR_Bits = 0x01 << 18 //+ USART 3 clock enable.
	UART4EN  APB1ENR_Bits = 0x01 << 19 //+ UART 4 clock enable.
	UART5EN  APB1ENR_Bits = 0x01 << 20 //+ UART 5 clock enable.
	I2C1EN   APB1ENR_Bits = 0x01 << 21 //+ I2C 1 clock enable.
	I2C2EN   APB1ENR_Bits = 0x01 << 22 //+ I2C 2 clock enable.
	USBEN    APB1ENR_Bits = 0x01 << 23 //+ USB clock enable.
	PWREN    APB1ENR_Bits = 0x01 << 28 //+ Power interface clock enable.
	DACEN    APB1ENR_Bits = 0x01 << 29 //+ DAC interface clock enable.
	COMPEN   APB1ENR_Bits = 0x01 << 31 //+ Comparator interface clock enable.
)

const (
	GPIOALPEN AHBLPENR_Bits = 0x01 << 0  //+ GPIO port A clock enabled in sleep mode.
	GPIOBLPEN AHBLPENR_Bits = 0x01 << 1  //+ GPIO port B clock enabled in sleep mode.
	GPIOCLPEN AHBLPENR_Bits = 0x01 << 2  //+ GPIO port C clock enabled in sleep mode.
	GPIODLPEN AHBLPENR_Bits = 0x01 << 3  //+ GPIO port D clock enabled in sleep mode.
	GPIOELPEN AHBLPENR_Bits = 0x01 << 4  //+ GPIO port E clock enabled in sleep mode.
	GPIOHLPEN AHBLPENR_Bits = 0x01 << 5  //+ GPIO port H clock enabled in sleep mode.
	GPIOFLPEN AHBLPENR_Bits = 0x01 << 6  //+ GPIO port F clock enabled in sleep mode.
	GPIOGLPEN AHBLPENR_Bits = 0x01 << 7  //+ GPIO port G clock enabled in sleep mode.
	CRCLPEN   AHBLPENR_Bits = 0x01 << 12 //+ CRC clock enabled in sleep mode.
	FLITFLPEN AHBLPENR_Bits = 0x01 << 15 //+ Flash Interface clock enabled in sleep mode.
	SRAMLPEN  AHBLPENR_Bits = 0x01 << 16 //+ SRAM clock enabled in sleep mode.
	DMA1LPEN  AHBLPENR_Bits = 0x01 << 24 //+ DMA1 clock enabled in sleep mode.
	DMA2LPEN  AHBLPENR_Bits = 0x01 << 25 //+ DMA2 clock enabled in sleep mode.
	AESLPEN   AHBLPENR_Bits = 0x01 << 27 //+ AES clock enabled in sleep mode.
	FSMCLPEN  AHBLPENR_Bits = 0x01 << 30 //+ FSMC clock enabled in sleep mode.
)

const (
	SYSCFGLPEN APB2LPENR_Bits = 0x01 << 0  //+ System Configuration SYSCFG clock enabled in sleep mode.
	TIM9LPEN   APB2LPENR_Bits = 0x01 << 2  //+ TIM9 interface clock enabled in sleep mode.
	TIM10LPEN  APB2LPENR_Bits = 0x01 << 3  //+ TIM10 interface clock enabled in sleep mode.
	TIM11LPEN  APB2LPENR_Bits = 0x01 << 4  //+ TIM11 Timer clock enabled in sleep mode.
	ADC1LPEN   APB2LPENR_Bits = 0x01 << 9  //+ ADC1 clock enabled in sleep mode.
	SDIOLPEN   APB2LPENR_Bits = 0x01 << 11 //+ SDIO clock enabled in sleep mode.
	SPI1LPEN   APB2LPENR_Bits = 0x01 << 12 //+ SPI1 clock enabled in sleep mode.
	USART1LPEN APB2LPENR_Bits = 0x01 << 14 //+ USART1 clock enabled in sleep mode.
)

const (
	TIM2LPEN   APB1LPENR_Bits = 0x01 << 0  //+ Timer 2 clock enabled in sleep mode.
	TIM3LPEN   APB1LPENR_Bits = 0x01 << 1  //+ Timer 3 clock enabled in sleep mode.
	TIM4LPEN   APB1LPENR_Bits = 0x01 << 2  //+ Timer 4 clock enabled in sleep mode.
	TIM5LPEN   APB1LPENR_Bits = 0x01 << 3  //+ Timer 5 clock enabled in sleep mode.
	TIM6LPEN   APB1LPENR_Bits = 0x01 << 4  //+ Timer 6 clock enabled in sleep mode.
	TIM7LPEN   APB1LPENR_Bits = 0x01 << 5  //+ Timer 7 clock enabled in sleep mode.
	LCDLPEN    APB1LPENR_Bits = 0x01 << 9  //+ LCD clock enabled in sleep mode.
	WWDGLPEN   APB1LPENR_Bits = 0x01 << 11 //+ Window Watchdog clock enabled in sleep mode.
	SPI2LPEN   APB1LPENR_Bits = 0x01 << 14 //+ SPI 2 clock enabled in sleep mode.
	SPI3LPEN   APB1LPENR_Bits = 0x01 << 15 //+ SPI 3 clock enabled in sleep mode.
	USART2LPEN APB1LPENR_Bits = 0x01 << 17 //+ USART 2 clock enabled in sleep mode.
	USART3LPEN APB1LPENR_Bits = 0x01 << 18 //+ USART 3 clock enabled in sleep mode.
	UART4LPEN  APB1LPENR_Bits = 0x01 << 19 //+ UART 4 clock enabled in sleep mode.
	UART5LPEN  APB1LPENR_Bits = 0x01 << 20 //+ UART 5 clock enabled in sleep mode.
	I2C1LPEN   APB1LPENR_Bits = 0x01 << 21 //+ I2C 1 clock enabled in sleep mode.
	I2C2LPEN   APB1LPENR_Bits = 0x01 << 22 //+ I2C 2 clock enabled in sleep mode.
	USBLPEN    APB1LPENR_Bits = 0x01 << 23 //+ USB clock enabled in sleep mode.
	PWRLPEN    APB1LPENR_Bits = 0x01 << 28 //+ Power interface clock enabled in sleep mode.
	DACLPEN    APB1LPENR_Bits = 0x01 << 29 //+ DAC interface clock enabled in sleep mode.
	COMPLPEN   APB1LPENR_Bits = 0x01 << 31 //+ Comparator interface clock enabled in sleep mode.
)

const (
	LSION          CSR_Bits = 0x01 << 0  //+ Internal Low Speed oscillator enable.
	LSIRDY         CSR_Bits = 0x01 << 1  //+ Internal Low Speed oscillator Ready.
	LSEON          CSR_Bits = 0x01 << 8  //+ External Low Speed oscillator enable.
	LSERDY         CSR_Bits = 0x01 << 9  //+ External Low Speed oscillator Ready.
	LSEBYP         CSR_Bits = 0x01 << 10 //+ External Low Speed oscillator Bypass.
	LSECSSON       CSR_Bits = 0x01 << 11 //+ External Low Speed oscillator CSS Enable.
	LSECSSD        CSR_Bits = 0x01 << 12 //+ External Low Speed oscillator CSS Detected.
	RTCSEL         CSR_Bits = 0x03 << 16 //+ RTCSEL[1:0] bits (RTC clock source selection).
	RTCSEL_0       CSR_Bits = 0x01 << 16 //  Bit 0.
	RTCSEL_1       CSR_Bits = 0x02 << 16 //  Bit 1.
	RTCSEL_NOCLOCK CSR_Bits = 0x00 << 16 //  No clock.
	RTCSEL_LSE     CSR_Bits = 0x01 << 16 //  LSE oscillator clock used as RTC clock.
	RTCSEL_LSI     CSR_Bits = 0x02 << 16 //  LSI oscillator clock used as RTC clock.
	RTCSEL_HSE     CSR_Bits = 0x03 << 16 //  HSE oscillator clock divided by 2, 4, 8 or 16 by RTCPRE used as RTC clock.
	RTCEN          CSR_Bits = 0x01 << 22 //+ RTC clock enable.
	RTCRST         CSR_Bits = 0x01 << 23 //+ RTC reset.
	RMVF           CSR_Bits = 0x01 << 24 //+ Remove reset flag.
	OBLRSTF        CSR_Bits = 0x01 << 25 //+ Option Bytes Loader reset flag.
	PINRSTF        CSR_Bits = 0x01 << 26 //+ PIN reset flag.
	PORRSTF        CSR_Bits = 0x01 << 27 //+ POR/PDR reset flag.
	SFTRSTF        CSR_Bits = 0x01 << 28 //+ Software Reset flag.
	IWDGRSTF       CSR_Bits = 0x01 << 29 //+ Independent Watchdog reset flag.
	WWDGRSTF       CSR_Bits = 0x01 << 30 //+ Window watchdog reset flag.
	LPWRRSTF       CSR_Bits = 0x01 << 31 //+ Low-Power reset flag.
)
