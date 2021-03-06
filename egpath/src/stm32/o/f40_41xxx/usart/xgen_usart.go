package usart

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type USART_Periph struct {
	SR   RSR
	_    uint16
	DR   RDR
	_    uint16
	BRR  RBRR
	_    uint16
	CR1  RCR1
	_    uint16
	CR2  RCR2
	_    uint16
	CR3  RCR3
	_    uint16
	GTPR RGTPR
}

func (p *USART_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var USART2 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.USART2_BASE)))

//emgo:const
var USART3 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.USART3_BASE)))

//emgo:const
var UART4 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.UART4_BASE)))

//emgo:const
var UART5 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.UART5_BASE)))

//emgo:const
var UART7 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.UART7_BASE)))

//emgo:const
var UART8 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.UART8_BASE)))

//emgo:const
var USART1 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.USART1_BASE)))

//emgo:const
var USART6 = (*USART_Periph)(unsafe.Pointer(uintptr(mmap.USART6_BASE)))

type SR uint16

func (b SR) Field(mask SR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR) J(v int) SR {
	return SR(bits.Make32(v, uint32(mask)))
}

type RSR struct{ mmio.U16 }

func (r *RSR) Bits(mask SR) SR      { return SR(r.U16.Bits(uint16(mask))) }
func (r *RSR) StoreBits(mask, b SR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RSR) SetBits(mask SR)      { r.U16.SetBits(uint16(mask)) }
func (r *RSR) ClearBits(mask SR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RSR) Load() SR             { return SR(r.U16.Load()) }
func (r *RSR) Store(b SR)           { r.U16.Store(uint16(b)) }

type RMSR struct{ mmio.UM16 }

func (rm RMSR) Load() SR   { return SR(rm.UM16.Load()) }
func (rm RMSR) Store(b SR) { rm.UM16.Store(uint16(b)) }

func (p *USART_Periph) PE() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(PE)}}
}

func (p *USART_Periph) FE() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(FE)}}
}

func (p *USART_Periph) NE() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(NE)}}
}

func (p *USART_Periph) ORE() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(ORE)}}
}

func (p *USART_Periph) IDLE() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(IDLE)}}
}

func (p *USART_Periph) RXNE() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(RXNE)}}
}

func (p *USART_Periph) TC() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(TC)}}
}

func (p *USART_Periph) TXE() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(TXE)}}
}

func (p *USART_Periph) LBD() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(LBD)}}
}

func (p *USART_Periph) CTS() RMSR {
	return RMSR{mmio.UM16{&p.SR.U16, uint16(CTS)}}
}

type DR uint16

func (b DR) Field(mask DR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR) J(v int) DR {
	return DR(bits.Make32(v, uint32(mask)))
}

type RDR struct{ mmio.U16 }

func (r *RDR) Bits(mask DR) DR      { return DR(r.U16.Bits(uint16(mask))) }
func (r *RDR) StoreBits(mask, b DR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RDR) SetBits(mask DR)      { r.U16.SetBits(uint16(mask)) }
func (r *RDR) ClearBits(mask DR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RDR) Load() DR             { return DR(r.U16.Load()) }
func (r *RDR) Store(b DR)           { r.U16.Store(uint16(b)) }

type RMDR struct{ mmio.UM16 }

func (rm RMDR) Load() DR   { return DR(rm.UM16.Load()) }
func (rm RMDR) Store(b DR) { rm.UM16.Store(uint16(b)) }

type BRR uint16

func (b BRR) Field(mask BRR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BRR) J(v int) BRR {
	return BRR(bits.Make32(v, uint32(mask)))
}

type RBRR struct{ mmio.U16 }

func (r *RBRR) Bits(mask BRR) BRR     { return BRR(r.U16.Bits(uint16(mask))) }
func (r *RBRR) StoreBits(mask, b BRR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RBRR) SetBits(mask BRR)      { r.U16.SetBits(uint16(mask)) }
func (r *RBRR) ClearBits(mask BRR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RBRR) Load() BRR             { return BRR(r.U16.Load()) }
func (r *RBRR) Store(b BRR)           { r.U16.Store(uint16(b)) }

type RMBRR struct{ mmio.UM16 }

func (rm RMBRR) Load() BRR   { return BRR(rm.UM16.Load()) }
func (rm RMBRR) Store(b BRR) { rm.UM16.Store(uint16(b)) }

func (p *USART_Periph) DIV_Fraction() RMBRR {
	return RMBRR{mmio.UM16{&p.BRR.U16, uint16(DIV_Fraction)}}
}

func (p *USART_Periph) DIV_Mantissa() RMBRR {
	return RMBRR{mmio.UM16{&p.BRR.U16, uint16(DIV_Mantissa)}}
}

type CR1 uint16

func (b CR1) Field(mask CR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1) J(v int) CR1 {
	return CR1(bits.Make32(v, uint32(mask)))
}

type RCR1 struct{ mmio.U16 }

func (r *RCR1) Bits(mask CR1) CR1     { return CR1(r.U16.Bits(uint16(mask))) }
func (r *RCR1) StoreBits(mask, b CR1) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCR1) SetBits(mask CR1)      { r.U16.SetBits(uint16(mask)) }
func (r *RCR1) ClearBits(mask CR1)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCR1) Load() CR1             { return CR1(r.U16.Load()) }
func (r *RCR1) Store(b CR1)           { r.U16.Store(uint16(b)) }

type RMCR1 struct{ mmio.UM16 }

func (rm RMCR1) Load() CR1   { return CR1(rm.UM16.Load()) }
func (rm RMCR1) Store(b CR1) { rm.UM16.Store(uint16(b)) }

func (p *USART_Periph) SBK() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(SBK)}}
}

func (p *USART_Periph) RWU() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(RWU)}}
}

func (p *USART_Periph) RE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(RE)}}
}

func (p *USART_Periph) TE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(TE)}}
}

func (p *USART_Periph) IDLEIE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(IDLEIE)}}
}

func (p *USART_Periph) RXNEIE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(RXNEIE)}}
}

func (p *USART_Periph) TCIE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(TCIE)}}
}

func (p *USART_Periph) TXEIE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(TXEIE)}}
}

func (p *USART_Periph) PEIE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(PEIE)}}
}

func (p *USART_Periph) PS() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(PS)}}
}

func (p *USART_Periph) PCE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(PCE)}}
}

func (p *USART_Periph) WAKE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(WAKE)}}
}

func (p *USART_Periph) M() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(M)}}
}

func (p *USART_Periph) UE() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(UE)}}
}

func (p *USART_Periph) OVER8() RMCR1 {
	return RMCR1{mmio.UM16{&p.CR1.U16, uint16(OVER8)}}
}

type CR2 uint16

func (b CR2) Field(mask CR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2) J(v int) CR2 {
	return CR2(bits.Make32(v, uint32(mask)))
}

type RCR2 struct{ mmio.U16 }

func (r *RCR2) Bits(mask CR2) CR2     { return CR2(r.U16.Bits(uint16(mask))) }
func (r *RCR2) StoreBits(mask, b CR2) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCR2) SetBits(mask CR2)      { r.U16.SetBits(uint16(mask)) }
func (r *RCR2) ClearBits(mask CR2)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCR2) Load() CR2             { return CR2(r.U16.Load()) }
func (r *RCR2) Store(b CR2)           { r.U16.Store(uint16(b)) }

type RMCR2 struct{ mmio.UM16 }

func (rm RMCR2) Load() CR2   { return CR2(rm.UM16.Load()) }
func (rm RMCR2) Store(b CR2) { rm.UM16.Store(uint16(b)) }

func (p *USART_Periph) ADD() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(ADD)}}
}

func (p *USART_Periph) LBDL() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(LBDL)}}
}

func (p *USART_Periph) LBDIE() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(LBDIE)}}
}

func (p *USART_Periph) LBCL() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(LBCL)}}
}

func (p *USART_Periph) CPHA() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(CPHA)}}
}

func (p *USART_Periph) CPOL() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(CPOL)}}
}

func (p *USART_Periph) CLKEN() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(CLKEN)}}
}

func (p *USART_Periph) STOP() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(STOP)}}
}

func (p *USART_Periph) LINEN() RMCR2 {
	return RMCR2{mmio.UM16{&p.CR2.U16, uint16(LINEN)}}
}

type CR3 uint16

func (b CR3) Field(mask CR3) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR3) J(v int) CR3 {
	return CR3(bits.Make32(v, uint32(mask)))
}

type RCR3 struct{ mmio.U16 }

func (r *RCR3) Bits(mask CR3) CR3     { return CR3(r.U16.Bits(uint16(mask))) }
func (r *RCR3) StoreBits(mask, b CR3) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCR3) SetBits(mask CR3)      { r.U16.SetBits(uint16(mask)) }
func (r *RCR3) ClearBits(mask CR3)    { r.U16.ClearBits(uint16(mask)) }
func (r *RCR3) Load() CR3             { return CR3(r.U16.Load()) }
func (r *RCR3) Store(b CR3)           { r.U16.Store(uint16(b)) }

type RMCR3 struct{ mmio.UM16 }

func (rm RMCR3) Load() CR3   { return CR3(rm.UM16.Load()) }
func (rm RMCR3) Store(b CR3) { rm.UM16.Store(uint16(b)) }

func (p *USART_Periph) EIE() RMCR3 {
	return RMCR3{mmio.UM16{&p.CR3.U16, uint16(EIE)}}
}

func (p *USART_Periph) IREN() RMCR3 {
	return RMCR3{mmio.UM16{&p.CR3.U16, uint16(IREN)}}
}

func (p *USART_Periph) IRLP() RMCR3 {
	return RMCR3{mmio.UM16{&p.CR3.U16, uint16(IRLP)}}
}

func (p *USART_Periph) HDSEL() RMCR3 {
	return RMCR3{mmio.UM16{&p.CR3.U16, uint16(HDSEL)}}
}

func (p *USART_Periph) NACK() RMCR3 {
	return RMCR3{mmio.UM16{&p.CR3.U16, uint16(NACK)}}
}

func (p *USART_Periph) SCEN() RMCR3 {
	return RMCR3{mmio.UM16{&p.CR3.U16, uint16(SCEN)}}
}

func (p *USART_Periph) DMAR() RMCR3 {
	return RMCR3{mmio.UM16{&p.CR3.U16, uint16(DMAR)}}
}

func (p *USART_Periph) DMAT() RMCR3 {
	return RMCR3{mmio.UM16{&p.CR3.U16, uint16(DMAT)}}
}

func (p *USART_Periph) RTSE() RMCR3 {
	return RMCR3{mmio.UM16{&p.CR3.U16, uint16(RTSE)}}
}

func (p *USART_Periph) CTSE() RMCR3 {
	return RMCR3{mmio.UM16{&p.CR3.U16, uint16(CTSE)}}
}

func (p *USART_Periph) CTSIE() RMCR3 {
	return RMCR3{mmio.UM16{&p.CR3.U16, uint16(CTSIE)}}
}

func (p *USART_Periph) ONEBIT() RMCR3 {
	return RMCR3{mmio.UM16{&p.CR3.U16, uint16(ONEBIT)}}
}

type GTPR uint16

func (b GTPR) Field(mask GTPR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask GTPR) J(v int) GTPR {
	return GTPR(bits.Make32(v, uint32(mask)))
}

type RGTPR struct{ mmio.U16 }

func (r *RGTPR) Bits(mask GTPR) GTPR    { return GTPR(r.U16.Bits(uint16(mask))) }
func (r *RGTPR) StoreBits(mask, b GTPR) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RGTPR) SetBits(mask GTPR)      { r.U16.SetBits(uint16(mask)) }
func (r *RGTPR) ClearBits(mask GTPR)    { r.U16.ClearBits(uint16(mask)) }
func (r *RGTPR) Load() GTPR             { return GTPR(r.U16.Load()) }
func (r *RGTPR) Store(b GTPR)           { r.U16.Store(uint16(b)) }

type RMGTPR struct{ mmio.UM16 }

func (rm RMGTPR) Load() GTPR   { return GTPR(rm.UM16.Load()) }
func (rm RMGTPR) Store(b GTPR) { rm.UM16.Store(uint16(b)) }

func (p *USART_Periph) PSC() RMGTPR {
	return RMGTPR{mmio.UM16{&p.GTPR.U16, uint16(PSC)}}
}

func (p *USART_Periph) GT() RMGTPR {
	return RMGTPR{mmio.UM16{&p.GTPR.U16, uint16(GT)}}
}
