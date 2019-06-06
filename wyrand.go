package wyrand

import "sync/atomic"

// WyRand is a struct that contains the WyRand state.
type WyRand struct {
	state uint64
}

// New returns a new WyRand seeded with the given value.
func New(seed uint64) *WyRand {
	return &WyRand{
		state: seed,
	}
}

// Next returns a non-negative pseudo-random 64-bit integer as an uint64 value.
func (w *WyRand) Next() uint64 {
	state := atomic.LoadUint64(&w.state)

	state += 0xa0761d6478bd642f
	t := state ^ 0xe7037ed1a0b428db*state

	return atomic.AddUint64(&w.state, (t>>64)^t)
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an uint64 value.
func (w *WyRand) Int63() uint64 {
	return w.Next() >> 1
}
