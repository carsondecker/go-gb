package cpu

type Registers struct {
	A *Register
	F *Register
	B *Register
	C *Register
	D *Register
	E *Register
	H *Register
	L *Register

	AF *RegisterPair
	BC *RegisterPair
	DE *RegisterPair
	HL *RegisterPair

	SP uint16
	PC uint16
}

func NewRegisters() *Registers {
	A := NewRegister()
	F := NewRegister()
	B := NewRegister()
	C := NewRegister()
	D := NewRegister()
	E := NewRegister()
	H := NewRegister()
	L := NewRegister()

	return &Registers{
		A: A,
		F: F,
		B: B,
		C: C,
		D: D,
		E: E,
		H: H,
		L: L,

		AF: NewRegisterPair(A, F),
		BC: NewRegisterPair(B, C),
		DE: NewRegisterPair(D, E),
		HL: NewRegisterPair(H, L),

		SP: 0,
		PC: 0,
	}
}

func (r *Registers) GetZeroFlag() bool {
	return r.F.CheckBit(7)
}

func (r *Registers) SetZeroFlag(val bool) {
	r.F.SetBit(7, val)
}

func (r *Registers) GetSubtractionFlag() bool {
	return r.F.CheckBit(6)
}

func (r *Registers) SetSubtractionFlag(val bool) {
	r.F.SetBit(6, val)
}

func (r *Registers) GetHalfCarryFlag() bool {
	return r.F.CheckBit(5)
}

func (r *Registers) SetHalfCarryFlag(val bool) {
	r.F.SetBit(5, val)
}

func (r *Registers) GetCarryFlag(val bool) bool {
	return r.F.CheckBit(4)
}

func (r *Registers) SetCarryFlag(val bool) {
	r.F.SetBit(5, val)
}
