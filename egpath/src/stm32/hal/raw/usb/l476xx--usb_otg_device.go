// +build l476xx

// Peripheral: USB_OTG_Device_Periph  USB_OTG_device_Registers.
// Instances:
// Registers:
//  0x00 32  DCFG       dev Configuration Register   800h.
//  0x04 32  DCTL       dev Control Register         804h.
//  0x08 32  DSTS       dev Status Register (RO)     808h.
//  0x10 32  DIEPMSK    dev IN Endpoint Mask         810h.
//  0x14 32  DOEPMSK    dev OUT Endpoint Mask        814h.
//  0x18 32  DAINT      dev All Endpoints Itr Reg    818h.
//  0x1C 32  DAINTMSK   dev All Endpoints Itr Mask   81Ch.
//  0x28 32  DVBUSDIS   dev VBUS discharge Register  828h.
//  0x2C 32  DVBUSPULSE dev VBUS Pulse Register      82Ch.
//  0x30 32  DTHRCTL    dev thr                      830h.
//  0x34 32  DIEPEMPMSK dev empty msk             834h.
//  0x38 32  DEACHINT   dedicated EP interrupt       838h.
//  0x3C 32  DEACHMSK   dedicated EP msk             83Ch.
//  0x44 32  DINEP1MSK  dedicated EP mask           844h.
//  0x84 32  DOUTEP1MSK dedicated EP msk            884h.
// Import:
//  stm32/o/l476xx/mmap
package usb

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	DSPD      DCFG_Bits = 0x03 << 0  //+ Device speed.
	NZLSOHSK  DCFG_Bits = 0x01 << 2  //+ Nonzero-length status OUT handshake.
	DAD       DCFG_Bits = 0x7F << 4  //+ Device address.
	PFIVL     DCFG_Bits = 0x03 << 11 //+ Periodic (micro)frame interval.
	PERSCHIVL DCFG_Bits = 0x03 << 24 //+ Periodic scheduling interval.
)

const (
	DSPDn      = 0
	NZLSOHSKn  = 2
	DADn       = 4
	PFIVLn     = 11
	PERSCHIVLn = 24
)

const (
	RWUSIG   DCTL_Bits = 0x01 << 0  //+ Remote wakeup signaling.
	SDIS     DCTL_Bits = 0x01 << 1  //+ Soft disconnect.
	GINSTS   DCTL_Bits = 0x01 << 2  //+ Global IN NAK status.
	GONSTS   DCTL_Bits = 0x01 << 3  //+ Global OUT NAK status.
	TCTL     DCTL_Bits = 0x07 << 4  //+ Test control.
	SGINAK   DCTL_Bits = 0x01 << 7  //+ Set global IN NAK.
	CGINAK   DCTL_Bits = 0x01 << 8  //+ Clear global IN NAK.
	SGONAK   DCTL_Bits = 0x01 << 9  //+ Set global OUT NAK.
	CGONAK   DCTL_Bits = 0x01 << 10 //+ Clear global OUT NAK.
	POPRGDNE DCTL_Bits = 0x01 << 11 //+ Power-on programming done.
)

const (
	RWUSIGn   = 0
	SDISn     = 1
	GINSTSn   = 2
	GONSTSn   = 3
	TCTLn     = 4
	SGINAKn   = 7
	CGINAKn   = 8
	SGONAKn   = 9
	CGONAKn   = 10
	POPRGDNEn = 11
)

const (
	SUSPSTS DSTS_Bits = 0x01 << 0   //+ Suspend status.
	ENUMSPD DSTS_Bits = 0x03 << 1   //+ Enumerated speed.
	EERR    DSTS_Bits = 0x01 << 3   //+ Erratic error.
	FNSOF   DSTS_Bits = 0x3FFF << 8 //+ Frame number of the received SOF.
)

const (
	SUSPSTSn = 0
	ENUMSPDn = 1
	EERRn    = 3
	FNSOFn   = 8
)

const (
	XFRCM     DIEPMSK_Bits = 0x01 << 0 //+ Transfer completed interrupt mask.
	EPDM      DIEPMSK_Bits = 0x01 << 1 //+ Endpoint disabled interrupt mask.
	TOM       DIEPMSK_Bits = 0x01 << 3 //+ Timeout condition mask (nonisochronous endpoints).
	ITTXFEMSK DIEPMSK_Bits = 0x01 << 4 //+ IN token received when TxFIFO empty mask.
	INEPNMM   DIEPMSK_Bits = 0x01 << 5 //+ IN token received with EP mismatch mask.
	INEPNEM   DIEPMSK_Bits = 0x01 << 6 //+ IN endpoint NAK effective mask.
	TXFURM    DIEPMSK_Bits = 0x01 << 8 //+ FIFO underrun mask.
	BIM       DIEPMSK_Bits = 0x01 << 9 //+ BNA interrupt mask.
)

const (
	XFRCMn     = 0
	EPDMn      = 1
	TOMn       = 3
	ITTXFEMSKn = 4
	INEPNMMn   = 5
	INEPNEMn   = 6
	TXFURMn    = 8
	BIMn       = 9
)

const (
	XFRCM   DOEPMSK_Bits = 0x01 << 0 //+ Transfer completed interrupt mask.
	EPDM    DOEPMSK_Bits = 0x01 << 1 //+ Endpoint disabled interrupt mask.
	STUPM   DOEPMSK_Bits = 0x01 << 3 //+ SETUP phase done mask.
	OTEPDM  DOEPMSK_Bits = 0x01 << 4 //+ OUT token received when endpoint disabled mask.
	B2BSTUP DOEPMSK_Bits = 0x01 << 6 //+ Back-to-back SETUP packets received mask.
	OPEM    DOEPMSK_Bits = 0x01 << 8 //+ OUT packet error mask.
	BOIM    DOEPMSK_Bits = 0x01 << 9 //+ BNA interrupt mask.
)

const (
	XFRCMn   = 0
	EPDMn    = 1
	STUPMn   = 3
	OTEPDMn  = 4
	B2BSTUPn = 6
	OPEMn    = 8
	BOIMn    = 9
)

const (
	IEPINT DAINT_Bits = 0xFFFF << 0  //+ IN endpoint interrupt bits.
	OEPINT DAINT_Bits = 0xFFFF << 16 //+ OUT endpoint interrupt bits.
)

const (
	IEPINTn = 0
	OEPINTn = 16
)

const (
	IEPM DAINTMSK_Bits = 0xFFFF << 0  //+ IN EP interrupt mask bits.
	OEPM DAINTMSK_Bits = 0xFFFF << 16 //+ OUT EP interrupt mask bits.
)

const (
	IEPMn = 0
	OEPMn = 16
)

const (
	VBUSDT DVBUSDIS_Bits = 0xFFFF << 0 //+ Device VBUS discharge time.
)

const (
	VBUSDTn = 0
)

const (
	DVBUSP DVBUSPULSE_Bits = 0xFFF << 0 //+ Device VBUS pulsing time.
)

const (
	DVBUSPn = 0
)

const (
	NONISOTHREN DTHRCTL_Bits = 0x01 << 0   //+ Nonisochronous IN endpoints threshold enable.
	ISOTHREN    DTHRCTL_Bits = 0x01 << 1   //+ ISO IN endpoint threshold enable.
	TXTHRLEN    DTHRCTL_Bits = 0x1FF << 2  //+ Transmit threshold length.
	RXTHREN     DTHRCTL_Bits = 0x01 << 16  //+ Receive threshold enable.
	RXTHRLEN    DTHRCTL_Bits = 0x1FF << 17 //+ Receive threshold length.
	ARPEN       DTHRCTL_Bits = 0x01 << 27  //+ Arbiter parking enable.
)

const (
	NONISOTHRENn = 0
	ISOTHRENn    = 1
	TXTHRLENn    = 2
	RXTHRENn     = 16
	RXTHRLENn    = 17
	ARPENn       = 27
)

const (
	INEPTXFEM DIEPEMPMSK_Bits = 0xFFFF << 0 //+ IN EP Tx FIFO empty interrupt mask bits.
)

const (
	INEPTXFEMn = 0
)

const (
	IEP1INT DEACHINT_Bits = 0x01 << 1  //+ IN endpoint 1interrupt bit.
	OEP1INT DEACHINT_Bits = 0x01 << 17 //+ OUT endpoint 1 interrupt bit.
)

const (
	IEP1INTn = 1
	OEP1INTn = 17
)