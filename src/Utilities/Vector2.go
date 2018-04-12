package Utilities

import "math"

// Definition of 2D Vector
type Vector2 struct {
	X float64
	Y float64
}

func (v Vector2) Add(op Vector2) Vector2 {
	return Vector2 {v.X + op.X, v.Y + op.Y}
}

func (v Vector2) Subtract(op Vector2) Vector2 {
	return Vector2 {v.X - op.X, v.Y - op.Y}
}

func (v Vector2) ScalarMultiply(op float64) Vector2 {
	return Vector2 {v.X * op, v.Y * op}
}

func (v Vector2) ScalarDivide(op float64) Vector2 {
	return Vector2 {v.X / op, v.Y / op}
}

func (v Vector2) Equals(op Vector2) bool {
	return (v.X == op.X) && (v.Y == op.Y)
}

// Gets the magnitude of a vector
func (v Vector2) GetMagnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

// Gets a unit vector that has the same direction as this vector
func (v Vector2) GetNormal() Vector2 {
	mag := v.GetMagnitude()
	return Vector2{v.X / mag, v.Y / mag}
}

// Returns the value of the dot product of 'this' vector and the input vector
func (v Vector2) Dot(op Vector2) float64 {
	return (v.X * op.X) + (v.Y * op.Y)
}

// Returns 'this' vector projected onto the input vector
func (v Vector2) Project(op Vector2) Vector2 {
	return op.GetNormal().ScalarMultiply(v.Dot(op.GetNormal()))
}

// Returns 'this' vector rotated 90 degrees clockwise
func (v Vector2) CWNormal() Vector2 {
	return Vector2 {v.Y, -v.X}
}

// Returns 'this' vector rotated 90 degrees counter-clockwise
func (v Vector2) CCWNormal() Vector2 {
	return Vector2 {-v.Y, v.X}
}

// Computes a cross (b cross c)
func DoubleCross(a Vector2, b Vector2, c Vector2) Vector2 {
	return b.ScalarMultiply(a.Dot(c)).Subtract(c.ScalarMultiply(a.Dot(b)))
}

// Computes the magnitude of the cross product of two vectors
func CrossMagnitude(a Vector2, b Vector2) float64 {
	return a.X * b.Y - b.X * a.Y
}




