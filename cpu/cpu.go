package cpu

import "github.com/carsondecker/go-gb/mmu"

type CPU struct {
	Registers *Registers
	Bus       *mmu.MemoryBus
}

func (cpu *CPU) ExecuteInstruction() {
	nextByte := cpu.Bus.ReadByte(cpu.Registers.PC)
	cpu.Registers.PC++
	switch nextByte {
	case 0x00:
		NOP()
	case 0x01:
		data := cpu.getNextWord()
		LD_r16_n16(cpu.Registers.BC, data)
	case 0x02:
		cpu.LD_ar16_n8(cpu.Registers.BC, cpu.Registers.A)
	case 0x03:
		INC_r16(cpu.Registers.BC)
	case 0x04:
		cpu.INC_r8(&cpu.Registers.B)
	case 0x05:
		cpu.DEC_r8(&cpu.Registers.B)
	case 0x06:
		data := cpu.getNextByte()
		LD_r8_n8(&cpu.Registers.B, data)
	case 0x07:
		cpu.RLC(&cpu.Registers.A)
	case 0x08:
		data := cpu.getNextWord()
		cpu.LD_an16_n16(data, cpu.Registers.SP)
	}
}

func (cpu *CPU) getNextWord() uint16 {
	high := cpu.Bus.ReadByte(cpu.Registers.PC)
	cpu.Registers.PC++
	low := cpu.Bus.ReadByte(cpu.Registers.PC)
	cpu.Registers.PC++
	return uint16(high)<<8 + uint16(low)
}

func (cpu *CPU) getNextByte() byte {
	data := cpu.Bus.ReadByte(cpu.Registers.PC)
	cpu.Registers.PC++
	return data
}
