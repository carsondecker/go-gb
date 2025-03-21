package cpu

import (
	"testing"
)

func TestRegisterSetGetReset(t *testing.T) {
	reg := NewRegister()
	if reg.Get() != 0 {
		t.Errorf("value %d in register after register creation", reg.val)
	}

	reg.Set(123)
	if reg.Get() != 123 {
		t.Errorf("value %d in register after setting register to 123", reg.Get())
	}

	reg.Reset()
	if reg.Get() != 0 {
		t.Errorf("value %d in register reset", reg.Get())
	}
}

func TestRegisterSetCheckBit(t *testing.T) {
	reg := NewRegister()
	reg.Set(32)
	if !reg.CheckBit(5) {
		t.Errorf("bit 5 in 32 incorrectly read as 0")
	}
	if reg.CheckBit(4) {
		t.Errorf("bit 4 in 32 incorrectly read as 1")
	}

	reg.SetBit(4, true)
	if !reg.CheckBit(4) {
		t.Errorf("bit 4 incorrectly read as 0 after being set to 1")
	}
}

func TestRegisterSetGetPair(t *testing.T) {
	high := NewRegister()
	low := NewRegister()
	r := NewRegisterPair(high, low)
	if r.Get() != 0 {
		t.Errorf("value %d in pair after register creation", r.Get())
	}

	r.Set(450)
	if r.Get() != 450 {
		t.Errorf("value %d in pair after being set to 450", r.Get())
	}
}

func TestRegisterIncrementDecrementPair(t *testing.T) {
	high := NewRegister()
	low := NewRegister()
	r := NewRegisterPair(high, low)
	r.Set(255)
	r.Increment()
	if r.Get() != 256 {
		t.Errorf("value %d in pair after incrementing 255", r.Get())
	}
	r.Decrement()
	if r.Get() != 255 {
		t.Errorf("value %d in pair after decrementing 256", r.Get())
	}
	r.Decrement()
	if r.Get() != 254 {
		t.Errorf("value %d in pair after decrementing 255", r.Get())
	}
}
