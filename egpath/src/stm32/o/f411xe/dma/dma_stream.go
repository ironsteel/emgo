// +build f411xe

// Peripheral: DMA_Stream_Periph  DMA Controller.
// Instances:
//  DMA1_Stream0  mmap.DMA1_Stream0_BASE
//  DMA1_Stream1  mmap.DMA1_Stream1_BASE
//  DMA1_Stream2  mmap.DMA1_Stream2_BASE
//  DMA1_Stream3  mmap.DMA1_Stream3_BASE
//  DMA1_Stream4  mmap.DMA1_Stream4_BASE
//  DMA1_Stream5  mmap.DMA1_Stream5_BASE
//  DMA1_Stream6  mmap.DMA1_Stream6_BASE
//  DMA1_Stream7  mmap.DMA1_Stream7_BASE
//  DMA2_Stream0  mmap.DMA2_Stream0_BASE
//  DMA2_Stream1  mmap.DMA2_Stream1_BASE
//  DMA2_Stream2  mmap.DMA2_Stream2_BASE
//  DMA2_Stream3  mmap.DMA2_Stream3_BASE
//  DMA2_Stream4  mmap.DMA2_Stream4_BASE
//  DMA2_Stream5  mmap.DMA2_Stream5_BASE
//  DMA2_Stream6  mmap.DMA2_Stream6_BASE
//  DMA2_Stream7  mmap.DMA2_Stream7_BASE
// Registers:
//  0x00 32  CR   DMA stream x configuration register.
//  0x04 32  NDTR DMA stream x number of data register.
//  0x08 32  PAR  DMA stream x peripheral address register.
//  0x0C 32  M0AR DMA stream x memory 0 address register.
//  0x10 32  M1AR DMA stream x memory 1 address register.
//  0x14 32  FCR  DMA stream x FIFO control register.
// Import:
//  stm32/o/f411xe/mmap
package dma
