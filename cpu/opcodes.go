package cpu

import "fmt"

// Abbreviations based on these: https://rgbds.gbdev.io/docs/v0.9.1/gbz80.7

func NOP() {
	// no operation
	fmt.Println("No operation performed")
}

// LD Operations

func LD_r8_n8(r *byte, n byte) {
	*r = n
	fmt.Printf("Loaded %d into register at address %p\n", n, r)
}

func LD_r16_n16(r *WordRegister, n uint16) {
	r.Set(n)
	fmt.Printf("Loaded %d into register at address %p\n", n, r)
}

func (cpu *CPU) LD_rADDR_n8(r *WordRegister, n byte) {
	cpu.Bus.SetByte(r.Get(), n)
}
