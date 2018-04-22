package Objects

import u "Utilities"

type collider struct {
	center u.Vector2
	shape u.Shape
}

type Collider interface {
	GetCenter() u.Vector2
	GetShape() u.Shape
}

// Gets the geometric center of this collider
func (c collider) GetCenter() u.Vector2 {
	return c.center
}

func (c collider) GetShape() u.Shape {
	return c.shape
}

// Factory function that initializes a new collider from a given shape
func MakeCollider(shape u.Shape) Collider {
	return collider {shape.GetGeometricCenter(), shape}
}
