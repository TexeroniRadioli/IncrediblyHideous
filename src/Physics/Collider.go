package Physics

import u "Utilities"

type Collider struct {
	center u.Vector2
	shape u.Shape
}

// Gets the geometric center of this collider
func (c Collider) GetCenter() u.Vector2 {
	return c.center
}

func (c Collider) GetShape() u.Shape {
	return c.shape
}

// Factory function that initializes a new collider from a given shape
func MakeCollider(shape u.Shape) Collider {
	return Collider {shape.GetGeometricCenter(), shape}
}
