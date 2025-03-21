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
	if GetPair(high, low) != 0 {
		t.Errorf("value %d in pair after register creation", GetPair(high, low))
	}

	SetPair(high, low, 450)
	if GetPair(high, low) != 450 {
		t.Errorf("value %d in pair after being set to 450", GetPair(high, low))
	}
}

func TestRegisterIncrementDecrementPair(t *testing.T) {
	high := NewRegister()
	low := NewRegister()
	SetPair(high, low, 255)
	IncrementPair(high, low)
	if GetPair(high, low) != 256 {
		t.Errorf("value %d in pair after incrementing 255", GetPair(high, low))
	}
	DecrementPair(high, low)
	if GetPair(high, low) != 255 {
		t.Errorf("value %d in pair after decrementing 256", GetPair(high, low))
	}
}
