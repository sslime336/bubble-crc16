package main

func Checksum(alg Algorithm, byts []byte) uint16 {
	var crc uint16
	if alg.Refin {
		crc = reverseUint16(alg.Init)
	} else {
		crc = alg.Init
	}
	poly := uint16(alg.Poly)

	for _, byt := range byts {
		if alg.Refin {
			byt = reverseByte(byt)
		}
		crc ^= uint16(byt) << 8
		for range 8 {
			if crc&0x8000 > 0 {
				crc = (crc << 1) ^ poly
			} else {
				crc <<= 1
			}
		}

	}

	if alg.Refout {
		crc = reverseUint16(crc)
	}

	return crc ^ uint16(alg.Xorout)
}

func reverseUint16(num uint16) (res uint16) {
	for i := range 8 {
		var f0, f1 uint16
		if num&(1<<i) > 0 {
			f0 = 1
		}
		if num&(1<<(15-i)) > 0 {
			f1 = 1
		}
		res |= (f0 << (15 - i)) + (f1 << i)
	}
	return
}

func reverseByte(b byte) byte {
	b = (b&0xF0)>>4 | (b&0x0F)<<4
	b = (b&0xCC)>>2 | (b&0x33)<<2
	b = (b&0xAA)>>1 | (b&0x55)<<1
	return b
}

type Algorithm struct {
	Poly   int
	Init   uint16
	Refin  bool
	Refout bool
	Xorout int
	Check  int
}

var CRC16Algorithm = []Algorithm{
	{Poly: 0x8005, Init: 0x0000, Refin: true, Refout: true, Xorout: 0x0000, Check: 0xbb3d},
	{Poly: 0xc867, Init: 0xffff, Refin: false, Refout: false, Xorout: 0x0000, Check: 0x4c06},
	{Poly: 0x8005, Init: 0xffff, Refin: false, Refout: false, Xorout: 0x0000, Check: 0xaee7},
	{Poly: 0x8005, Init: 0x800d, Refin: false, Refout: false, Xorout: 0x0000, Check: 0x9ecf},
	{Poly: 0x0589, Init: 0x0000, Refin: false, Refout: false, Xorout: 0x0001, Check: 0x007e},
	{Poly: 0x0589, Init: 0x0000, Refin: false, Refout: false, Xorout: 0x0000, Check: 0x007f},
	{Poly: 0x3d65, Init: 0x0000, Refin: true, Refout: true, Xorout: 0xffff, Check: 0xea82},
	{Poly: 0x3d65, Init: 0x0000, Refin: false, Refout: false, Xorout: 0xffff, Check: 0xc2b7},
	{Poly: 0x1021, Init: 0xffff, Refin: false, Refout: false, Xorout: 0xffff, Check: 0xd64e},
	{Poly: 0x1021, Init: 0x0000, Refin: false, Refout: false, Xorout: 0xffff, Check: 0xce3c},
	{Poly: 0x1021, Init: 0xffff, Refin: false, Refout: false, Xorout: 0x0000, Check: 0x29b1},
	{Poly: 0x1021, Init: 0xffff, Refin: true, Refout: true, Xorout: 0xffff, Check: 0x906e},
	{Poly: 0x1021, Init: 0xc6c6, Refin: true, Refout: true, Xorout: 0x0000, Check: 0xbf05},
	{Poly: 0x1021, Init: 0x0000, Refin: true, Refout: true, Xorout: 0x0000, Check: 0x2189},
	{Poly: 0x6f63, Init: 0x0000, Refin: false, Refout: false, Xorout: 0x0000, Check: 0xbdf4},
	{Poly: 0x5935, Init: 0xffff, Refin: false, Refout: false, Xorout: 0x0000, Check: 0x772b},
	{Poly: 0x8005, Init: 0x0000, Refin: true, Refout: true, Xorout: 0xffff, Check: 0x44c2},
	{Poly: 0x1021, Init: 0xffff, Refin: true, Refout: true, Xorout: 0x0000, Check: 0x6f91},
	{Poly: 0x8005, Init: 0xffff, Refin: true, Refout: true, Xorout: 0x0000, Check: 0x4b37},
	{Poly: 0x080b, Init: 0xffff, Refin: true, Refout: true, Xorout: 0x0000, Check: 0xa066},
	{Poly: 0x5935, Init: 0x0000, Refin: false, Refout: false, Xorout: 0x0000, Check: 0x5d38},
	{Poly: 0x755b, Init: 0x0000, Refin: false, Refout: false, Xorout: 0x0000, Check: 0x20fe},
	{Poly: 0x1dcf, Init: 0xffff, Refin: false, Refout: false, Xorout: 0xffff, Check: 0xa819},
	{Poly: 0x1021, Init: 0xb2aa, Refin: true, Refout: true, Xorout: 0x0000, Check: 0x63d0},
	{Poly: 0x1021, Init: 0x1d0f, Refin: false, Refout: false, Xorout: 0x0000, Check: 0xe5cc},
	{Poly: 0x8bb7, Init: 0x0000, Refin: false, Refout: false, Xorout: 0x0000, Check: 0xd0db},
	{Poly: 0xa097, Init: 0x0000, Refin: false, Refout: false, Xorout: 0x0000, Check: 0x0fb3},
	{Poly: 0x1021, Init: 0x89ec, Refin: true, Refout: true, Xorout: 0x0000, Check: 0x26b1},
	{Poly: 0x8005, Init: 0x0000, Refin: false, Refout: false, Xorout: 0x0000, Check: 0xfee8},
	{Poly: 0x8005, Init: 0xffff, Refin: true, Refout: true, Xorout: 0xffff, Check: 0xb4c8},
	{Poly: 0x1021, Init: 0x0000, Refin: false, Refout: false, Xorout: 0x0000, Check: 0x31c3},
}

const SupportedCRC16ModelLen = 31

var CrcModelMap = func() map[CRC16Model]Algorithm {
	cmmap := make(map[CRC16Model]Algorithm, SupportedCRC16ModelLen)
	for modelIndex := range SupportedCRC16ModelLen {
		cmmap[CRC16Model(modelIndex)] = CRC16Algorithm[modelIndex]
	}

	return cmmap
}()

//go:generate stringer -type=CRC16Model
type CRC16Model int

const (
	CRC_16_ARC CRC16Model = iota
	CRC_16_CDMA2000
	CRC_16_CMS
	CRC_16_DDS_110
	CRC_16_DECT_R
	CRC_16_DECT_X
	CRC_16_DNP
	CRC_16_EN_13757
	CRC_16_GENIBUS
	CRC_16_GSM
	CRC_16_IBM_3740
	CRC_16_IBM_SDLC
	CRC_16_ISO_IEC_14443_3_A
	CRC_16_KERMIT
	CRC_16_LJ1200
	CRC_16_M17
	CRC_16_MAXIM_DOW
	CRC_16_MCRF4XX
	CRC_16_MODBUS
	CRC_16_NRSC_5
	CRC_16_OPENSAFETY_A
	CRC_16_OPENSAFETY_B
	CRC_16_PROFIBUS
	CRC_16_RIELLO
	CRC_16_SPI_FUJITSU
	CRC_16_T10_DIF
	CRC_16_TELEDISK
	CRC_16_TMS37157
	CRC_16_UMTS
	CRC_16_USB
	CRC_16_XMODEM
)
