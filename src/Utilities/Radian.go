package Utilities

import "math"

// Definition of radian struct
type radian struct {
	value float64
}

// Definition of Radian interface
type Radian interface {
	Value() float64
	Equals(Radian) bool
	Add(Radian) Radian
	Subtract(Radian) Radian
	MultiplyBy(float64) Radian
	DivideBy(float64) Radian
}

// Factory function that creates a radian from a float64 value
func MakeRadian(initialValue float64) Radian {
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

func (r radian) Equals(o Radian) bool {
	return r.value == o.Value()
}

func (r radian) Add(op Radian) Radian {
	return MakeRadian(r.value + op.Value())
}

func (r radian) Subtract(op Radian) Radian {
	return MakeRadian(r.value - op.Value())
}

func (r radian) MultiplyBy(op float64) Radian {
	return MakeRadian(r.value * op)
}

func (r radian) DivideBy(op float64) Radian {
	return MakeRadian(r.value / op)
}
