package cpu

func (cpu *CPU) execute(nextByte byte) {
	switch nextByte {
	case 0x00:
		NOP()
	case 0x01:
		LD_r16_n16(cpu.Registers.BC, cpu.getNextWord())
	case 0x02:
		cpu.LD_a16_n8(cpu.Registers.BC.Get(), cpu.Registers.A)
	case 0x03:
		INC_r16(cpu.Registers.BC)
	case 0x04:
		cpu.INC_r8(&cpu.Registers.B)
	case 0x05:
		cpu.DEC_r8(&cpu.Registers.B)
	case 0x06:
		LD_r8_n8(&cpu.Registers.B, cpu.getNextByte())
	case 0x07:
		cpu.RLC_r8(&cpu.Registers.A)
	case 0x08:
		cpu.LD_a16_SP(cpu.getNextWord())
	case 0x09:
		cpu.ADD_HL_n16(cpu.Registers.BC.Get())
	case 0x0A:
		LD_r8_n8(&cpu.Registers.A, cpu.Bus.ReadByte(cpu.Registers.BC.Get()))
	case 0x0B:
		DEC_r16(cpu.Registers.BC)
	case 0x0C:
		cpu.INC_r8(&cpu.Registers.C)
	case 0x0D:
		cpu.DEC_r8(&cpu.Registers.C)
	case 0x0E:
		LD_r8_n8(&cpu.Registers.C, cpu.getNextByte())
	case 0x0F:
		cpu.RRC_r8(&cpu.Registers.A)
	case 0x10:
		cpu.STOP()
	case 0x11:
		LD_r16_n16(cpu.Registers.DE, cpu.getNextWord())
	case 0x12:
		cpu.LD_a16_n8(cpu.Registers.DE.Get(), cpu.Registers.A)
	case 0x13:
		INC_r16(cpu.Registers.DE)
	case 0x14:
		cpu.INC_r8(&cpu.Registers.D)
	case 0x15:
		cpu.DEC_r8(&cpu.Registers.D)
	case 0x16:
		LD_r8_n8(&cpu.Registers.D, cpu.getNextByte())
	case 0x17:
		cpu.RL_r8(&cpu.Registers.A)
	case 0x18:
		cpu.JR_e8(int8(cpu.getNextByte()))
	case 0x19:
		cpu.ADD_HL_n16(cpu.Registers.DE.Get())
	case 0x1A:
		LD_r8_n8(&cpu.Registers.A, cpu.Bus.ReadByte(cpu.Registers.DE.Get()))
	case 0x1B:
		DEC_r16(cpu.Registers.DE)
	case 0x1C:
		cpu.INC_r8(&cpu.Registers.E)
	case 0x1D:
		cpu.DEC_r8(&cpu.Registers.E)
	case 0x1E:
		LD_r8_n8(&cpu.Registers.E, cpu.getNextByte())
	case 0x1F:
		cpu.RR_r8(&cpu.Registers.A)
	case 0x20:
		cpu.JR_NZ_e8(int8(cpu.getNextByte()))
	case 0x21:
		LD_r16_n16(cpu.Registers.HL, cpu.getNextWord())
	case 0x22:
		cpu.LD_HLI_A()
	case 0x23:
		INC_r16(cpu.Registers.HL)
	case 0x24:
		cpu.INC_r8(&cpu.Registers.H)
	case 0x25:
		cpu.DEC_r8(&cpu.Registers.H)
	case 0x26:
		LD_r8_n8(&cpu.Registers.H, cpu.getNextByte())
	case 0x27:
		cpu.DAA()
	case 0x28:
		cpu.JR_Z_e8(int8(cpu.getNextByte()))
	case 0x29:
		cpu.ADD_HL_n16(cpu.Registers.HL.Get())
	case 0x2A:
		cpu.LD_A_HLI()
	case 0x2B:
		DEC_r16(cpu.Registers.HL)
	case 0x2C:
		cpu.INC_r8(&cpu.Registers.L)
	case 0x2D:
		cpu.DEC_r8(&cpu.Registers.L)
	case 0x2E:
		LD_r8_n8(&cpu.Registers.L, cpu.getNextByte())
	case 0x2F:
		cpu.CPL()
	case 0x30:
		cpu.JR_NC_e8(int8(cpu.getNextByte()))
	case 0x31:
		cpu.LD_SP_n16(cpu.getNextWord())
	case 0x32:
		cpu.LD_HLD_A()
	case 0x33:
		cpu.INC_SP()
	case 0x34:
		cpu.INC_a16(cpu.Registers.HL.Get())
	case 0x35:
		cpu.DEC_a16(cpu.Registers.HL.Get())
	case 0x36:
		cpu.LD_a16_n8(cpu.Registers.HL.Get(), cpu.getNextByte())
	case 0x37:
		cpu.SCF()
	case 0x38:
		cpu.JR_C_e8(int8(cpu.getNextByte()))
	case 0x39:
		cpu.ADD_HL_n16(cpu.Registers.SP)
	case 0x3A:
		cpu.LD_A_HLD()
	case 0x3B:
		cpu.DEC_SP()
	case 0x3C:
		cpu.INC_r8(&cpu.Registers.A)
	case 0x3D:
		cpu.DEC_r8(&cpu.Registers.A)
	case 0x3E:
		LD_r8_n8(&cpu.Registers.A, cpu.getNextByte())
	case 0x3F:
		cpu.CCF()
	case 0x40:
		LD_r8_n8(&cpu.Registers.B, cpu.Registers.B)
	case 0x41:
		LD_r8_n8(&cpu.Registers.B, cpu.Registers.C)
	case 0x42:
		LD_r8_n8(&cpu.Registers.B, cpu.Registers.D)
	case 0x43:
		LD_r8_n8(&cpu.Registers.B, cpu.Registers.E)
	case 0x44:
		LD_r8_n8(&cpu.Registers.B, cpu.Registers.H)
	case 0x45:
		LD_r8_n8(&cpu.Registers.B, cpu.Registers.L)
	case 0x46:
		LD_r8_n8(&cpu.Registers.B, cpu.Bus.ReadByte(cpu.Registers.HL.Get()))
	case 0x47:
		LD_r8_n8(&cpu.Registers.B, cpu.Registers.A)
	case 0x48:
		LD_r8_n8(&cpu.Registers.C, cpu.Registers.B)
	case 0x49:
		LD_r8_n8(&cpu.Registers.C, cpu.Registers.C)
	case 0x4A:
		LD_r8_n8(&cpu.Registers.C, cpu.Registers.D)
	case 0x4B:
		LD_r8_n8(&cpu.Registers.C, cpu.Registers.E)
	case 0x4C:
		LD_r8_n8(&cpu.Registers.C, cpu.Registers.H)
	case 0x4D:
		LD_r8_n8(&cpu.Registers.C, cpu.Registers.L)
	case 0x4E:
		LD_r8_n8(&cpu.Registers.C, cpu.Bus.ReadByte(cpu.Registers.HL.Get()))
	case 0x4F:
		LD_r8_n8(&cpu.Registers.C, cpu.Registers.A)
	case 0x50:
		LD_r8_n8(&cpu.Registers.D, cpu.Registers.B)
	case 0x51:
		LD_r8_n8(&cpu.Registers.D, cpu.Registers.C)
	case 0x52:
		LD_r8_n8(&cpu.Registers.D, cpu.Registers.D)
	case 0x53:
		LD_r8_n8(&cpu.Registers.D, cpu.Registers.E)
	case 0x54:
		LD_r8_n8(&cpu.Registers.D, cpu.Registers.H)
	case 0x55:
		LD_r8_n8(&cpu.Registers.D, cpu.Registers.L)
	case 0x56:
		LD_r8_n8(&cpu.Registers.D, cpu.Bus.ReadByte(cpu.Registers.HL.Get()))
	case 0x57:
		LD_r8_n8(&cpu.Registers.D, cpu.Registers.A)
	case 0x58:
		LD_r8_n8(&cpu.Registers.E, cpu.Registers.B)
	case 0x59:
		LD_r8_n8(&cpu.Registers.E, cpu.Registers.C)
	case 0x5A:
		LD_r8_n8(&cpu.Registers.E, cpu.Registers.D)
	case 0x5B:
		LD_r8_n8(&cpu.Registers.E, cpu.Registers.E)
	case 0x5C:
		LD_r8_n8(&cpu.Registers.E, cpu.Registers.H)
	case 0x5D:
		LD_r8_n8(&cpu.Registers.E, cpu.Registers.L)
	case 0x5E:
		LD_r8_n8(&cpu.Registers.E, cpu.Bus.ReadByte(cpu.Registers.HL.Get()))
	case 0x5F:
		LD_r8_n8(&cpu.Registers.E, cpu.Registers.A)
	case 0x60:
		LD_r8_n8(&cpu.Registers.H, cpu.Registers.B)
	case 0x61:
		LD_r8_n8(&cpu.Registers.H, cpu.Registers.C)
	case 0x62:
		LD_r8_n8(&cpu.Registers.H, cpu.Registers.D)
	case 0x63:
		LD_r8_n8(&cpu.Registers.H, cpu.Registers.E)
	case 0x64:
		LD_r8_n8(&cpu.Registers.H, cpu.Registers.H)
	case 0x65:
		LD_r8_n8(&cpu.Registers.H, cpu.Registers.L)
	case 0x66:
		LD_r8_n8(&cpu.Registers.H, cpu.Bus.ReadByte(cpu.Registers.HL.Get()))
	case 0x67:
		LD_r8_n8(&cpu.Registers.H, cpu.Registers.A)
	case 0x68:
		LD_r8_n8(&cpu.Registers.L, cpu.Registers.B)
	case 0x69:
		LD_r8_n8(&cpu.Registers.L, cpu.Registers.C)
	case 0x6A:
		LD_r8_n8(&cpu.Registers.L, cpu.Registers.D)
	case 0x6B:
		LD_r8_n8(&cpu.Registers.L, cpu.Registers.E)
	case 0x6C:
		LD_r8_n8(&cpu.Registers.L, cpu.Registers.H)
	case 0x6D:
		LD_r8_n8(&cpu.Registers.L, cpu.Registers.L)
	case 0x6E:
		LD_r8_n8(&cpu.Registers.L, cpu.Bus.ReadByte(cpu.Registers.HL.Get()))
	case 0x6F:
		LD_r8_n8(&cpu.Registers.L, cpu.Registers.A)
	case 0x70:
		cpu.LD_a16_n8(cpu.Registers.HL.Get(), cpu.Registers.B)
	case 0x71:
		cpu.LD_a16_n8(cpu.Registers.HL.Get(), cpu.Registers.C)
	case 0x72:
		cpu.LD_a16_n8(cpu.Registers.HL.Get(), cpu.Registers.D)
	case 0x73:
		cpu.LD_a16_n8(cpu.Registers.HL.Get(), cpu.Registers.E)
	case 0x74:
		cpu.LD_a16_n8(cpu.Registers.HL.Get(), cpu.Registers.H)
	case 0x75:
		cpu.LD_a16_n8(cpu.Registers.HL.Get(), cpu.Registers.L)
	case 0x76:
		cpu.HALT()
	case 0x77:
		cpu.LD_a16_n8(cpu.Registers.HL.Get(), cpu.Registers.A)
	case 0x78:
		LD_r8_n8(&cpu.Registers.A, cpu.Registers.B)
	case 0x79:
		LD_r8_n8(&cpu.Registers.A, cpu.Registers.C)
	case 0x7A:
		LD_r8_n8(&cpu.Registers.A, cpu.Registers.D)
	case 0x7B:
		LD_r8_n8(&cpu.Registers.A, cpu.Registers.E)
	case 0x7C:
		LD_r8_n8(&cpu.Registers.A, cpu.Registers.H)
	case 0x7D:
		LD_r8_n8(&cpu.Registers.A, cpu.Registers.L)
	case 0x7E:
		LD_r8_n8(&cpu.Registers.A, cpu.Bus.ReadByte(cpu.Registers.HL.Get()))
	case 0x7F:
		LD_r8_n8(&cpu.Registers.A, cpu.Registers.A)
	case 0x80:
		cpu.ADD_A_n8(cpu.Registers.B)
	case 0x81:
		cpu.ADD_A_n8(cpu.Registers.C)
	case 0x82:
		cpu.ADD_A_n8(cpu.Registers.D)
	case 0x83:
		cpu.ADD_A_n8(cpu.Registers.E)
	case 0x84:
		cpu.ADD_A_n8(cpu.Registers.H)
	case 0x85:
		cpu.ADD_A_n8(cpu.Registers.L)
	case 0x86:
		cpu.ADD_A_n8(cpu.Bus.ReadByte((cpu.Registers.HL.Get())))
	case 0x87:
		cpu.ADD_A_n8(cpu.Registers.A)
	case 0x88:
		cpu.ADC_A_n8(cpu.Registers.B)
	case 0x89:
		cpu.ADC_A_n8(cpu.Registers.C)
	case 0x8A:
		cpu.ADC_A_n8(cpu.Registers.D)
	case 0x8B:
		cpu.ADC_A_n8(cpu.Registers.E)
	case 0x8C:
		cpu.ADC_A_n8(cpu.Registers.H)
	case 0x8D:
		cpu.ADC_A_n8(cpu.Registers.L)
	case 0x8E:
		cpu.ADC_A_n8(cpu.Bus.ReadByte((cpu.Registers.HL.Get())))
	case 0x8F:
		cpu.ADC_A_n8(cpu.Registers.A)
	case 0x90:
		cpu.SUB_A_n8(cpu.Registers.B)
	case 0x91:
		cpu.SUB_A_n8(cpu.Registers.C)
	case 0x92:
		cpu.SUB_A_n8(cpu.Registers.D)
	case 0x93:
		cpu.SUB_A_n8(cpu.Registers.E)
	case 0x94:
		cpu.SUB_A_n8(cpu.Registers.H)
	case 0x95:
		cpu.SUB_A_n8(cpu.Registers.L)
	case 0x96:
		cpu.SUB_A_n8(cpu.Bus.ReadByte((cpu.Registers.HL.Get())))
	case 0x97:
		cpu.SUB_A_n8(cpu.Registers.A)
	case 0x98:
		cpu.SBC_A_n8(cpu.Registers.B)
	case 0x99:
		cpu.SBC_A_n8(cpu.Registers.C)
	case 0x9A:
		cpu.SBC_A_n8(cpu.Registers.D)
	case 0x9B:
		cpu.SBC_A_n8(cpu.Registers.E)
	case 0x9C:
		cpu.SBC_A_n8(cpu.Registers.H)
	case 0x9D:
		cpu.SBC_A_n8(cpu.Registers.L)
	case 0x9E:
		cpu.SBC_A_n8(cpu.Bus.ReadByte((cpu.Registers.HL.Get())))
	case 0x9F:
		cpu.SBC_A_n8(cpu.Registers.A)
	case 0xA0:
		cpu.AND_A_n8(cpu.Registers.B)
	case 0xA1:
		cpu.AND_A_n8(cpu.Registers.C)
	case 0xA2:
		cpu.AND_A_n8(cpu.Registers.D)
	case 0xA3:
		cpu.AND_A_n8(cpu.Registers.E)
	case 0xA4:
		cpu.AND_A_n8(cpu.Registers.H)
	case 0xA5:
		cpu.AND_A_n8(cpu.Registers.L)
	case 0xA6:
		cpu.AND_A_n8(cpu.Bus.ReadByte((cpu.Registers.HL.Get())))
	case 0xA7:
		cpu.AND_A_n8(cpu.Registers.A)
	case 0xA8:
		cpu.XOR_A_n8(cpu.Registers.B)
	case 0xA9:
		cpu.XOR_A_n8(cpu.Registers.C)
	case 0xAA:
		cpu.XOR_A_n8(cpu.Registers.D)
	case 0xAB:
		cpu.XOR_A_n8(cpu.Registers.E)
	case 0xAC:
		cpu.XOR_A_n8(cpu.Registers.H)
	case 0xAD:
		cpu.XOR_A_n8(cpu.Registers.L)
	case 0xAE:
		cpu.XOR_A_n8(cpu.Bus.ReadByte((cpu.Registers.HL.Get())))
	case 0xAF:
		cpu.XOR_A_n8(cpu.Registers.A)
	case 0xB0:
		cpu.AND_A_n8(cpu.Registers.B)
	case 0xB1:
		cpu.AND_A_n8(cpu.Registers.C)
	case 0xB2:
		cpu.AND_A_n8(cpu.Registers.D)
	case 0xB3:
		cpu.AND_A_n8(cpu.Registers.E)
	case 0xB4:
		cpu.AND_A_n8(cpu.Registers.H)
	case 0xB5:
		cpu.AND_A_n8(cpu.Registers.L)
	case 0xB6:
		cpu.AND_A_n8(cpu.Bus.ReadByte((cpu.Registers.HL.Get())))
	case 0xB7:
		cpu.AND_A_n8(cpu.Registers.A)
	case 0xB8:
		cpu.CP_A_n8(cpu.Registers.B)
	case 0xB9:
		cpu.CP_A_n8(cpu.Registers.C)
	case 0xBA:
		cpu.CP_A_n8(cpu.Registers.D)
	case 0xBB:
		cpu.CP_A_n8(cpu.Registers.E)
	case 0xBC:
		cpu.CP_A_n8(cpu.Registers.H)
	case 0xBD:
		cpu.CP_A_n8(cpu.Registers.L)
	case 0xBE:
		cpu.CP_A_n8(cpu.Bus.ReadByte((cpu.Registers.HL.Get())))
	case 0xBF:
		cpu.CP_A_n8(cpu.Registers.A)
	case 0xC0:
		cpu.RET_NZ()
	case 0xC1:
		cpu.POP_r16(cpu.Registers.BC)
	case 0xC2:
		cpu.JP_NZ_n16(cpu.getNextWord())
	case 0xC3:
		cpu.JP_n16(cpu.getNextWord())
	case 0xC4:
		cpu.CALL_NZ_n16(cpu.getNextWord())
	case 0xC5:
		cpu.PUSH_r16(cpu.Registers.BC)
	case 0xC6:
		cpu.ADD_A_n8(cpu.getNextByte())
	case 0xC7:
		cpu.RST(0x00)
	case 0xC8:
		cpu.RET_Z()
	case 0xC9:
		cpu.RET()
	case 0xCA:
		cpu.JP_Z_n16(cpu.getNextWord())
	case 0xCB:
		cpu.prefixed(cpu.getNextByte())
	case 0xCC:
		cpu.CALL_Z_n16(cpu.getNextWord())
	case 0xCD:
		cpu.CALL_n16(cpu.getNextWord())
	case 0xCE:
		cpu.ADC_A_n8(cpu.getNextByte())
	case 0xCF:
		cpu.RST(0x08)
	case 0xD0:
		cpu.RET_NC()
	case 0xD1:
		cpu.POP_r16(cpu.Registers.DE)
	case 0xD2:
		cpu.JP_NC_n16(cpu.getNextWord())
	case 0xD3:
		// no operation
	case 0xD4:
		cpu.CALL_NC_n16(cpu.getNextWord())
	case 0xD5:
		cpu.PUSH_r16(cpu.Registers.DE)
	case 0xD6:
		cpu.SUB_A_n8(cpu.getNextByte())
	case 0xD7:
		cpu.RST(0x10)
	case 0xD8:
		cpu.RET_C()
	case 0xD9:
		cpu.RETI()
	case 0xDA:
		cpu.JP_C_n16(cpu.getNextWord())
	case 0xDB:
		// no operation
	case 0xDC:
		cpu.CALL_C_n16(cpu.getNextWord())
	case 0xDD:
		// no operation
	case 0xDE:
		cpu.SBC_A_n8(cpu.getNextByte())
	case 0xDF:
		cpu.RST(0x18)
	case 0xE0:
		cpu.LDH_a8_r8(cpu.getNextByte(), cpu.Registers.A)
	case 0xE1:
		cpu.POP_r16(cpu.Registers.HL)
	case 0xE2:
		cpu.LDH_a8_r8(cpu.Registers.C, cpu.Registers.A)
	case 0xE3:
		// no operation
	case 0xE4:
		// no operation
	case 0xE5:
		cpu.PUSH_r16(cpu.Registers.HL)
	case 0xE6:
		cpu.AND_A_n8(cpu.getNextByte())
	case 0xE7:
		cpu.RST(0x20)
	case 0xE8:
		cpu.ADD_SP_e8(int8(cpu.getNextByte()))
	case 0xE9:
		cpu.JP_n16(cpu.Registers.HL.Get())
	case 0xEA:
		cpu.LD_a16_n8(cpu.getNextWord(), cpu.Registers.A)
	case 0xEB:
		// no operation
	case 0xEC:
		// no operation
	case 0xED:
		// no operation
	case 0xEE:
		cpu.XOR_A_n8(cpu.getNextByte())
	case 0xEF:
		cpu.RST(0x18)
	case 0xF0:
		cpu.LDH_r8_a16(&cpu.Registers.A, cpu.getNextWord())
	case 0xF1:
		cpu.POP_r16(cpu.Registers.AF)
	case 0xF2:
		cpu.LDH_r8_a8(&cpu.Registers.A, cpu.Registers.C)
	case 0xF3:
		cpu.DI()
	case 0xF4:
		// no operation
	case 0xF5:
		cpu.PUSH_r16(cpu.Registers.AF)
	case 0xF6:
		cpu.OR_A_n8(cpu.getNextByte())
	case 0xF7:
		cpu.RST(0x30)
	case 0xF8:
		cpu.LD_HL_SP_e8(int8(cpu.getNextByte()))
	case 0xF9:
		cpu.LD_SP_n16(cpu.Registers.HL.Get())
	case 0xFA:
		cpu.LDH_r8_a16(&cpu.Registers.A, cpu.getNextWord())
	case 0xFB:
		cpu.EI()
	case 0xFC:
		// no operation
	case 0xFD:
		// no operation
	case 0xFE:
		cpu.CP_A_n8(cpu.getNextByte())
	case 0xFF:
		cpu.RST(0x38)
	}
}

func (cpu *CPU) prefixed(nextByte byte) {
	switch nextByte {
	case 0x00:
		cpu.RLC_r8(&cpu.Registers.B)
	case 0x01:
		cpu.RLC_r8(&cpu.Registers.C)
	case 0x02:
		cpu.RLC_r8(&cpu.Registers.D)
	case 0x03:
		cpu.RLC_r8(&cpu.Registers.E)
	case 0x04:
		cpu.RLC_r8(&cpu.Registers.H)
	case 0x05:
		cpu.RLC_r8(&cpu.Registers.L)
	case 0x06:
		cpu.RLC_a16(cpu.Registers.HL.Get())
	case 0x07:
		cpu.RLC_r8(&cpu.Registers.A)
	case 0x08:
		cpu.RRC_r8(&cpu.Registers.B)
	case 0x09:
		cpu.RRC_r8(&cpu.Registers.C)
	case 0x0A:
		cpu.RRC_r8(&cpu.Registers.D)
	case 0x0B:
		cpu.RRC_r8(&cpu.Registers.E)
	case 0x0C:
		cpu.RRC_r8(&cpu.Registers.H)
	case 0x0D:
		cpu.RRC_r8(&cpu.Registers.L)
	case 0x0E:
		cpu.RRC_a16(cpu.Registers.HL.Get())
	case 0x0F:
		cpu.RRC_r8(&cpu.Registers.A)
	case 0x10:
		cpu.RL_r8(&cpu.Registers.B)
	case 0x11:
		cpu.RL_r8(&cpu.Registers.C)
	case 0x12:
		cpu.RL_r8(&cpu.Registers.D)
	case 0x13:
		cpu.RL_r8(&cpu.Registers.E)
	case 0x14:
		cpu.RL_r8(&cpu.Registers.H)
	case 0x15:
		cpu.RL_r8(&cpu.Registers.L)
	case 0x16:
		cpu.RL_a16(cpu.Registers.HL.Get())
	case 0x17:
		cpu.RL_r8(&cpu.Registers.A)
	case 0x18:
		cpu.RR_r8(&cpu.Registers.B)
	case 0x19:
		cpu.RR_r8(&cpu.Registers.C)
	case 0x1A:
		cpu.RR_r8(&cpu.Registers.D)
	case 0x1B:
		cpu.RR_r8(&cpu.Registers.E)
	case 0x1C:
		cpu.RR_r8(&cpu.Registers.H)
	case 0x1D:
		cpu.RR_r8(&cpu.Registers.L)
	case 0x1E:
		cpu.RR_a16(cpu.Registers.HL.Get())
	case 0x1F:
		cpu.RR_r8(&cpu.Registers.A)
	case 0x20:
		cpu.SLA_r8(&cpu.Registers.B)
	case 0x21:
		cpu.SLA_r8(&cpu.Registers.C)
	case 0x22:
		cpu.SLA_r8(&cpu.Registers.D)
	case 0x23:
		cpu.SLA_r8(&cpu.Registers.E)
	case 0x24:
		cpu.SLA_r8(&cpu.Registers.H)
	case 0x25:
		cpu.SLA_r8(&cpu.Registers.L)
	case 0x26:
		cpu.SLA_a16(cpu.Registers.HL.Get())
	case 0x27:
		cpu.SLA_r8(&cpu.Registers.A)
	case 0x28:
		cpu.SRA_r8(&cpu.Registers.B)
	case 0x29:
		cpu.SRA_r8(&cpu.Registers.C)
	case 0x2A:
		cpu.SRA_r8(&cpu.Registers.D)
	case 0x2B:
		cpu.SRA_r8(&cpu.Registers.E)
	case 0x2C:
		cpu.SRA_r8(&cpu.Registers.H)
	case 0x2D:
		cpu.SRA_r8(&cpu.Registers.L)
	case 0x2E:
		cpu.SRA_a16(cpu.Registers.HL.Get())
	case 0x2F:
		cpu.SRA_r8(&cpu.Registers.A)
	case 0x30:
		cpu.SWAP_r8(&cpu.Registers.B)
	case 0x31:
		cpu.SWAP_r8(&cpu.Registers.C)
	case 0x32:
		cpu.SWAP_r8(&cpu.Registers.D)
	case 0x33:
		cpu.SWAP_r8(&cpu.Registers.E)
	case 0x34:
		cpu.SWAP_r8(&cpu.Registers.H)
	case 0x35:
		cpu.SWAP_r8(&cpu.Registers.L)
	case 0x36:
		cpu.SWAP_a16(cpu.Registers.HL.Get())
	case 0x37:
		cpu.SWAP_r8(&cpu.Registers.A)
	case 0x38:
		cpu.SRL_r8(&cpu.Registers.B)
	case 0x39:
		cpu.SRL_r8(&cpu.Registers.C)
	case 0x3A:
		cpu.SRL_r8(&cpu.Registers.D)
	case 0x3B:
		cpu.SRL_r8(&cpu.Registers.E)
	case 0x3C:
		cpu.SRL_r8(&cpu.Registers.H)
	case 0x3D:
		cpu.SRL_r8(&cpu.Registers.L)
	case 0x3E:
		cpu.SRL_a16(cpu.Registers.HL.Get())
	case 0x3F:
		cpu.SRL_r8(&cpu.Registers.A)
	case 0x40:
		cpu.BIT_n8_n8(byte(0), cpu.Registers.B)
	case 0x41:
		cpu.BIT_n8_n8(byte(0), cpu.Registers.B)
	case 0x42:
		cpu.BIT_n8_n8(byte(0), cpu.Registers.B)
	case 0x43:
		cpu.BIT_n8_n8(byte(0), cpu.Registers.B)
	case 0x44:
		cpu.BIT_n8_n8(byte(0), cpu.Registers.B)
	case 0x45:
		cpu.BIT_n8_n8(byte(0), cpu.Registers.B)
	case 0x46:
		cpu.BIT_n8_n8(byte(0), cpu.Bus.ReadByte(cpu.Registers.HL.Get()))
	case 0x47:
		cpu.BIT_n8_n8(byte(0), cpu.Registers.A)
	case 0x48:
		cpu.BIT_n8_n8(byte(1), cpu.Registers.B)
	case 0x49:
		cpu.BIT_n8_n8(byte(1), cpu.Registers.B)
	case 0x4A:
		cpu.BIT_n8_n8(byte(1), cpu.Registers.B)
	case 0x4B:
		cpu.BIT_n8_n8(byte(1), cpu.Registers.B)
	case 0x4C:
		cpu.BIT_n8_n8(byte(1), cpu.Registers.B)
	case 0x4D:
		cpu.BIT_n8_n8(byte(1), cpu.Registers.B)
	case 0x4E:
		cpu.BIT_n8_n8(byte(1), cpu.Bus.ReadByte(cpu.Registers.HL.Get()))
	case 0x4F:
		cpu.BIT_n8_n8(byte(1), cpu.Registers.A)
	case 0x50:
		cpu.BIT_n8_n8(byte(2), cpu.Registers.B)
	case 0x51:
		cpu.BIT_n8_n8(byte(2), cpu.Registers.B)
	case 0x52:
		cpu.BIT_n8_n8(byte(2), cpu.Registers.B)
	case 0x53:
		cpu.BIT_n8_n8(byte(2), cpu.Registers.B)
	case 0x54:
		cpu.BIT_n8_n8(byte(2), cpu.Registers.B)
	case 0x55:
		cpu.BIT_n8_n8(byte(2), cpu.Registers.B)
	case 0x56:
		cpu.BIT_n8_n8(byte(2), cpu.Bus.ReadByte(cpu.Registers.HL.Get()))
	case 0x57:
		cpu.BIT_n8_n8(byte(2), cpu.Registers.A)
	case 0x58:
		cpu.BIT_n8_n8(byte(3), cpu.Registers.B)
	case 0x59:
		cpu.BIT_n8_n8(byte(3), cpu.Registers.B)
	case 0x5A:
		cpu.BIT_n8_n8(byte(3), cpu.Registers.B)
	case 0x5B:
		cpu.BIT_n8_n8(byte(3), cpu.Registers.B)
	case 0x5C:
		cpu.BIT_n8_n8(byte(3), cpu.Registers.B)
	case 0x5D:
		cpu.BIT_n8_n8(byte(3), cpu.Registers.B)
	case 0x5E:
		cpu.BIT_n8_n8(byte(3), cpu.Bus.ReadByte(cpu.Registers.HL.Get()))
	case 0x5F:
		cpu.BIT_n8_n8(byte(3), cpu.Registers.A)
	case 0x60:
		cpu.BIT_n8_n8(byte(4), cpu.Registers.B)
	case 0x61:
		cpu.BIT_n8_n8(byte(4), cpu.Registers.B)
	case 0x62:
		cpu.BIT_n8_n8(byte(4), cpu.Registers.B)
	case 0x63:
		cpu.BIT_n8_n8(byte(4), cpu.Registers.B)
	case 0x64:
		cpu.BIT_n8_n8(byte(4), cpu.Registers.B)
	case 0x65:
		cpu.BIT_n8_n8(byte(4), cpu.Registers.B)
	case 0x66:
		cpu.BIT_n8_n8(byte(4), cpu.Bus.ReadByte(cpu.Registers.HL.Get()))
	case 0x67:
		cpu.BIT_n8_n8(byte(4), cpu.Registers.A)
	case 0x68:
		cpu.BIT_n8_n8(byte(5), cpu.Registers.B)
	case 0x69:
		cpu.BIT_n8_n8(byte(5), cpu.Registers.B)
	case 0x6A:
		cpu.BIT_n8_n8(byte(5), cpu.Registers.B)
	case 0x6B:
		cpu.BIT_n8_n8(byte(5), cpu.Registers.B)
	case 0x6C:
		cpu.BIT_n8_n8(byte(5), cpu.Registers.B)
	case 0x6D:
		cpu.BIT_n8_n8(byte(5), cpu.Registers.B)
	case 0x6E:
		cpu.BIT_n8_n8(byte(5), cpu.Bus.ReadByte(cpu.Registers.HL.Get()))
	case 0x6F:
		cpu.BIT_n8_n8(byte(5), cpu.Registers.A)
	case 0x70:
		cpu.BIT_n8_n8(byte(6), cpu.Registers.B)
	case 0x71:
		cpu.BIT_n8_n8(byte(6), cpu.Registers.B)
	case 0x72:
		cpu.BIT_n8_n8(byte(6), cpu.Registers.B)
	case 0x73:
		cpu.BIT_n8_n8(byte(6), cpu.Registers.B)
	case 0x74:
		cpu.BIT_n8_n8(byte(6), cpu.Registers.B)
	case 0x75:
		cpu.BIT_n8_n8(byte(6), cpu.Registers.B)
	case 0x76:
		cpu.BIT_n8_n8(byte(6), cpu.Bus.ReadByte(cpu.Registers.HL.Get()))
	case 0x77:
		cpu.BIT_n8_n8(byte(6), cpu.Registers.A)
	case 0x78:
		cpu.BIT_n8_n8(byte(7), cpu.Registers.B)
	case 0x79:
		cpu.BIT_n8_n8(byte(7), cpu.Registers.B)
	case 0x7A:
		cpu.BIT_n8_n8(byte(7), cpu.Registers.B)
	case 0x7B:
		cpu.BIT_n8_n8(byte(7), cpu.Registers.B)
	case 0x7C:
		cpu.BIT_n8_n8(byte(7), cpu.Registers.B)
	case 0x7D:
		cpu.BIT_n8_n8(byte(7), cpu.Registers.B)
	case 0x7E:
		cpu.BIT_n8_n8(byte(7), cpu.Bus.ReadByte(cpu.Registers.HL.Get()))
	case 0x7F:
		cpu.BIT_n8_n8(byte(7), cpu.Registers.A)
	case 0x80:
		RES_n8_r8(byte(0), &cpu.Registers.B)
	case 0x81:
		RES_n8_r8(byte(0), &cpu.Registers.B)
	case 0x82:
		RES_n8_r8(byte(0), &cpu.Registers.B)
	case 0x83:
		RES_n8_r8(byte(0), &cpu.Registers.B)
	case 0x84:
		RES_n8_r8(byte(0), &cpu.Registers.B)
	case 0x85:
		RES_n8_r8(byte(0), &cpu.Registers.B)
	case 0x86:
		cpu.RES_n8_a16(byte(0), cpu.Registers.HL.Get())
	case 0x87:
		RES_n8_r8(byte(0), &cpu.Registers.A)
	case 0x88:
		RES_n8_r8(byte(1), &cpu.Registers.B)
	case 0x89:
		RES_n8_r8(byte(1), &cpu.Registers.B)
	case 0x8A:
		RES_n8_r8(byte(1), &cpu.Registers.B)
	case 0x8B:
		RES_n8_r8(byte(1), &cpu.Registers.B)
	case 0x8C:
		RES_n8_r8(byte(1), &cpu.Registers.B)
	case 0x8D:
		RES_n8_r8(byte(1), &cpu.Registers.B)
	case 0x8E:
		cpu.RES_n8_a16(byte(1), cpu.Registers.HL.Get())
	case 0x8F:
		RES_n8_r8(byte(1), &cpu.Registers.A)
	case 0x90:
		RES_n8_r8(byte(2), &cpu.Registers.B)
	case 0x91:
		RES_n8_r8(byte(2), &cpu.Registers.B)
	case 0x92:
		RES_n8_r8(byte(2), &cpu.Registers.B)
	case 0x93:
		RES_n8_r8(byte(2), &cpu.Registers.B)
	case 0x94:
		RES_n8_r8(byte(2), &cpu.Registers.B)
	case 0x95:
		RES_n8_r8(byte(2), &cpu.Registers.B)
	case 0x96:
		cpu.RES_n8_a16(byte(2), cpu.Registers.HL.Get())
	case 0x97:
		RES_n8_r8(byte(2), &cpu.Registers.A)
	case 0x98:
		RES_n8_r8(byte(3), &cpu.Registers.B)
	case 0x99:
		RES_n8_r8(byte(3), &cpu.Registers.B)
	case 0x9A:
		RES_n8_r8(byte(3), &cpu.Registers.B)
	case 0x9B:
		RES_n8_r8(byte(3), &cpu.Registers.B)
	case 0x9C:
		RES_n8_r8(byte(3), &cpu.Registers.B)
	case 0x9D:
		RES_n8_r8(byte(3), &cpu.Registers.B)
	case 0x9E:
		cpu.RES_n8_a16(byte(3), cpu.Registers.HL.Get())
	case 0x9F:
		RES_n8_r8(byte(3), &cpu.Registers.A)
	case 0xA0:
		RES_n8_r8(byte(4), &cpu.Registers.B)
	case 0xA1:
		RES_n8_r8(byte(4), &cpu.Registers.B)
	case 0xA2:
		RES_n8_r8(byte(4), &cpu.Registers.B)
	case 0xA3:
		RES_n8_r8(byte(4), &cpu.Registers.B)
	case 0xA4:
		RES_n8_r8(byte(4), &cpu.Registers.B)
	case 0xA5:
		RES_n8_r8(byte(4), &cpu.Registers.B)
	case 0xA6:
		cpu.RES_n8_a16(byte(4), cpu.Registers.HL.Get())
	case 0xA7:
		RES_n8_r8(byte(4), &cpu.Registers.A)
	case 0xA8:
		RES_n8_r8(byte(5), &cpu.Registers.B)
	case 0xA9:
		RES_n8_r8(byte(5), &cpu.Registers.B)
	case 0xAA:
		RES_n8_r8(byte(5), &cpu.Registers.B)
	case 0xAB:
		RES_n8_r8(byte(5), &cpu.Registers.B)
	case 0xAC:
		RES_n8_r8(byte(5), &cpu.Registers.B)
	case 0xAD:
		RES_n8_r8(byte(5), &cpu.Registers.B)
	case 0xAE:
		cpu.RES_n8_a16(byte(5), cpu.Registers.HL.Get())
	case 0xAF:
		RES_n8_r8(byte(5), &cpu.Registers.A)
	case 0xB0:
		RES_n8_r8(byte(6), &cpu.Registers.B)
	case 0xB1:
		RES_n8_r8(byte(6), &cpu.Registers.B)
	case 0xB2:
		RES_n8_r8(byte(6), &cpu.Registers.B)
	case 0xB3:
		RES_n8_r8(byte(6), &cpu.Registers.B)
	case 0xB4:
		RES_n8_r8(byte(6), &cpu.Registers.B)
	case 0xB5:
		RES_n8_r8(byte(6), &cpu.Registers.B)
	case 0xB6:
		cpu.RES_n8_a16(byte(6), cpu.Registers.HL.Get())
	case 0xB7:
		RES_n8_r8(byte(6), &cpu.Registers.A)
	case 0xB8:
		RES_n8_r8(byte(7), &cpu.Registers.B)
	case 0xB9:
		RES_n8_r8(byte(7), &cpu.Registers.B)
	case 0xBA:
		RES_n8_r8(byte(7), &cpu.Registers.B)
	case 0xBB:
		RES_n8_r8(byte(7), &cpu.Registers.B)
	case 0xBC:
		RES_n8_r8(byte(7), &cpu.Registers.B)
	case 0xBD:
		RES_n8_r8(byte(7), &cpu.Registers.B)
	case 0xBE:
		cpu.RES_n8_a16(byte(7), cpu.Registers.HL.Get())
	case 0xBF:
		RES_n8_r8(byte(7), &cpu.Registers.A)
	case 0xC0:
		SET_n8_r8(byte(0), &cpu.Registers.B)
	case 0xC1:
		SET_n8_r8(byte(0), &cpu.Registers.B)
	case 0xC2:
		SET_n8_r8(byte(0), &cpu.Registers.B)
	case 0xC3:
		SET_n8_r8(byte(0), &cpu.Registers.B)
	case 0xC4:
		SET_n8_r8(byte(0), &cpu.Registers.B)
	case 0xC5:
		SET_n8_r8(byte(0), &cpu.Registers.B)
	case 0xC6:
		cpu.SET_n8_a16(byte(0), cpu.Registers.HL.Get())
	case 0xC7:
		SET_n8_r8(byte(0), &cpu.Registers.A)
	case 0xC8:
		SET_n8_r8(byte(1), &cpu.Registers.B)
	case 0xC9:
		SET_n8_r8(byte(1), &cpu.Registers.B)
	case 0xCA:
		SET_n8_r8(byte(1), &cpu.Registers.B)
	case 0xCB:
		SET_n8_r8(byte(1), &cpu.Registers.B)
	case 0xCC:
		SET_n8_r8(byte(1), &cpu.Registers.B)
	case 0xCD:
		SET_n8_r8(byte(1), &cpu.Registers.B)
	case 0xCE:
		cpu.SET_n8_a16(byte(1), cpu.Registers.HL.Get())
	case 0xCF:
		SET_n8_r8(byte(1), &cpu.Registers.A)
	case 0xD0:
		SET_n8_r8(byte(2), &cpu.Registers.B)
	case 0xD1:
		SET_n8_r8(byte(2), &cpu.Registers.B)
	case 0xD2:
		SET_n8_r8(byte(2), &cpu.Registers.B)
	case 0xD3:
		SET_n8_r8(byte(2), &cpu.Registers.B)
	case 0xD4:
		SET_n8_r8(byte(2), &cpu.Registers.B)
	case 0xD5:
		SET_n8_r8(byte(2), &cpu.Registers.B)
	case 0xD6:
		cpu.SET_n8_a16(byte(2), cpu.Registers.HL.Get())
	case 0xD7:
		SET_n8_r8(byte(2), &cpu.Registers.A)
	case 0xD8:
		SET_n8_r8(byte(3), &cpu.Registers.B)
	case 0xD9:
		SET_n8_r8(byte(3), &cpu.Registers.B)
	case 0xDA:
		SET_n8_r8(byte(3), &cpu.Registers.B)
	case 0xDB:
		SET_n8_r8(byte(3), &cpu.Registers.B)
	case 0xDC:
		SET_n8_r8(byte(3), &cpu.Registers.B)
	case 0xDD:
		SET_n8_r8(byte(3), &cpu.Registers.B)
	case 0xDE:
		cpu.SET_n8_a16(byte(3), cpu.Registers.HL.Get())
	case 0xDF:
		SET_n8_r8(byte(3), &cpu.Registers.A)
	case 0xE0:
		SET_n8_r8(byte(4), &cpu.Registers.B)
	case 0xE1:
		SET_n8_r8(byte(4), &cpu.Registers.B)
	case 0xE2:
		SET_n8_r8(byte(4), &cpu.Registers.B)
	case 0xE3:
		SET_n8_r8(byte(4), &cpu.Registers.B)
	case 0xE4:
		SET_n8_r8(byte(4), &cpu.Registers.B)
	case 0xE5:
		SET_n8_r8(byte(4), &cpu.Registers.B)
	case 0xE6:
		cpu.SET_n8_a16(byte(4), cpu.Registers.HL.Get())
	case 0xE7:
		SET_n8_r8(byte(4), &cpu.Registers.A)
	case 0xE8:
		SET_n8_r8(byte(5), &cpu.Registers.B)
	case 0xE9:
		SET_n8_r8(byte(5), &cpu.Registers.B)
	case 0xEA:
		SET_n8_r8(byte(5), &cpu.Registers.B)
	case 0xEB:
		SET_n8_r8(byte(5), &cpu.Registers.B)
	case 0xEC:
		SET_n8_r8(byte(5), &cpu.Registers.B)
	case 0xED:
		SET_n8_r8(byte(5), &cpu.Registers.B)
	case 0xEE:
		cpu.SET_n8_a16(byte(5), cpu.Registers.HL.Get())
	case 0xEF:
		SET_n8_r8(byte(5), &cpu.Registers.A)
	case 0xF0:
		SET_n8_r8(byte(6), &cpu.Registers.B)
	case 0xF1:
		SET_n8_r8(byte(6), &cpu.Registers.B)
	case 0xF2:
		SET_n8_r8(byte(6), &cpu.Registers.B)
	case 0xF3:
		SET_n8_r8(byte(6), &cpu.Registers.B)
	case 0xF4:
		SET_n8_r8(byte(6), &cpu.Registers.B)
	case 0xF5:
		SET_n8_r8(byte(6), &cpu.Registers.B)
	case 0xF6:
		cpu.SET_n8_a16(byte(6), cpu.Registers.HL.Get())
	case 0xF7:
		SET_n8_r8(byte(6), &cpu.Registers.A)
	case 0xF8:
		SET_n8_r8(byte(7), &cpu.Registers.B)
	case 0xF9:
		SET_n8_r8(byte(7), &cpu.Registers.B)
	case 0xFA:
		SET_n8_r8(byte(7), &cpu.Registers.B)
	case 0xFB:
		SET_n8_r8(byte(7), &cpu.Registers.B)
	case 0xFC:
		SET_n8_r8(byte(7), &cpu.Registers.B)
	case 0xFD:
		SET_n8_r8(byte(7), &cpu.Registers.B)
	case 0xFE:
		cpu.SET_n8_a16(byte(7), cpu.Registers.HL.Get())
	case 0xFF:
		SET_n8_r8(byte(7), &cpu.Registers.A)
	}
}
