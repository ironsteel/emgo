// +build f10x_hd

// Peripheral: RCC_Periph  Reset and Clock Control.
// Instances:
//  RCC  mmap.RCC_BASE
// Registers:
//  0x00 32  CR
//  0x04 32  CFGR
//  0x08 32  CIR
//  0x0C 32  APB2RSTR
//  0x10 32  APB1RSTR
//  0x14 32  AHBENR
//  0x18 32  APB2ENR
//  0x1C 32  APB1ENR
//  0x20 32  BDCR
//  0x24 32  CSR
// Import:
//  stm32/o/f10x_hd/mmap
package rcc

const (
	HSION   CR_Bits = 0x01 << 0  //+ Internal High Speed clock enable.
	HSIRDY  CR_Bits = 0x01 << 1  //+ Internal High Speed clock ready flag.
	HSITRIM CR_Bits = 0x1F << 3  //+ Internal High Speed clock trimming.
	HSICAL  CR_Bits = 0xFF << 8  //+ Internal High Speed clock Calibration.
	HSEON   CR_Bits = 0x01 << 16 //+ External High Speed clock enable.
	HSERDY  CR_Bits = 0x01 << 17 //+ External High Speed clock ready flag.
	HSEBYP  CR_Bits = 0x01 << 18 //+ External High Speed clock Bypass.
	CSSON   CR_Bits = 0x01 << 19 //+ Clock Security System enable.
	PLLON   CR_Bits = 0x01 << 24 //+ PLL enable.
	PLLRDY  CR_Bits = 0x01 << 25 //+ PLL clock ready flag.
)

const (
	SW                CFGR_Bits = 0x03 << 0  //+ SW[1:0] bits (System clock Switch).
	SW_0              CFGR_Bits = 0x01 << 0  //  Bit 0.
	SW_1              CFGR_Bits = 0x02 << 0  //  Bit 1.
	SW_HSI            CFGR_Bits = 0x00 << 0  //  HSI selected as system clock.
	SW_HSE            CFGR_Bits = 0x01 << 0  //  HSE selected as system clock.
	SW_PLL            CFGR_Bits = 0x02 << 0  //  PLL selected as system clock.
	SWS               CFGR_Bits = 0x03 << 2  //+ SWS[1:0] bits (System Clock Switch Status).
	SWS_0             CFGR_Bits = 0x01 << 2  //  Bit 0.
	SWS_1             CFGR_Bits = 0x02 << 2  //  Bit 1.
	SWS_HSI           CFGR_Bits = 0x00 << 2  //  HSI oscillator used as system clock.
	SWS_HSE           CFGR_Bits = 0x01 << 2  //  HSE oscillator used as system clock.
	SWS_PLL           CFGR_Bits = 0x02 << 2  //  PLL used as system clock.
	HPRE              CFGR_Bits = 0x0F << 4  //+ HPRE[3:0] bits (AHB prescaler).
	HPRE_0            CFGR_Bits = 0x01 << 4  //  Bit 0.
	HPRE_1            CFGR_Bits = 0x02 << 4  //  Bit 1.
	HPRE_2            CFGR_Bits = 0x04 << 4  //  Bit 2.
	HPRE_3            CFGR_Bits = 0x08 << 4  //  Bit 3.
	HPRE_DIV1         CFGR_Bits = 0x00 << 4  //  SYSCLK not divided.
	HPRE_DIV2         CFGR_Bits = 0x08 << 4  //  SYSCLK divided by 2.
	HPRE_DIV4         CFGR_Bits = 0x09 << 4  //  SYSCLK divided by 4.
	HPRE_DIV8         CFGR_Bits = 0x0A << 4  //  SYSCLK divided by 8.
	HPRE_DIV16        CFGR_Bits = 0x0B << 4  //  SYSCLK divided by 16.
	HPRE_DIV64        CFGR_Bits = 0x0C << 4  //  SYSCLK divided by 64.
	HPRE_DIV128       CFGR_Bits = 0x0D << 4  //  SYSCLK divided by 128.
	HPRE_DIV256       CFGR_Bits = 0x0E << 4  //  SYSCLK divided by 256.
	HPRE_DIV512       CFGR_Bits = 0x0F << 4  //  SYSCLK divided by 512.
	PPRE1             CFGR_Bits = 0x07 << 8  //+ PRE1[2:0] bits (APB1 prescaler).
	PPRE1_0           CFGR_Bits = 0x01 << 8  //  Bit 0.
	PPRE1_1           CFGR_Bits = 0x02 << 8  //  Bit 1.
	PPRE1_2           CFGR_Bits = 0x04 << 8  //  Bit 2.
	PPRE1_DIV1        CFGR_Bits = 0x00 << 8  //  HCLK not divided.
	PPRE1_DIV2        CFGR_Bits = 0x04 << 8  //  HCLK divided by 2.
	PPRE1_DIV4        CFGR_Bits = 0x05 << 8  //  HCLK divided by 4.
	PPRE1_DIV8        CFGR_Bits = 0x06 << 8  //  HCLK divided by 8.
	PPRE1_DIV16       CFGR_Bits = 0x07 << 8  //  HCLK divided by 16.
	PPRE2             CFGR_Bits = 0x07 << 11 //+ PRE2[2:0] bits (APB2 prescaler).
	PPRE2_0           CFGR_Bits = 0x01 << 11 //  Bit 0.
	PPRE2_1           CFGR_Bits = 0x02 << 11 //  Bit 1.
	PPRE2_2           CFGR_Bits = 0x04 << 11 //  Bit 2.
	PPRE2_DIV1        CFGR_Bits = 0x00 << 11 //  HCLK not divided.
	PPRE2_DIV2        CFGR_Bits = 0x04 << 11 //  HCLK divided by 2.
	PPRE2_DIV4        CFGR_Bits = 0x05 << 11 //  HCLK divided by 4.
	PPRE2_DIV8        CFGR_Bits = 0x06 << 11 //  HCLK divided by 8.
	PPRE2_DIV16       CFGR_Bits = 0x07 << 11 //  HCLK divided by 16.
	ADCPRE            CFGR_Bits = 0x03 << 14 //+ ADCPRE[1:0] bits (ADC prescaler).
	ADCPRE_0          CFGR_Bits = 0x01 << 14 //  Bit 0.
	ADCPRE_1          CFGR_Bits = 0x02 << 14 //  Bit 1.
	ADCPRE_DIV2       CFGR_Bits = 0x00 << 14 //  PCLK2 divided by 2.
	ADCPRE_DIV4       CFGR_Bits = 0x01 << 14 //  PCLK2 divided by 4.
	ADCPRE_DIV6       CFGR_Bits = 0x02 << 14 //  PCLK2 divided by 6.
	ADCPRE_DIV8       CFGR_Bits = 0x03 << 14 //  PCLK2 divided by 8.
	PLLSRC            CFGR_Bits = 0x01 << 16 //+ PLL entry clock source.
	PLLXTPRE          CFGR_Bits = 0x01 << 17 //+ HSE divider for PLL entry.
	PLLMULL           CFGR_Bits = 0x0F << 18 //+ PLLMUL[3:0] bits (PLL multiplication factor).
	PLLMULL_0         CFGR_Bits = 0x01 << 18 //  Bit 0.
	PLLMULL_1         CFGR_Bits = 0x02 << 18 //  Bit 1.
	PLLMULL_2         CFGR_Bits = 0x04 << 18 //  Bit 2.
	PLLMULL_3         CFGR_Bits = 0x08 << 18 //  Bit 3.
	PLLSRC_HSI_Div2   CFGR_Bits = 0x00 << 18 //  HSI clock divided by 2 selected as PLL entry clock source.
	PLLSRC_HSE        CFGR_Bits = 0x01 << 16 //  HSE clock selected as PLL entry clock source.
	PLLXTPRE_HSE      CFGR_Bits = 0x00 << 18 //  HSE clock not divided for PLL entry.
	PLLXTPRE_HSE_Div2 CFGR_Bits = 0x01 << 17 //  HSE clock divided by 2 for PLL entry.
	PLLMULL2          CFGR_Bits = 0x00 << 18 //  PLL input clock*2.
	PLLMULL3          CFGR_Bits = 0x01 << 18 //  PLL input clock*3.
	PLLMULL4          CFGR_Bits = 0x02 << 18 //  PLL input clock*4.
	PLLMULL5          CFGR_Bits = 0x03 << 18 //  PLL input clock*5.
	PLLMULL6          CFGR_Bits = 0x04 << 18 //  PLL input clock*6.
	PLLMULL7          CFGR_Bits = 0x05 << 18 //  PLL input clock*7.
	PLLMULL8          CFGR_Bits = 0x06 << 18 //  PLL input clock*8.
	PLLMULL9          CFGR_Bits = 0x07 << 18 //  PLL input clock*9.
	PLLMULL10         CFGR_Bits = 0x08 << 18 //  PLL input clock10.
	PLLMULL11         CFGR_Bits = 0x09 << 18 //  PLL input clock*11.
	PLLMULL12         CFGR_Bits = 0x0A << 18 //  PLL input clock*12.
	PLLMULL13         CFGR_Bits = 0x0B << 18 //  PLL input clock*13.
	PLLMULL14         CFGR_Bits = 0x0C << 18 //  PLL input clock*14.
	PLLMULL15         CFGR_Bits = 0x0D << 18 //  PLL input clock*15.
	PLLMULL16         CFGR_Bits = 0x0E << 18 //  PLL input clock*16.
	USBPRE            CFGR_Bits = 0x01 << 22 //+ USB Device prescaler.
	MCO               CFGR_Bits = 0x07 << 24 //+ MCO[2:0] bits (Microcontroller Clock Output).
	MCO_0             CFGR_Bits = 0x01 << 24 //  Bit 0.
	MCO_1             CFGR_Bits = 0x02 << 24 //  Bit 1.
	MCO_2             CFGR_Bits = 0x04 << 24 //  Bit 2.
	MCO_NOCLOCK       CFGR_Bits = 0x00 << 24 //  No clock.
	MCO_SYSCLK        CFGR_Bits = 0x04 << 24 //  System clock selected as MCO source.
	MCO_HSI           CFGR_Bits = 0x05 << 24 //  HSI clock selected as MCO source.
	MCO_HSE           CFGR_Bits = 0x06 << 24 //  HSE clock selected as MCO source.
	MCO_PLL           CFGR_Bits = 0x07 << 24 //  PLL clock divided by 2 selected as MCO source.
)

const (
	LSIRDYF  CIR_Bits = 0x01 << 0  //+ LSI Ready Interrupt flag.
	LSERDYF  CIR_Bits = 0x01 << 1  //+ LSE Ready Interrupt flag.
	HSIRDYF  CIR_Bits = 0x01 << 2  //+ HSI Ready Interrupt flag.
	HSERDYF  CIR_Bits = 0x01 << 3  //+ HSE Ready Interrupt flag.
	PLLRDYF  CIR_Bits = 0x01 << 4  //+ PLL Ready Interrupt flag.
	CSSF     CIR_Bits = 0x01 << 7  //+ Clock Security System Interrupt flag.
	LSIRDYIE CIR_Bits = 0x01 << 8  //+ LSI Ready Interrupt Enable.
	LSERDYIE CIR_Bits = 0x01 << 9  //+ LSE Ready Interrupt Enable.
	HSIRDYIE CIR_Bits = 0x01 << 10 //+ HSI Ready Interrupt Enable.
	HSERDYIE CIR_Bits = 0x01 << 11 //+ HSE Ready Interrupt Enable.
	PLLRDYIE CIR_Bits = 0x01 << 12 //+ PLL Ready Interrupt Enable.
	LSIRDYC  CIR_Bits = 0x01 << 16 //+ LSI Ready Interrupt Clear.
	LSERDYC  CIR_Bits = 0x01 << 17 //+ LSE Ready Interrupt Clear.
	HSIRDYC  CIR_Bits = 0x01 << 18 //+ HSI Ready Interrupt Clear.
	HSERDYC  CIR_Bits = 0x01 << 19 //+ HSE Ready Interrupt Clear.
	PLLRDYC  CIR_Bits = 0x01 << 20 //+ PLL Ready Interrupt Clear.
	CSSC     CIR_Bits = 0x01 << 23 //+ Clock Security System Interrupt Clear.
)

const (
	AFIORST   APB2RSTR_Bits = 0x01 << 0  //+ Alternate Function I/O reset.
	IOPARST   APB2RSTR_Bits = 0x01 << 2  //+ I/O port A reset.
	IOPBRST   APB2RSTR_Bits = 0x01 << 3  //+ I/O port B reset.
	IOPCRST   APB2RSTR_Bits = 0x01 << 4  //+ I/O port C reset.
	IOPDRST   APB2RSTR_Bits = 0x01 << 5  //+ I/O port D reset.
	ADC1RST   APB2RSTR_Bits = 0x01 << 9  //+ ADC 1 interface reset.
	ADC2RST   APB2RSTR_Bits = 0x01 << 10 //+ ADC 2 interface reset.
	TIM1RST   APB2RSTR_Bits = 0x01 << 11 //+ TIM1 Timer reset.
	SPI1RST   APB2RSTR_Bits = 0x01 << 12 //+ SPI 1 reset.
	USART1RST APB2RSTR_Bits = 0x01 << 14 //+ USART1 reset.
	IOPERST   APB2RSTR_Bits = 0x01 << 6  //+ I/O port E reset.
	IOPFRST   APB2RSTR_Bits = 0x01 << 7  //+ I/O port F reset.
	IOPGRST   APB2RSTR_Bits = 0x01 << 8  //+ I/O port G reset.
	TIM8RST   APB2RSTR_Bits = 0x01 << 13 //+ TIM8 Timer reset.
	ADC3RST   APB2RSTR_Bits = 0x01 << 15 //+ ADC3 interface reset.
)

const (
	TIM2RST   APB1RSTR_Bits = 0x01 << 0  //+ Timer 2 reset.
	TIM3RST   APB1RSTR_Bits = 0x01 << 1  //+ Timer 3 reset.
	WWDGRST   APB1RSTR_Bits = 0x01 << 11 //+ Window Watchdog reset.
	USART2RST APB1RSTR_Bits = 0x01 << 17 //+ USART 2 reset.
	I2C1RST   APB1RSTR_Bits = 0x01 << 21 //+ I2C 1 reset.
	CAN1RST   APB1RSTR_Bits = 0x01 << 25 //+ CAN1 reset.
	BKPRST    APB1RSTR_Bits = 0x01 << 27 //+ Backup interface reset.
	PWRRST    APB1RSTR_Bits = 0x01 << 28 //+ Power interface reset.
	TIM4RST   APB1RSTR_Bits = 0x01 << 2  //+ Timer 4 reset.
	SPI2RST   APB1RSTR_Bits = 0x01 << 14 //+ SPI 2 reset.
	USART3RST APB1RSTR_Bits = 0x01 << 18 //+ USART 3 reset.
	I2C2RST   APB1RSTR_Bits = 0x01 << 22 //+ I2C 2 reset.
	USBRST    APB1RSTR_Bits = 0x01 << 23 //+ USB Device reset.
	TIM5RST   APB1RSTR_Bits = 0x01 << 3  //+ Timer 5 reset.
	TIM6RST   APB1RSTR_Bits = 0x01 << 4  //+ Timer 6 reset.
	TIM7RST   APB1RSTR_Bits = 0x01 << 5  //+ Timer 7 reset.
	SPI3RST   APB1RSTR_Bits = 0x01 << 15 //+ SPI 3 reset.
	UART4RST  APB1RSTR_Bits = 0x01 << 19 //+ UART 4 reset.
	UART5RST  APB1RSTR_Bits = 0x01 << 20 //+ UART 5 reset.
	DACRST    APB1RSTR_Bits = 0x01 << 29 //+ DAC interface reset.
)

const (
	DMA1EN  AHBENR_Bits = 0x01 << 0  //+ DMA1 clock enable.
	SRAMEN  AHBENR_Bits = 0x01 << 2  //+ SRAM interface clock enable.
	FLITFEN AHBENR_Bits = 0x01 << 4  //+ FLITF clock enable.
	CRCEN   AHBENR_Bits = 0x01 << 6  //+ CRC clock enable.
	DMA2EN  AHBENR_Bits = 0x01 << 1  //+ DMA2 clock enable.
	FSMCEN  AHBENR_Bits = 0x01 << 8  //+ FSMC clock enable.
	SDIOEN  AHBENR_Bits = 0x01 << 10 //+ SDIO clock enable.
)

const (
	AFIOEN   APB2ENR_Bits = 0x01 << 0  //+ Alternate Function I/O clock enable.
	IOPAEN   APB2ENR_Bits = 0x01 << 2  //+ I/O port A clock enable.
	IOPBEN   APB2ENR_Bits = 0x01 << 3  //+ I/O port B clock enable.
	IOPCEN   APB2ENR_Bits = 0x01 << 4  //+ I/O port C clock enable.
	IOPDEN   APB2ENR_Bits = 0x01 << 5  //+ I/O port D clock enable.
	ADC1EN   APB2ENR_Bits = 0x01 << 9  //+ ADC 1 interface clock enable.
	ADC2EN   APB2ENR_Bits = 0x01 << 10 //+ ADC 2 interface clock enable.
	TIM1EN   APB2ENR_Bits = 0x01 << 11 //+ TIM1 Timer clock enable.
	SPI1EN   APB2ENR_Bits = 0x01 << 12 //+ SPI 1 clock enable.
	USART1EN APB2ENR_Bits = 0x01 << 14 //+ USART1 clock enable.
	IOPEEN   APB2ENR_Bits = 0x01 << 6  //+ I/O port E clock enable.
	IOPFEN   APB2ENR_Bits = 0x01 << 7  //+ I/O port F clock enable.
	IOPGEN   APB2ENR_Bits = 0x01 << 8  //+ I/O port G clock enable.
	TIM8EN   APB2ENR_Bits = 0x01 << 13 //+ TIM8 Timer clock enable.
	ADC3EN   APB2ENR_Bits = 0x01 << 15 //+ DMA1 clock enable.
)

const (
	TIM2EN   APB1ENR_Bits = 0x01 << 0  //+ Timer 2 clock enabled.
	TIM3EN   APB1ENR_Bits = 0x01 << 1  //+ Timer 3 clock enable.
	WWDGEN   APB1ENR_Bits = 0x01 << 11 //+ Window Watchdog clock enable.
	USART2EN APB1ENR_Bits = 0x01 << 17 //+ USART 2 clock enable.
	I2C1EN   APB1ENR_Bits = 0x01 << 21 //+ I2C 1 clock enable.
	CAN1EN   APB1ENR_Bits = 0x01 << 25 //+ CAN1 clock enable.
	BKPEN    APB1ENR_Bits = 0x01 << 27 //+ Backup interface clock enable.
	PWREN    APB1ENR_Bits = 0x01 << 28 //+ Power interface clock enable.
	TIM4EN   APB1ENR_Bits = 0x01 << 2  //+ Timer 4 clock enable.
	SPI2EN   APB1ENR_Bits = 0x01 << 14 //+ SPI 2 clock enable.
	USART3EN APB1ENR_Bits = 0x01 << 18 //+ USART 3 clock enable.
	I2C2EN   APB1ENR_Bits = 0x01 << 22 //+ I2C 2 clock enable.
	USBEN    APB1ENR_Bits = 0x01 << 23 //+ USB Device clock enable.
	TIM5EN   APB1ENR_Bits = 0x01 << 3  //+ Timer 5 clock enable.
	TIM6EN   APB1ENR_Bits = 0x01 << 4  //+ Timer 6 clock enable.
	TIM7EN   APB1ENR_Bits = 0x01 << 5  //+ Timer 7 clock enable.
	SPI3EN   APB1ENR_Bits = 0x01 << 15 //+ SPI 3 clock enable.
	UART4EN  APB1ENR_Bits = 0x01 << 19 //+ UART 4 clock enable.
	UART5EN  APB1ENR_Bits = 0x01 << 20 //+ UART 5 clock enable.
	DACEN    APB1ENR_Bits = 0x01 << 29 //+ DAC interface clock enable.
)

const (
	LSEON          BDCR_Bits = 0x01 << 0  //+ External Low Speed oscillator enable.
	LSERDY         BDCR_Bits = 0x01 << 1  //+ External Low Speed oscillator Ready.
	LSEBYP         BDCR_Bits = 0x01 << 2  //+ External Low Speed oscillator Bypass.
	RTCSEL         BDCR_Bits = 0x03 << 8  //+ RTCSEL[1:0] bits (RTC clock source selection).
	RTCSEL_0       BDCR_Bits = 0x01 << 8  //  Bit 0.
	RTCSEL_1       BDCR_Bits = 0x02 << 8  //  Bit 1.
	RTCSEL_NOCLOCK BDCR_Bits = 0x00 << 8  //  No clock.
	RTCSEL_LSE     BDCR_Bits = 0x01 << 8  //  LSE oscillator clock used as RTC clock.
	RTCSEL_LSI     BDCR_Bits = 0x02 << 8  //  LSI oscillator clock used as RTC clock.
	RTCSEL_HSE     BDCR_Bits = 0x03 << 8  //  HSE oscillator clock divided by 128 used as RTC clock.
	RTCEN          BDCR_Bits = 0x01 << 15 //+ RTC clock enable.
	BDRST          BDCR_Bits = 0x01 << 16 //+ Backup domain software reset.
)

const (
	LSION    CSR_Bits = 0x01 << 0  //+ Internal Low Speed oscillator enable.
	LSIRDY   CSR_Bits = 0x01 << 1  //+ Internal Low Speed oscillator Ready.
	RMVF     CSR_Bits = 0x01 << 24 //+ Remove reset flag.
	PINRSTF  CSR_Bits = 0x01 << 26 //+ PIN reset flag.
	PORRSTF  CSR_Bits = 0x01 << 27 //+ POR/PDR reset flag.
	SFTRSTF  CSR_Bits = 0x01 << 28 //+ Software Reset flag.
	IWDGRSTF CSR_Bits = 0x01 << 29 //+ Independent Watchdog reset flag.
	WWDGRSTF CSR_Bits = 0x01 << 30 //+ Window watchdog reset flag.
	LPWRRSTF CSR_Bits = 0x01 << 31 //+ Low-Power reset flag.
)
