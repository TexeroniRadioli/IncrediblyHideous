package Objects

import u "Utilities"

// Definition of a Transform of a PhysicsObject
type Transform struct {
	Position u.Vector2
	Rotation u.Radian
	Scale float64
}
