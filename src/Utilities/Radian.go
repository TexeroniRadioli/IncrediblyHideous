package Utilities

import "math"

// Definition of radian struct
type radian struct {
	value float64
}

// Definition of Radian interface
type Radian interface {
	Value() float64
	Equals(radian) bool
	Add(radian) radian
	Subtract(radian) radian
	MultiplyBy(float64) radian
	DivideBy(float64) radian
}

// Factory function that creates a radian from a float64 value
func MakeRadian(initialValue float64) radian {
	tpi := math.Pi * 2
	// Ensure that our Radian struct has a value between 0 and 2pi
	for initialValue < 0 {
		initialValue += tpi
	}
	for initialValue >= tpi {
		initialValue -= tpi
	}

	return radian {initialValue}
}

func (r radian) Value() float64 {
	return r.value
}

func (r radian) Equals(o radian) bool {
	return r.value == o.value
}

func (r radian) Add(op radian) radian {
	return MakeRadian(r.value + op.value)
}

func (r radian) Subtract(op radian) radian {
	return MakeRadian(r.value - op.value)
}

func (r radian) MultiplyBy(op float64) radian {
	return MakeRadian(r.value * op)
}

func (r radian) DivideBy(op float64) radian {
	return MakeRadian(r.value / op)
}
