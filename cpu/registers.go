package cpu

type Register struct {
	val byte
}

func (r *Register) set(newVal byte) {
	r.val = newVal
}

func (r *Register) value() byte {
	return r.val
}

func (r *Register) reset() {
	r.val = 0
}

func (r *Register) checkBit(bit byte) bool {
	return r.val&(1<<bit) != 0
}

func (r *Register) setBit(bit byte, set bool) {
	if set {
		r.val = r.val | (1 << bit)
	} else {
		r.val = r.val &^ (1 << bit)
	}
}

func (r *Register) increment() {
	r.val++
}

func (r *Register) decrement() {
	r.val--
}
