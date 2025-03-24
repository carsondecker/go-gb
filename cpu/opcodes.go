package cpu

// Abbreviations based on these: https://rgbds.gbdev.io/docs/v0.9.1/gbz80.7

// NOP
func NOP() {
	// no operation
}

// Load instructions

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
	cpu.Registers.HL.Set(cpu.Registers.HL.Get() + 1)
}

// Stack Manipulation Instructions

// LD [n16],SP
func (cpu *CPU) LD_a16_n16(a uint16, n uint16) {
	high := byte(n >> 8)
	low := byte(n)
	cpu.Bus.SetByte(a, low)
	cpu.Bus.SetByte(a+1, high)
}
