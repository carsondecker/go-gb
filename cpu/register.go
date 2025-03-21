package cpu

import "fmt"

type Register struct {
	val byte
}

func NewRegister() *Register {
	return &Register{
		val: 0,
	}
}

func (r *Register) Set(newVal byte) {
	r.val = newVal
}

func (r *Register) Get() byte {
	return r.val
}

func (r *Register) Reset() {
	r.val = 0
}

func (r *Register) CheckBit(bit byte) bool {
	return r.val&(1<<bit) != 0
}

func (r *Register) SetBit(bit byte, val bool) {
	if val {
		r.val = r.val | (1 << bit)
	} else {
		r.val = r.val &^ (1 << bit)
	}
}

func (r *Register) Increment() {
	r.val++
}

func (r *Register) Decrement() {
	r.val--
}

func (r *Register) String() string {
	return fmt.Sprintf("%08b", r.Get())
}

type RegisterPair struct {
	high *Register
	low  *Register
}

func NewRegisterPair(high *Register, low *Register) *RegisterPair {
	return &RegisterPair{
		high,
		low,
	}
}

func (r *RegisterPair) Get() uint16 {
	return uint16(r.high.Get())<<8 + uint16(r.low.Get())
}

func (r *RegisterPair) Set(val uint16) {
	r.high.Set(byte(val >> 8))
	r.low.Set(byte(val))
}

func (r *RegisterPair) Increment() {
	r.Set(r.Get() + 1)
}

func (r *RegisterPair) Decrement() {
	r.Set(r.Get() - 1)
}
