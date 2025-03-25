package cpu

// Abbreviations based on these: https://rgbds.gbdev.io/docs/v0.9.1/gbz80.7

// NOP
func NOP() {
	// no operation
}

// Load Instructions

// LD r8,r8 | LD r8,n8 | LD r8,[HL] | LD A,[r16] | LD A,[n16]
func LD_r8_n8(r *byte, n byte) {
	*r = n
}

// LD r16,n16 | LD SP,n16 | LD SP,HL
func LD_r16_n16(r *WordRegister, n uint16) {
	r.Set(n)
}

// LD [HL],r8 | LD [HL],n8 | LD [r16],A | LD [n16],A
func (cpu *CPU) LD_a16_n8(a uint16, n byte) {
	cpu.Bus.SetByte(a, n)
}

// LDH [n16],A
func (cpu *CPU) LDH_a16_r8(a uint16, r byte) {
	if a >= 0xFF00 {
		cpu.Bus.SetByte(a, r)
	}
}

// LDH [C],A
func (cpu *CPU) LDH_a8_r8(a byte, r byte) {
	cpu.Bus.SetByte(0xFF00+uint16(a), r)
}

// LDH A,[n16]
func (cpu *CPU) LDH_r8_a16(r *byte, a uint16) {
	if a >= 0xFF00 {
		*r = cpu.Bus.ReadByte(a)
	}
}

// LDH A,[C]
func (cpu *CPU) LDH_r8_a8(r *byte, a byte) {
	*r = cpu.Bus.ReadByte(0xFF00 + uint16(a))
}

// LD [HLI],A
func (cpu *CPU) LD_HLI_A() {
	cpu.Bus.SetByte(cpu.Registers.HL.Get(), cpu.Registers.A)
	cpu.Registers.HL.Set(cpu.Registers.HL.Get() + 1)
}

// LD [HLD],A
func (cpu *CPU) LD_HLD_A() {
	cpu.Bus.SetByte(cpu.Registers.HL.Get(), cpu.Registers.A)
	cpu.Registers.HL.Set(cpu.Registers.HL.Get() - 1)
}

// LD A,[HLI]
func (cpu *CPU) LD_A_HLI() {
	cpu.Registers.A = cpu.Bus.ReadByte(cpu.Registers.HL.Get())
	cpu.Registers.HL.Set(cpu.Registers.HL.Get() + 1)
}

// LD A,[HLD]
func (cpu *CPU) LD_A_HLD() {
	cpu.Registers.A = cpu.Bus.ReadByte(cpu.Registers.HL.Get())
	cpu.Registers.HL.Set(cpu.Registers.HL.Get() - 1)
}

// 8-Bit Arithmetic Instructions

// ADC A,r8 | ADC A,[HL] | ADC A,n8
func (cpu *CPU) ADC_A_n8(n byte) {
	var c byte
	if cpu.Registers.GetCarryFlag() {
		c = 1
	} else {
		c = 0
	}
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag((((cpu.Registers.A & 0x0F) + (n & 0x0F) + c) & 0x10) == 0x10)
	cpu.Registers.SetCarryFlag(uint16(cpu.Registers.A)+uint16(n)+uint16(c) > 0xFF)
	cpu.Registers.A += n + c
	cpu.Registers.SetZeroFlag(cpu.Registers.A == 0)
}

// ADD A,r8 | ADD A,[HL] | ADD A,n8
func (cpu *CPU) ADD_A_n8(n byte) {
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag((((cpu.Registers.A & 0x0F) + (n & 0x0F)) & 0x10) == 0x10)
	cpu.Registers.SetCarryFlag(uint16(cpu.Registers.A)+uint16(n) > 0xFF)
	cpu.Registers.A += n
	cpu.Registers.SetZeroFlag(cpu.Registers.A == 0)
}

// CP A,r8 | CP A,[HL] | CP A,n8
func (cpu *CPU) CP_A_n8(n byte) {
	cpu.Registers.SetSubtractionFlag(true)
	cpu.Registers.SetHalfCarryFlag((n & 0x0F) > (cpu.Registers.A & 0x0F))
	cpu.Registers.SetCarryFlag(n > cpu.Registers.A)
	cpu.Registers.SetZeroFlag(cpu.Registers.A-n == 0)
}

// DEC r8
func (cpu *CPU) DEC_r8(r *byte) {
	cpu.Registers.SetHalfCarryFlag((*r & 0x0F) == 0)
	cpu.Registers.SetSubtractionFlag(true)
	(*r)--
	cpu.Registers.SetZeroFlag(*r == 0)
}

// DEC [HL]
func (cpu *CPU) DEC_a16(a uint16) {
	currVal := cpu.Bus.ReadByte(a)
	cpu.Registers.SetHalfCarryFlag((currVal & 0x0F) == 0)
	cpu.Registers.SetSubtractionFlag(true)
	cpu.Bus.SetByte(a, currVal-1)
	cpu.Registers.SetZeroFlag(currVal-1 == 0)
}

// INC r8
func (cpu *CPU) INC_r8(r *byte) {
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag((*r)&0x0F == 0x0F)
	(*r)++
	cpu.Registers.SetZeroFlag(*r == 0)
}

// INC [HL]
func (cpu *CPU) INC_a16(a uint16) {
	currVal := cpu.Bus.ReadByte(a)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(currVal&0x0F == 0x0F)
	cpu.Bus.SetByte(a, currVal+1)
	cpu.Registers.SetZeroFlag(currVal+1 == 0)
}

// SBC A,r8 | SBC A,[HL] | SBC A,n8
func (cpu *CPU) SBC_A_n8(n byte) {
	var c byte
	if cpu.Registers.GetCarryFlag() {
		c = 1
	} else {
		c = 0
	}
	cpu.Registers.SetHalfCarryFlag((cpu.Registers.A & 0x0F) < ((n + c) & 0x0F))
	cpu.Registers.SetCarryFlag((n + c) > cpu.Registers.A)
	cpu.Registers.SetSubtractionFlag(true)
	cpu.Registers.A -= n + c
	cpu.Registers.SetZeroFlag(cpu.Registers.A == 0)
}

// SUB A,r8 | SUB A,[HL] | SUB A,n8
func (cpu *CPU) SUB_A_n8(n byte) {
	cpu.Registers.SetHalfCarryFlag((cpu.Registers.A & 0x0F) < (n & 0x0F))
	cpu.Registers.SetCarryFlag(n > cpu.Registers.A)
	cpu.Registers.SetSubtractionFlag(true)
	cpu.Registers.A -= n
	cpu.Registers.SetZeroFlag(cpu.Registers.A == 0)
}

// 16-Bit Arithmetic Instructions

// ADD HL,r16
func (cpu *CPU) ADD_HL_n16(n uint16) {
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag((cpu.Registers.HL.Get()&0x0FFF)+(n&0x0FFF)&0x1FFF == 0x1FFF)
	cpu.Registers.SetCarryFlag(int(cpu.Registers.HL.Get())+int(n) > 0xFFFF)
	cpu.Registers.HL.Set(cpu.Registers.HL.Get() + n)
	cpu.Registers.SetZeroFlag(cpu.Registers.HL.Get() == 0)
}

// DEC r16
func DEC_r16(r *WordRegister) {
	r.Set(r.Get() - 1)
}

// INC r16
func INC_r16(r *WordRegister) {
	r.Set(r.Get() + 1)
}

// Bitwise Logic Instructions

// AND A,r8 | AND A,[HL] | AND A,n8
func (cpu *CPU) AND_A_n8(n byte) {
	cpu.Registers.A &= n
	cpu.Registers.SetZeroFlag(cpu.Registers.A == 0)
	cpu.Registers.SetCarryFlag(false)
	cpu.Registers.SetHalfCarryFlag(true)
	cpu.Registers.SetSubtractionFlag(false)
}

// CPL
func (cpu *CPU) CPL() {
	cpu.Registers.A = ^cpu.Registers.A
	cpu.Registers.SetSubtractionFlag(true)
	cpu.Registers.SetHalfCarryFlag(true)
}

// OR A,r8 | OR A,[HL] | OR A,n8
func (cpu *CPU) OR_A_n8(n byte) {
	cpu.Registers.A |= n
	cpu.Registers.SetZeroFlag(cpu.Registers.A == 0)
	cpu.Registers.SetCarryFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	cpu.Registers.SetSubtractionFlag(false)
}

// XOR A,r8 | XOR A,[HL] | XOR A,n8
func (cpu *CPU) XOR_A_n8(n byte) {
	cpu.Registers.A ^= n
	cpu.Registers.SetZeroFlag(cpu.Registers.A == 0)
	cpu.Registers.SetCarryFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	cpu.Registers.SetSubtractionFlag(false)
}

// Stack Manipulation Instructions

// LD [n16],SP
func (cpu *CPU) LD_a16_n16(a uint16, n uint16) {
	high := byte(n >> 8)
	low := byte(n)
	cpu.Bus.SetByte(a, low)
	cpu.Bus.SetByte(a+1, high)
}
