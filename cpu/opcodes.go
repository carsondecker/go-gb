package cpu

// Abbreviations based on these: https://rgbds.gbdev.io/docs/v0.9.1/gbz80.7

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

// ADD HL,r16 | ADD HL,SP
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

// Bit Flag Instructions

// BIT u3,r8 | BIT u3,[HL]
func (cpu *CPU) BIT_n8_n8(bit byte, n byte) {
	cpu.Registers.SetZeroFlag(!CheckBit(n, bit))
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(true)
}

// RES u3,r8
func RES_n8_r8(bit byte, r *byte) {
	SetBit(r, bit, false)
}

// RES u3,[HL]
func (cpu *CPU) RES_n8_a16(bit byte, a uint16) {
	cpu.Bus.SetByte(a, cpu.Bus.ReadByte(a)&^(1<<bit))
}

// SET u3,r8
func SET_n8_r8(bit byte, r *byte) {
	SetBit(r, bit, true)
}

// SET u3,[HL]
func (cpu *CPU) SET_n8_a16(bit byte, a uint16) {
	cpu.Bus.SetByte(a, cpu.Bus.ReadByte(a)|(1<<bit))
}

// Bit Shift Instructions

// RL r8 | RLA
func (cpu *CPU) RL_r8(r *byte) {
	var c byte
	if cpu.Registers.GetCarryFlag() {
		c = 1
	} else {
		c = 0
	}
	cpu.Registers.SetCarryFlag(*r&0x80 != 0)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	*r = ((*r) << 1) | c
	cpu.Registers.SetZeroFlag(*r == 0)
}

// RL [HL]
func (cpu *CPU) RL_a16(a uint16) {
	var c byte
	if cpu.Registers.GetCarryFlag() {
		c = 1
	} else {
		c = 0
	}
	val := cpu.Bus.ReadByte(a)
	cpu.Registers.SetCarryFlag(val&0x80 != 0)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	cpu.Bus.SetByte(a, (val<<1)|c)
	cpu.Registers.SetZeroFlag(cpu.Bus.ReadByte(a) == 0)
}

// RLC r8 | RLCA
func (cpu *CPU) RLC_r8(r *byte) {
	cpu.Registers.SetCarryFlag(*r&0x80 != 0)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	*r = *r<<1 | *r>>7
	cpu.Registers.SetZeroFlag(*r == 0)
}

// RLC [HL]
func (cpu *CPU) RLC_a16(a uint16) {
	val := cpu.Bus.ReadByte(a)
	cpu.Registers.SetCarryFlag(val&0x80 != 0)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	cpu.Bus.SetByte(a, val<<1|val>>7)
	cpu.Registers.SetZeroFlag(cpu.Bus.ReadByte(a) == 0)
}

// RR r8 | RRA
func (cpu *CPU) RR_r8(r *byte) {
	var c byte
	if cpu.Registers.GetCarryFlag() {
		c = 1
	} else {
		c = 0
	}
	cpu.Registers.SetCarryFlag(*r&0x01 != 0)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	*r = ((*r) >> 1) | c<<7
	cpu.Registers.SetZeroFlag(*r == 0)
}

// RR [HL]
func (cpu *CPU) RR_a16(a uint16) {
	var c byte
	if cpu.Registers.GetCarryFlag() {
		c = 1
	} else {
		c = 0
	}
	val := cpu.Bus.ReadByte(a)
	cpu.Registers.SetCarryFlag(val&0x01 != 0)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	cpu.Bus.SetByte(a, (val>>1)|c<<7)
	cpu.Registers.SetZeroFlag(cpu.Bus.ReadByte(a) == 0)
}

// RRC r8 | RRCA
func (cpu *CPU) RRC_r8(r *byte) {
	cpu.Registers.SetCarryFlag(*r&0x01 != 0)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	*r = (*r)>>1 | (*r)<<7
	cpu.Registers.SetZeroFlag(*r == 0)
}

// RRC a16
func (cpu *CPU) RRC_a16(a uint16) {
	val := cpu.Bus.ReadByte(a)
	cpu.Registers.SetCarryFlag(val&0x01 != 0)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	cpu.Bus.SetByte(a, val>>1|val<<7)
	cpu.Registers.SetZeroFlag(cpu.Bus.ReadByte(a) == 0)
}

// SLA r8
func (cpu *CPU) SLA_r8(r *byte) {
	cpu.Registers.SetCarryFlag((*r)&0x80 != 0)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	*r = (*r) << 1
	cpu.Registers.SetZeroFlag(*r == 0)
}

// SLA [HL]
func (cpu *CPU) SLA_a16(a uint16) {
	val := cpu.Bus.ReadByte(a)
	cpu.Registers.SetCarryFlag(val&0x80 != 0)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	cpu.Bus.SetByte(a, val<<1)
	cpu.Registers.SetZeroFlag(cpu.Bus.ReadByte(a) == 0)
}

// SRA r8
func (cpu *CPU) SRA_r8(r *byte) {
	cpu.Registers.SetCarryFlag((*r)&0x01 != 0)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	*r = ((*r) >> 1) | ((*r) & 0x80)
	cpu.Registers.SetZeroFlag(*r == 0)
}

// SRA [HL]
func (cpu *CPU) SRA_a16(a uint16) {
	val := cpu.Bus.ReadByte(a)
	cpu.Registers.SetCarryFlag(val&0x80 != 0)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	cpu.Bus.SetByte(a, (val>>1)|(val&0x80))
	cpu.Registers.SetZeroFlag(cpu.Bus.ReadByte(a) == 0)
}

// SRL r8
func (cpu *CPU) SRL_r8(r *byte) {
	cpu.Registers.SetCarryFlag((*r)&0x01 != 0)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	*r = (*r) >> 1
	cpu.Registers.SetZeroFlag(*r == 0)
}

// SRL [HL]
func (cpu *CPU) SRL_a16(a uint16) {
	val := cpu.Bus.ReadByte(a)
	cpu.Registers.SetCarryFlag(val&0x80 != 0)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	cpu.Bus.SetByte(a, val>>1)
	cpu.Registers.SetZeroFlag(cpu.Bus.ReadByte(a) == 0)
}

// SWAP r8
func (cpu *CPU) SWAP_r8(r *byte) {
	cpu.Registers.SetCarryFlag(false)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	*r = (((*r) & 0xF0) >> 4) | (((*r) & 0x0F) << 4)
	cpu.Registers.SetZeroFlag(*r == 0)
}

// SWAP [HL]
func (cpu *CPU) SWAP_a16(a uint16) {
	val := cpu.Bus.ReadByte(a)
	cpu.Registers.SetCarryFlag(false)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(false)
	cpu.Bus.SetByte(a, (((val)&0xF0)>>4)|(((val)&0x0F)<<4))
	cpu.Registers.SetZeroFlag(cpu.Bus.ReadByte(a) == 0)
}

// Jumps and Subroutine Instructions

// CALL n16
func (cpu *CPU) CALL_n16(n uint16) {
	cpu.Registers.SP--
	cpu.Bus.SetByte(cpu.Registers.SP, byte(cpu.Registers.PC))
	cpu.Registers.SP--
	cpu.Bus.SetByte(cpu.Registers.SP, byte(cpu.Registers.PC>>4))
	cpu.Registers.PC = n
}

// CALL Z,n16
func (cpu *CPU) CALL_Z_n16(n uint16) {
	if cpu.Registers.GetZeroFlag() {
		cpu.CALL_n16(n)
	}
}

// CALL NZ,n16
func (cpu *CPU) CALL_NZ_n16(n uint16) {
	if !cpu.Registers.GetZeroFlag() {
		cpu.CALL_n16(n)
	}
}

// CALL C,n16
func (cpu *CPU) CALL_C_n16(n uint16) {
	if cpu.Registers.GetCarryFlag() {
		cpu.CALL_n16(n)
	}
}

// CALL NC,n16
func (cpu *CPU) CALL_NC_n16(n uint16) {
	if !cpu.Registers.GetCarryFlag() {
		cpu.CALL_n16(n)
	}
}

// JP HL | JP n16
func (cpu *CPU) JP_n16(n uint16) {
	cpu.Registers.PC = n
}

// JP Z,n16
func (cpu *CPU) JP_Z_n16(n uint16) {
	if cpu.Registers.GetZeroFlag() {
		cpu.Registers.PC = n
	}
}

// JP NZ,n16
func (cpu *CPU) JP_NZ_n16(n uint16) {
	if !cpu.Registers.GetZeroFlag() {
		cpu.Registers.PC = n
	}
}

// JP C,n16
func (cpu *CPU) JP_C_n16(n uint16) {
	if cpu.Registers.GetCarryFlag() {
		cpu.Registers.PC = n
	}
}

// JP NC,n16
func (cpu *CPU) JP_NC_n16(n uint16) {
	if !cpu.Registers.GetCarryFlag() {
		cpu.Registers.PC = n
	}
}

// JR n16? JR e8?
// TODO: Check this one again
func (cpu *CPU) JR_e8(e int8) {
	cpu.Registers.PC = uint16(int16(cpu.Registers.PC) + int16(e))
}

// JR Z,e8
func (cpu *CPU) JR_Z_e8(e int8) {
	if cpu.Registers.GetZeroFlag() {
		cpu.JR_e8(e)
	}
}

// JR NZ,e8
func (cpu *CPU) JR_NZ_e8(e int8) {
	if !cpu.Registers.GetZeroFlag() {
		cpu.JR_e8(e)
	}
}

// JR C,e8
func (cpu *CPU) JR_C_e8(e int8) {
	if cpu.Registers.GetCarryFlag() {
		cpu.JR_e8(e)
	}
}

// JR NC,e8
func (cpu *CPU) JR_NC_e8(e int8) {
	if !cpu.Registers.GetCarryFlag() {
		cpu.JR_e8(e)
	}
}

// RET
func (cpu *CPU) RET() {
	low := cpu.Bus.ReadByte(cpu.Registers.SP)
	cpu.Registers.SP++
	high := cpu.Bus.ReadByte(cpu.Registers.SP)
	cpu.Registers.SP++
	cpu.Registers.PC = uint16(high)<<8 | uint16(low)
}

// RET Z
func (cpu *CPU) RET_Z() {
	if cpu.Registers.GetZeroFlag() {
		cpu.RET()
	}
}

// RET NZ
func (cpu *CPU) RET_NZ() {
	if !cpu.Registers.GetZeroFlag() {
		cpu.RET()
	}
}

// RET C
func (cpu *CPU) RET_C() {
	if cpu.Registers.GetCarryFlag() {
		cpu.RET()
	}
}

// RET NC
func (cpu *CPU) RET_NC() {
	if !cpu.Registers.GetCarryFlag() {
		cpu.RET()
	}
}

// RETI
func (cpu *CPU) RETI() {
	cpu.EI()
	cpu.RET()
}

// RST vec
func (cpu *CPU) RST(v byte) {
	cpu.CALL_n16(uint16(v))
}

// Carry Flag Instructions

// CCF
func (cpu *CPU) CCF() {
	cpu.Registers.SetCarryFlag(!cpu.Registers.GetCarryFlag())
}

// SCF
func (cpu *CPU) SCF() {
	cpu.Registers.SetCarryFlag(true)
}

// Stack Manipulation Instructions

// ADD SP,e8
func (cpu *CPU) ADD_SP_e8(e int8) {
	cpu.Registers.SetZeroFlag(false)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(((cpu.Registers.SP & 0x0F) + (uint16(e) & 0x0F)) > 0x0F)
	cpu.Registers.SetCarryFlag(((cpu.Registers.SP & 0xFF) + (uint16(e) & 0xFF)) > 0xFF)
	cpu.Registers.SP = uint16(int16(cpu.Registers.SP) + int16(e))
}

// DEC SP
func (cpu *CPU) DEC_SP() {
	cpu.Registers.SP--
}

// INC SP
func (cpu *CPU) INC_SP() {
	cpu.Registers.SP++
}

// LD SP,n16 | LD SP,HL
func (cpu *CPU) LD_SP_n16(n uint16) {
	cpu.Registers.SP = n
}

// LD [n16],SP
func (cpu *CPU) LD_a16_SP(a uint16, n uint16) {
	high := byte(n >> 8)
	low := byte(n)
	cpu.Bus.SetByte(a, low)
	cpu.Bus.SetByte(a+1, high)
}

// LD HL,SP+e8
func (cpu *CPU) LD_SP_e8(e int8) {
	cpu.Registers.SetZeroFlag(false)
	cpu.Registers.SetSubtractionFlag(false)
	cpu.Registers.SetHalfCarryFlag(((cpu.Registers.SP & 0x0F) + (uint16(e) & 0x0F)) > 0x0F)
	cpu.Registers.SetCarryFlag(((cpu.Registers.SP & 0xFF) + (uint16(e) & 0xFF)) > 0xFF)
	cpu.Registers.HL.Set(uint16(int16(cpu.Registers.SP) + int16(e)))
}

// POP AF | POP r16
func (cpu *CPU) POP_r16(r *WordRegister) {
	*r.low = cpu.Bus.ReadByte(cpu.Registers.SP)
	cpu.Registers.SP++
	*r.high = cpu.Bus.ReadByte(cpu.Registers.SP)
	cpu.Registers.SP++
}

// PUSH AF | PUSH r16
func (cpu *CPU) PUSH_r16(r *WordRegister) {
	cpu.Registers.SP--
	cpu.Bus.SetByte(cpu.Registers.SP, *r.high)
	cpu.Registers.SP--
	cpu.Bus.SetByte(cpu.Registers.SP, *r.high)
}

// Interrupt-Related Instructions

// DI
func (cpu *CPU) DI() {
	cpu.IME = false
}

// EI
func (cpu *CPU) EI() {
	cpu.IME = true
}

// Other

// DAA
func (cpu *CPU) DAA() {
	if cpu.Registers.GetSubtractionFlag() {
		var adj byte = 0
		if cpu.Registers.GetHalfCarryFlag() {
			adj += 0x6
		}
		if cpu.Registers.GetCarryFlag() {
			adj += 0x60
		}
		cpu.Registers.A -= adj
	} else {
		var adj byte = 0
		if cpu.Registers.GetHalfCarryFlag() || (cpu.Registers.A&0xF) > 0x9 {
			adj += 0x6
		}
		if cpu.Registers.GetCarryFlag() || cpu.Registers.A > 0x99 {
			adj += 0x60
			cpu.Registers.SetCarryFlag(true)
		}
		cpu.Registers.A += adj
	}
	cpu.Registers.SetHalfCarryFlag(false)
	cpu.Registers.SetZeroFlag(cpu.Registers.A == 0)
}

// NOP
func NOP() {
	// no operation
}
