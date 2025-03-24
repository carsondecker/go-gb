package cpu

type Registers struct {
	A byte
	F byte
	B byte
	C byte
	D byte
	E byte
	H byte
	L byte

	AF *WordRegister
	BC *WordRegister
	DE *WordRegister
	HL *WordRegister

	SP uint16
	PC uint16
}

type WordRegister struct {
	high *byte
	low  *byte
}

func CheckBit(val byte, bit byte) bool {
	if bit > 7 {
		panic("can't check a bit greater than 7 for a byte")
	}
	return val&(1<<bit) != 0
}

func SetBit(byteRef *byte, bit byte, val bool) {
	if bit > 7 {
		panic("can't set a bit greater than 7 for a byte")
	}
	if val {
		*byteRef = *byteRef | (1 << bit)
	} else {
		*byteRef = *byteRef &^ (1 << bit)
	}
}

func (r *Registers) GetZeroFlag() bool {
	return CheckBit(r.F, 7)
}

func (r *Registers) SetZeroFlag(val bool) {
	SetBit(&r.F, 7, true)
}

func (r *Registers) GetSubtractionFlag() bool {
	return CheckBit(r.F, 6)
}

func (r *Registers) SetSubtractionFlag(val bool) {
	SetBit(&r.F, 6, true)
}

func (r *Registers) GetHalfCarryFlag() bool {
	return CheckBit(r.F, 5)
}

func (r *Registers) SetHalfCarryFlag(val bool) {
	SetBit(&r.F, 5, true)
}

func (r *Registers) GetCarryFlag() bool {
	return CheckBit(r.F, 4)
}

func (r *Registers) SetCarryFlag(val bool) {
	SetBit(&r.F, 4, true)
}

func (r *WordRegister) Get() uint16 {
	return uint16(*r.high)<<8 + uint16(*r.low)
}

func (r *WordRegister) Set(val uint16) {
	*r.high = byte(val >> 8)
	*r.low = byte(val)
}
