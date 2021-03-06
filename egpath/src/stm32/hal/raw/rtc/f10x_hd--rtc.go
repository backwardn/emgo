// +build f10x_hd

// Peripheral: RTC_Periph  Real-Time Clock.
// Instances:
//  RTC  mmap.RTC_BASE
// Registers:
//  0x00 16  CRH
//  0x04 16  CRL
//  0x08 16  PRLH
//  0x0C 16  PRLL
//  0x10 16  DIVH
//  0x14 16  DIVL
//  0x18 16  CNTH
//  0x1C 16  CNTL
//  0x20 16  ALRH
//  0x24 16  ALRL
// Import:
//  stm32/o/f10x_hd/mmap
package rtc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	SECIE CRH = 0x01 << 0 //+ Second Interrupt Enable.
	ALRIE CRH = 0x01 << 1 //+ Alarm Interrupt Enable.
	OWIE  CRH = 0x01 << 2 //+ OverfloW Interrupt Enable.
)

const (
	SECIEn = 0
	ALRIEn = 1
	OWIEn  = 2
)

const (
	SECF  CRL = 0x01 << 0 //+ Second Flag.
	ALRF  CRL = 0x01 << 1 //+ Alarm Flag.
	OWF   CRL = 0x01 << 2 //+ OverfloW Flag.
	RSF   CRL = 0x01 << 3 //+ Registers Synchronized Flag.
	CNF   CRL = 0x01 << 4 //+ Configuration Flag.
	RTOFF CRL = 0x01 << 5 //+ RTC operation OFF.
)

const (
	SECFn  = 0
	ALRFn  = 1
	OWFn   = 2
	RSFn   = 3
	CNFn   = 4
	RTOFFn = 5
)
