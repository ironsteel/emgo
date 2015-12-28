// +build f40_41xxx

// Package irq provides list of external interrupts.
package irq

import "arch/cortexm/nvic"

const (
	WWDG               nvic.IRQ = 0  // Window WatchDog Interrupt.
	PVD                nvic.IRQ = 1  // PVD through EXTI Line detection Interrupt.
	TAMP_STAMP         nvic.IRQ = 2  // Tamper and TimeStamp interrupts through the EXTI line.
	RTC_WKUP           nvic.IRQ = 3  // RTC Wakeup interrupt through the EXTI line.
	FLASH              nvic.IRQ = 4  // FLASH global Interrupt.
	RCC                nvic.IRQ = 5  // RCC global Interrupt.
	EXTI0              nvic.IRQ = 6  // EXTI Line0 Interrupt.
	EXTI1              nvic.IRQ = 7  // EXTI Line1 Interrupt.
	EXTI2              nvic.IRQ = 8  // EXTI Line2 Interrupt.
	EXTI3              nvic.IRQ = 9  // EXTI Line3 Interrupt.
	EXTI4              nvic.IRQ = 10 // EXTI Line4 Interrupt.
	DMA1_Stream0       nvic.IRQ = 11 // DMA1 Stream 0 global Interrupt.
	DMA1_Stream1       nvic.IRQ = 12 // DMA1 Stream 1 global Interrupt.
	DMA1_Stream2       nvic.IRQ = 13 // DMA1 Stream 2 global Interrupt.
	DMA1_Stream3       nvic.IRQ = 14 // DMA1 Stream 3 global Interrupt.
	DMA1_Stream4       nvic.IRQ = 15 // DMA1 Stream 4 global Interrupt.
	DMA1_Stream5       nvic.IRQ = 16 // DMA1 Stream 5 global Interrupt.
	DMA1_Stream6       nvic.IRQ = 17 // DMA1 Stream 6 global Interrupt.
	ADC                nvic.IRQ = 18 // ADC1, ADC2 and ADC3 global Interrupts.
	CAN1_TX            nvic.IRQ = 19 // CAN1 TX Interrupt.
	CAN1_RX0           nvic.IRQ = 20 // CAN1 RX0 Interrupt.
	CAN1_RX1           nvic.IRQ = 21 // CAN1 RX1 Interrupt.
	CAN1_SCE           nvic.IRQ = 22 // CAN1 SCE Interrupt.
	EXTI9_5            nvic.IRQ = 23 // External Line[9:5] Interrupts.
	TIM1_BRK_TIM9      nvic.IRQ = 24 // TIM1 Break interrupt and TIM9 global interrupt.
	TIM1_UP_TIM10      nvic.IRQ = 25 // TIM1 Update Interrupt and TIM10 global interrupt.
	TIM1_TRG_COM_TIM11 nvic.IRQ = 26 // TIM1 Trigger and Commutation Interrupt and TIM11 global interrupt.
	TIM1_CC            nvic.IRQ = 27 // TIM1 Capture Compare Interrupt.
	TIM2               nvic.IRQ = 28 // TIM2 global Interrupt.
	TIM3               nvic.IRQ = 29 // TIM3 global Interrupt.
	TIM4               nvic.IRQ = 30 // TIM4 global Interrupt.
	I2C1_EV            nvic.IRQ = 31 // I2C1 Event Interrupt.
	I2C1_ER            nvic.IRQ = 32 // I2C1 Error Interrupt.
	I2C2_EV            nvic.IRQ = 33 // I2C2 Event Interrupt.
	I2C2_ER            nvic.IRQ = 34 // I2C2 Error Interrupt.
	SPI1               nvic.IRQ = 35 // SPI1 global Interrupt.
	SPI2               nvic.IRQ = 36 // SPI2 global Interrupt.
	USART1             nvic.IRQ = 37 // USART1 global Interrupt.
	USART2             nvic.IRQ = 38 // USART2 global Interrupt.
	USART3             nvic.IRQ = 39 // USART3 global Interrupt.
	EXTI15_10          nvic.IRQ = 40 // External Line[15:10] Interrupts.
	RTC_Alarm          nvic.IRQ = 41 // RTC Alarm (A and B) through EXTI Line Interrupt.
	OTG_FS_WKUP        nvic.IRQ = 42 // USB OTG FS Wakeup through EXTI line interrupt.
	TIM8_BRK_TIM12     nvic.IRQ = 43 // TIM8 Break Interrupt and TIM12 global interrupt.
	TIM8_UP_TIM13      nvic.IRQ = 44 // TIM8 Update Interrupt and TIM13 global interrupt.
	TIM8_TRG_COM_TIM14 nvic.IRQ = 45 // TIM8 Trigger and Commutation Interrupt and TIM14 global interrupt.
	TIM8_CC            nvic.IRQ = 46 // TIM8 Capture Compare Interrupt.
	DMA1_Stream7       nvic.IRQ = 47 // DMA1 Stream7 Interrupt.
	FSMC               nvic.IRQ = 48 // FSMC global Interrupt.
	SDIO               nvic.IRQ = 49 // SDIO global Interrupt.
	TIM5               nvic.IRQ = 50 // TIM5 global Interrupt.
	SPI3               nvic.IRQ = 51 // SPI3 global Interrupt.
	UART4              nvic.IRQ = 52 // UART4 global Interrupt.
	UART5              nvic.IRQ = 53 // UART5 global Interrupt.
	TIM6_DAC           nvic.IRQ = 54 // TIM6 global and DAC1&2 underrun error  interrupts.
	TIM7               nvic.IRQ = 55 // TIM7 global interrupt.
	DMA2_Stream0       nvic.IRQ = 56 // DMA2 Stream 0 global Interrupt.
	DMA2_Stream1       nvic.IRQ = 57 // DMA2 Stream 1 global Interrupt.
	DMA2_Stream2       nvic.IRQ = 58 // DMA2 Stream 2 global Interrupt.
	DMA2_Stream3       nvic.IRQ = 59 // DMA2 Stream 3 global Interrupt.
	DMA2_Stream4       nvic.IRQ = 60 // DMA2 Stream 4 global Interrupt.
	ETH                nvic.IRQ = 61 // Ethernet global Interrupt.
	ETH_WKUP           nvic.IRQ = 62 // Ethernet Wakeup through EXTI line Interrupt.
	CAN2_TX            nvic.IRQ = 63 // CAN2 TX Interrupt.
	CAN2_RX0           nvic.IRQ = 64 // CAN2 RX0 Interrupt.
	CAN2_RX1           nvic.IRQ = 65 // CAN2 RX1 Interrupt.
	CAN2_SCE           nvic.IRQ = 66 // CAN2 SCE Interrupt.
	OTG_FS             nvic.IRQ = 67 // USB OTG FS global Interrupt.
	DMA2_Stream5       nvic.IRQ = 68 // DMA2 Stream 5 global interrupt.
	DMA2_Stream6       nvic.IRQ = 69 // DMA2 Stream 6 global interrupt.
	DMA2_Stream7       nvic.IRQ = 70 // DMA2 Stream 7 global interrupt.
	USART6             nvic.IRQ = 71 // USART6 global interrupt.
	I2C3_EV            nvic.IRQ = 72 // I2C3 event interrupt.
	I2C3_ER            nvic.IRQ = 73 // I2C3 error interrupt.
	OTG_HS_EP1_OUT     nvic.IRQ = 74 // USB OTG HS End Point 1 Out global interrupt.
	OTG_HS_EP1_IN      nvic.IRQ = 75 // USB OTG HS End Point 1 In global interrupt.
	OTG_HS_WKUP        nvic.IRQ = 76 // USB OTG HS Wakeup through EXTI interrupt.
	OTG_HS             nvic.IRQ = 77 // USB OTG HS global interrupt.
	DCMI               nvic.IRQ = 78 // DCMI global interrupt.
	CRYP               nvic.IRQ = 79 // CRYP crypto global interrupt.
	HASH_RNG           nvic.IRQ = 80 // Hash and Rng global interrupt.
	FPU                nvic.IRQ = 81 // FPU global interrupt.
)
