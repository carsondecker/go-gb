package mmu

type MemoryBus struct {
	Memory [0xFFFF]byte
}

func (m *MemoryBus) ReadByte(address uint16) byte {
	return m.Memory[address]
}

func (m *MemoryBus) SetByte(address uint16, val byte) {
	m.Memory[address] = val
}
