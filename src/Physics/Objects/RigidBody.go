package Objects

import u "Utilities"

type RigidBody struct {
	IsAffectedByGravity bool
	IsRotationLocked bool
	IsAnchored bool
	IsColliding bool

	Mass float64
	Elasticity float64

	Velocity u.Vector2
	RVelocity float64

	Forces []u.Vector2
	NetForce u.Vector2
}
