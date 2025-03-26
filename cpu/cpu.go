package cpu

import "github.com/carsondecker/go-gb/mmu"

type CPU struct {
	Registers *Registers
	Bus       *mmu.MemoryBus
	IME       bool
	Stopped   bool
	Halted    bool
}

// TODO: Make HALT and STOP work in the step command
func (cpu *CPU) Step() {
	nextByte := cpu.Bus.ReadByte(cpu.Registers.PC)
	cpu.Registers.PC++
	cpu.execute(nextByte)
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

func (cpu *CPU) handleInterrupts() {
	if !cpu.IME {
		return
	}

	callableInterrupts := cpu.Bus.ReadByte(0xFFFF) & cpu.Bus.ReadByte(0xFF0F)
	if callableInterrupts&0x01 != 0 {
		// VBlank
		cpu.RST(0x40)
		cpu.Bus.SetByte(0xFF0F, cpu.Bus.ReadByte(0xFF0F)^0x01)
	} else if callableInterrupts&0x02 != 0 {
		// LCD
		cpu.RST(0x48)
		cpu.Bus.SetByte(0xFF0F, cpu.Bus.ReadByte(0xFF0F)^0x02)
	} else if callableInterrupts&0x04 != 0 {
		// Timer
		cpu.RST(0x50)
		cpu.Bus.SetByte(0xFF0F, cpu.Bus.ReadByte(0xFF0F)^0x04)
	} else if callableInterrupts&0x08 != 0 {
		// Serial
		cpu.RST(0x58)
		cpu.Bus.SetByte(0xFF0F, cpu.Bus.ReadByte(0xFF0F)^0x08)
	} else if callableInterrupts&0x10 != 0 {
		// JoyPad
		cpu.RST(0x60)
		cpu.Bus.SetByte(0xFF0F, cpu.Bus.ReadByte(0xFF0F)^0x10)
	} else {
		return
	}
	cpu.IME = false
}
