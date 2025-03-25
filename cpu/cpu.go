package cpu

import "github.com/carsondecker/go-gb/mmu"

type CPU struct {
	Registers *Registers
	Bus       *mmu.MemoryBus
	IME       bool
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
		cpu.LD_a16_n8(cpu.Registers.BC.Get(), cpu.Registers.A)
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
