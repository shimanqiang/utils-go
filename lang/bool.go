package lang

import "sync/atomic"

type Bool struct {
	v uint32
}

var _zeroBool bool

// NewBool creates a new Bool.
func NewBool(v bool) *Bool {
	x := &Bool{}
	if v != _zeroBool {
		x.Store(v)
	}
	return x
}

// Load atomically loads the wrapped bool.
func (x *Bool) Load() bool {
	return truthy(atomic.LoadUint32(&x.v))
}

// Store atomically stores the passed bool.
func (x *Bool) Store(n bool) {
	atomic.StoreUint32(&x.v, boolToInt(n))
}

// CAS is an atomic compare-and-swap for bool values.
func (x *Bool) CAS(o, n bool) bool {
	return atomic.CompareAndSwapUint32(&x.v, boolToInt(o), boolToInt(n))
}

// Swap atomically stores the given bool and returns the old
// value.
func (x *Bool) Swap(o bool) bool {
	return truthy(atomic.SwapUint32(&x.v, boolToInt(o)))
}

func truthy(n uint32) bool {
	return n == 1
}

func boolToInt(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}
