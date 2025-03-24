package cpu

// Abbreviations based on these: https://rgbds.gbdev.io/docs/v0.9.1/gbz80.7

func NOP() {
	// no operation
}

// LD Operations

func LD_r8_n8(r *byte, n byte) {
	*r = n
}

func LD_r16_n16(r *WordRegister, n uint16) {
	r.Set(n)
}

func (cpu *CPU) LD_ar16_n8(r *WordRegister, n byte) {
	cpu.Bus.SetByte(r.Get(), n)
}

func (cpu *CPU) LD_an16_n16(a uint16, n uint16) {
	high := byte(n >> 8)
	low := byte(n)
	cpu.Bus.SetByte(a, low)
	cpu.Bus.SetByte(a+1, high)
}

// INC Operations

func (cpu *CPU) INC_r8(r *byte) {
	(*r)++
	if *r == 0 {
		cpu.Registers.SetZeroFlag(true)
		cpu.Registers.SetHalfCarryFlag(true)
	}
	cpu.Registers.SetSubtractionFlag(false)
}

func INC_r16(r *WordRegister) {
	if *r.low == 255 {
		(*r.high)++
	}
	(*r.low)++
}

// DEC Operations

func (cpu *CPU) DEC_r8(r *byte) {
	(*r)--
	if *r == 0 {
		cpu.Registers.SetZeroFlag(true)
	} else if *r == 255 {
		cpu.Registers.SetHalfCarryFlag(true)
	}
	cpu.Registers.SetSubtractionFlag(true)
}

// Rotate Operations

func (cpu *CPU) RLC(r *byte) {
	leftBit := CheckBit(*r, 7)
	cpu.Registers.SetCarryFlag(leftBit)
	*r = (*r) << 1
	SetBit(r, 0, leftBit)
	if *r == 0 {
		cpu.Registers.SetZeroFlag(true)
	}
}
