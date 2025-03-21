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

func GetPair(high *Register, low *Register) uint16 {
	return uint16(high.Get())<<8 + uint16(low.Get())
}

func SetPair(high *Register, low *Register, val uint16) {
	high.Set(byte(val >> 8))
	low.Set(byte(val))
}

func IncrementPair(high *Register, low *Register) {
	if low.Get() == 255 {
		high.Increment()
		low.Increment()
	} else {
		low.Increment()
	}
}

func DecrementPair(high *Register, low *Register) {
	if low.Get() == 0 {
		high.Decrement()
		low.Decrement()
	} else {
		low.Decrement()
	}
}
