// +build f40_41xxx

package ltdc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type LTDC_Periph struct {
	_     [2]uint32
	SSCR  SSCR
	BPCR  BPCR
	AWCR  AWCR
	TWCR  TWCR
	GCR   GCR
	_     [2]uint32
	SRCR  SRCR
	_     uint32
	BCCR  BCCR
	_     uint32
	IER   IER
	ISR   ISR
	ICR   ICR
	LIPCR LIPCR
	CPSR  CPSR
	CDSR  CDSR
}

func (p *LTDC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var LTDC = (*LTDC_Periph)(unsafe.Pointer(uintptr(mmap.LTDC_BASE)))

type SSCR_Bits uint32

func (b SSCR_Bits) Field(mask SSCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SSCR_Bits) J(v int) SSCR_Bits {
	return SSCR_Bits(bits.Make32(v, uint32(mask)))
}

type SSCR struct{ mmio.U32 }

func (r *SSCR) Bits(mask SSCR_Bits) SSCR_Bits { return SSCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SSCR) StoreBits(mask, b SSCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SSCR) SetBits(mask SSCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *SSCR) ClearBits(mask SSCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *SSCR) Load() SSCR_Bits               { return SSCR_Bits(r.U32.Load()) }
func (r *SSCR) Store(b SSCR_Bits)             { r.U32.Store(uint32(b)) }

type SSCR_Mask struct{ mmio.UM32 }

func (rm SSCR_Mask) Load() SSCR_Bits   { return SSCR_Bits(rm.UM32.Load()) }
func (rm SSCR_Mask) Store(b SSCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) VSH() SSCR_Mask {
	return SSCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(VSH)}}
}

func (p *LTDC_Periph) HSW() SSCR_Mask {
	return SSCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(HSW)}}
}

type BPCR_Bits uint32

func (b BPCR_Bits) Field(mask BPCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BPCR_Bits) J(v int) BPCR_Bits {
	return BPCR_Bits(bits.Make32(v, uint32(mask)))
}

type BPCR struct{ mmio.U32 }

func (r *BPCR) Bits(mask BPCR_Bits) BPCR_Bits { return BPCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BPCR) StoreBits(mask, b BPCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BPCR) SetBits(mask BPCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *BPCR) ClearBits(mask BPCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *BPCR) Load() BPCR_Bits               { return BPCR_Bits(r.U32.Load()) }
func (r *BPCR) Store(b BPCR_Bits)             { r.U32.Store(uint32(b)) }

type BPCR_Mask struct{ mmio.UM32 }

func (rm BPCR_Mask) Load() BPCR_Bits   { return BPCR_Bits(rm.UM32.Load()) }
func (rm BPCR_Mask) Store(b BPCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) AVBP() BPCR_Mask {
	return BPCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(AVBP)}}
}

func (p *LTDC_Periph) AHBP() BPCR_Mask {
	return BPCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(AHBP)}}
}

type AWCR_Bits uint32

func (b AWCR_Bits) Field(mask AWCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AWCR_Bits) J(v int) AWCR_Bits {
	return AWCR_Bits(bits.Make32(v, uint32(mask)))
}

type AWCR struct{ mmio.U32 }

func (r *AWCR) Bits(mask AWCR_Bits) AWCR_Bits { return AWCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *AWCR) StoreBits(mask, b AWCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *AWCR) SetBits(mask AWCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *AWCR) ClearBits(mask AWCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *AWCR) Load() AWCR_Bits               { return AWCR_Bits(r.U32.Load()) }
func (r *AWCR) Store(b AWCR_Bits)             { r.U32.Store(uint32(b)) }

type AWCR_Mask struct{ mmio.UM32 }

func (rm AWCR_Mask) Load() AWCR_Bits   { return AWCR_Bits(rm.UM32.Load()) }
func (rm AWCR_Mask) Store(b AWCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) AAH() AWCR_Mask {
	return AWCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(AAH)}}
}

func (p *LTDC_Periph) AAW() AWCR_Mask {
	return AWCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(AAW)}}
}

type TWCR_Bits uint32

func (b TWCR_Bits) Field(mask TWCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TWCR_Bits) J(v int) TWCR_Bits {
	return TWCR_Bits(bits.Make32(v, uint32(mask)))
}

type TWCR struct{ mmio.U32 }

func (r *TWCR) Bits(mask TWCR_Bits) TWCR_Bits { return TWCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TWCR) StoreBits(mask, b TWCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TWCR) SetBits(mask TWCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *TWCR) ClearBits(mask TWCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *TWCR) Load() TWCR_Bits               { return TWCR_Bits(r.U32.Load()) }
func (r *TWCR) Store(b TWCR_Bits)             { r.U32.Store(uint32(b)) }

type TWCR_Mask struct{ mmio.UM32 }

func (rm TWCR_Mask) Load() TWCR_Bits   { return TWCR_Bits(rm.UM32.Load()) }
func (rm TWCR_Mask) Store(b TWCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) TOTALH() TWCR_Mask {
	return TWCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(TOTALH)}}
}

func (p *LTDC_Periph) TOTALW() TWCR_Mask {
	return TWCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(TOTALW)}}
}

type GCR_Bits uint32

func (b GCR_Bits) Field(mask GCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask GCR_Bits) J(v int) GCR_Bits {
	return GCR_Bits(bits.Make32(v, uint32(mask)))
}

type GCR struct{ mmio.U32 }

func (r *GCR) Bits(mask GCR_Bits) GCR_Bits { return GCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *GCR) StoreBits(mask, b GCR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GCR) SetBits(mask GCR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *GCR) ClearBits(mask GCR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *GCR) Load() GCR_Bits              { return GCR_Bits(r.U32.Load()) }
func (r *GCR) Store(b GCR_Bits)            { r.U32.Store(uint32(b)) }

type GCR_Mask struct{ mmio.UM32 }

func (rm GCR_Mask) Load() GCR_Bits   { return GCR_Bits(rm.UM32.Load()) }
func (rm GCR_Mask) Store(b GCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) LTDCEN() GCR_Mask {
	return GCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(LTDCEN)}}
}

func (p *LTDC_Periph) DBW() GCR_Mask {
	return GCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(DBW)}}
}

func (p *LTDC_Periph) DGW() GCR_Mask {
	return GCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(DGW)}}
}

func (p *LTDC_Periph) DRW() GCR_Mask {
	return GCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(DRW)}}
}

func (p *LTDC_Periph) DTEN() GCR_Mask {
	return GCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(DTEN)}}
}

func (p *LTDC_Periph) PCPOL() GCR_Mask {
	return GCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(PCPOL)}}
}

func (p *LTDC_Periph) DEPOL() GCR_Mask {
	return GCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(DEPOL)}}
}

func (p *LTDC_Periph) VSPOL() GCR_Mask {
	return GCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(VSPOL)}}
}

func (p *LTDC_Periph) HSPOL() GCR_Mask {
	return GCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(HSPOL)}}
}

type SRCR_Bits uint32

func (b SRCR_Bits) Field(mask SRCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SRCR_Bits) J(v int) SRCR_Bits {
	return SRCR_Bits(bits.Make32(v, uint32(mask)))
}

type SRCR struct{ mmio.U32 }

func (r *SRCR) Bits(mask SRCR_Bits) SRCR_Bits { return SRCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SRCR) StoreBits(mask, b SRCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SRCR) SetBits(mask SRCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *SRCR) ClearBits(mask SRCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *SRCR) Load() SRCR_Bits               { return SRCR_Bits(r.U32.Load()) }
func (r *SRCR) Store(b SRCR_Bits)             { r.U32.Store(uint32(b)) }

type SRCR_Mask struct{ mmio.UM32 }

func (rm SRCR_Mask) Load() SRCR_Bits   { return SRCR_Bits(rm.UM32.Load()) }
func (rm SRCR_Mask) Store(b SRCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) IMR() SRCR_Mask {
	return SRCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 36)), uint32(IMR)}}
}

func (p *LTDC_Periph) VBR() SRCR_Mask {
	return SRCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 36)), uint32(VBR)}}
}

type BCCR_Bits uint32

func (b BCCR_Bits) Field(mask BCCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BCCR_Bits) J(v int) BCCR_Bits {
	return BCCR_Bits(bits.Make32(v, uint32(mask)))
}

type BCCR struct{ mmio.U32 }

func (r *BCCR) Bits(mask BCCR_Bits) BCCR_Bits { return BCCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BCCR) StoreBits(mask, b BCCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BCCR) SetBits(mask BCCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *BCCR) ClearBits(mask BCCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *BCCR) Load() BCCR_Bits               { return BCCR_Bits(r.U32.Load()) }
func (r *BCCR) Store(b BCCR_Bits)             { r.U32.Store(uint32(b)) }

type BCCR_Mask struct{ mmio.UM32 }

func (rm BCCR_Mask) Load() BCCR_Bits   { return BCCR_Bits(rm.UM32.Load()) }
func (rm BCCR_Mask) Store(b BCCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) BCBLUE() BCCR_Mask {
	return BCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 44)), uint32(BCBLUE)}}
}

func (p *LTDC_Periph) BCGREEN() BCCR_Mask {
	return BCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 44)), uint32(BCGREEN)}}
}

func (p *LTDC_Periph) BCRED() BCCR_Mask {
	return BCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 44)), uint32(BCRED)}}
}

type IER_Bits uint32

func (b IER_Bits) Field(mask IER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IER_Bits) J(v int) IER_Bits {
	return IER_Bits(bits.Make32(v, uint32(mask)))
}

type IER struct{ mmio.U32 }

func (r *IER) Bits(mask IER_Bits) IER_Bits { return IER_Bits(r.U32.Bits(uint32(mask))) }
func (r *IER) StoreBits(mask, b IER_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IER) SetBits(mask IER_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *IER) ClearBits(mask IER_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *IER) Load() IER_Bits              { return IER_Bits(r.U32.Load()) }
func (r *IER) Store(b IER_Bits)            { r.U32.Store(uint32(b)) }

type IER_Mask struct{ mmio.UM32 }

func (rm IER_Mask) Load() IER_Bits   { return IER_Bits(rm.UM32.Load()) }
func (rm IER_Mask) Store(b IER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) LIE() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(LIE)}}
}

func (p *LTDC_Periph) FUIE() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(FUIE)}}
}

func (p *LTDC_Periph) TERRIE() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(TERRIE)}}
}

func (p *LTDC_Periph) RRIE() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(RRIE)}}
}

type ISR_Bits uint32

func (b ISR_Bits) Field(mask ISR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ISR_Bits) J(v int) ISR_Bits {
	return ISR_Bits(bits.Make32(v, uint32(mask)))
}

type ISR struct{ mmio.U32 }

func (r *ISR) Bits(mask ISR_Bits) ISR_Bits { return ISR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ISR) StoreBits(mask, b ISR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ISR) SetBits(mask ISR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ISR) ClearBits(mask ISR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ISR) Load() ISR_Bits              { return ISR_Bits(r.U32.Load()) }
func (r *ISR) Store(b ISR_Bits)            { r.U32.Store(uint32(b)) }

type ISR_Mask struct{ mmio.UM32 }

func (rm ISR_Mask) Load() ISR_Bits   { return ISR_Bits(rm.UM32.Load()) }
func (rm ISR_Mask) Store(b ISR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) LIF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(LIF)}}
}

func (p *LTDC_Periph) FUIF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(FUIF)}}
}

func (p *LTDC_Periph) TERRIF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(TERRIF)}}
}

func (p *LTDC_Periph) RRIF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(RRIF)}}
}

type ICR_Bits uint32

func (b ICR_Bits) Field(mask ICR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ICR_Bits) J(v int) ICR_Bits {
	return ICR_Bits(bits.Make32(v, uint32(mask)))
}

type ICR struct{ mmio.U32 }

func (r *ICR) Bits(mask ICR_Bits) ICR_Bits { return ICR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ICR) StoreBits(mask, b ICR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ICR) SetBits(mask ICR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ICR) ClearBits(mask ICR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ICR) Load() ICR_Bits              { return ICR_Bits(r.U32.Load()) }
func (r *ICR) Store(b ICR_Bits)            { r.U32.Store(uint32(b)) }

type ICR_Mask struct{ mmio.UM32 }

func (rm ICR_Mask) Load() ICR_Bits   { return ICR_Bits(rm.UM32.Load()) }
func (rm ICR_Mask) Store(b ICR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) CLIF() ICR_Mask {
	return ICR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(CLIF)}}
}

func (p *LTDC_Periph) CFUIF() ICR_Mask {
	return ICR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(CFUIF)}}
}

func (p *LTDC_Periph) CTERRIF() ICR_Mask {
	return ICR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(CTERRIF)}}
}

func (p *LTDC_Periph) CRRIF() ICR_Mask {
	return ICR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(CRRIF)}}
}

type LIPCR_Bits uint32

func (b LIPCR_Bits) Field(mask LIPCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LIPCR_Bits) J(v int) LIPCR_Bits {
	return LIPCR_Bits(bits.Make32(v, uint32(mask)))
}

type LIPCR struct{ mmio.U32 }

func (r *LIPCR) Bits(mask LIPCR_Bits) LIPCR_Bits { return LIPCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *LIPCR) StoreBits(mask, b LIPCR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *LIPCR) SetBits(mask LIPCR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *LIPCR) ClearBits(mask LIPCR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *LIPCR) Load() LIPCR_Bits                { return LIPCR_Bits(r.U32.Load()) }
func (r *LIPCR) Store(b LIPCR_Bits)              { r.U32.Store(uint32(b)) }

type LIPCR_Mask struct{ mmio.UM32 }

func (rm LIPCR_Mask) Load() LIPCR_Bits   { return LIPCR_Bits(rm.UM32.Load()) }
func (rm LIPCR_Mask) Store(b LIPCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) LIPOS() LIPCR_Mask {
	return LIPCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(LIPOS)}}
}

type CPSR_Bits uint32

func (b CPSR_Bits) Field(mask CPSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CPSR_Bits) J(v int) CPSR_Bits {
	return CPSR_Bits(bits.Make32(v, uint32(mask)))
}

type CPSR struct{ mmio.U32 }

func (r *CPSR) Bits(mask CPSR_Bits) CPSR_Bits { return CPSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CPSR) StoreBits(mask, b CPSR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CPSR) SetBits(mask CPSR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CPSR) ClearBits(mask CPSR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CPSR) Load() CPSR_Bits               { return CPSR_Bits(r.U32.Load()) }
func (r *CPSR) Store(b CPSR_Bits)             { r.U32.Store(uint32(b)) }

type CPSR_Mask struct{ mmio.UM32 }

func (rm CPSR_Mask) Load() CPSR_Bits   { return CPSR_Bits(rm.UM32.Load()) }
func (rm CPSR_Mask) Store(b CPSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) CYPOS() CPSR_Mask {
	return CPSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 68)), uint32(CYPOS)}}
}

func (p *LTDC_Periph) CXPOS() CPSR_Mask {
	return CPSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 68)), uint32(CXPOS)}}
}

type CDSR_Bits uint32

func (b CDSR_Bits) Field(mask CDSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CDSR_Bits) J(v int) CDSR_Bits {
	return CDSR_Bits(bits.Make32(v, uint32(mask)))
}

type CDSR struct{ mmio.U32 }

func (r *CDSR) Bits(mask CDSR_Bits) CDSR_Bits { return CDSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CDSR) StoreBits(mask, b CDSR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CDSR) SetBits(mask CDSR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CDSR) ClearBits(mask CDSR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CDSR) Load() CDSR_Bits               { return CDSR_Bits(r.U32.Load()) }
func (r *CDSR) Store(b CDSR_Bits)             { r.U32.Store(uint32(b)) }

type CDSR_Mask struct{ mmio.UM32 }

func (rm CDSR_Mask) Load() CDSR_Bits   { return CDSR_Bits(rm.UM32.Load()) }
func (rm CDSR_Mask) Store(b CDSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LTDC_Periph) VDES() CDSR_Mask {
	return CDSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 72)), uint32(VDES)}}
}

func (p *LTDC_Periph) HDES() CDSR_Mask {
	return CDSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 72)), uint32(HDES)}}
}

func (p *LTDC_Periph) VSYNCS() CDSR_Mask {
	return CDSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 72)), uint32(VSYNCS)}}
}

func (p *LTDC_Periph) HSYNCS() CDSR_Mask {
	return CDSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 72)), uint32(HSYNCS)}}
}