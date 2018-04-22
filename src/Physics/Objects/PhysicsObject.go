package Objects

import (
	u "Utilities"
	"math"
)

type physicsObject struct {
	transform Transform
	collider Collider
	absoluteCollider Collider
	rigidBody RigidBody
}

type PhysicsObject interface {
	TransformCollider (shape u.Shape) u.Shape
	UpdateAbsoluteColliderPosition ()
}

func (obj *physicsObject) TransformCollider (shape u.Shape) u.Shape {
	switch t := shape.(type) {
	case u.Polygon:
		return obj.transformPolygonCollider(t)
	case u.Circle:
		return obj.transformCircleCollider(t)
	default:
		return nil
	}
}

func (obj *physicsObject) transformPolygonCollider (input u.Polygon) u.Shape {
	t := obj.transform
	localPos := input.GetPoints()
	newPos := make([]u.Vector2, len(localPos))

	// Scales the input shape by this transform's scale field
	for i := range localPos {
		newPos[i] = u.Vector2{X: localPos[i].X * t.Scale, Y: localPos[i].X * t.Scale}
	}
	// Rotates the newly scaled shape around its origin according to this transform's rotation field
	for i := range newPos {
		magnitude := newPos[i].GetMagnitude()
		initialThetaY := u.MakeRadian(math.Asin(newPos[i].Y / magnitude))
		initialThetaX := u.MakeRadian(math.Acos(newPos[i].X / magnitude))
		var interpretedTheta u.Radian

		if newPos[i].X < 0 && newPos[i].Y < 0 {
			interpretedTheta = u.MakeRadian((initialThetaX.Value() + initialThetaY.Value()) / 2.0)
		} else if newPos[i].X < 0 {
			interpretedTheta = initialThetaX
		} else {
			interpretedTheta = initialThetaY
		}

		interpretedTheta = interpretedTheta.Add(t.Rotation)
		newX := math.Cos(interpretedTheta.Value()) * magnitude
		newY := math.Sin(interpretedTheta.Value()) * magnitude

		if math.Abs(newX - math.Round(newX)) < .00000001 {
			newX = math.Round(newX)
		}
		if math.Abs(newY - math.Round(newY)) < .00000001 {
			newY = math.Round(newY)
		}

		newPos[i] = u.Vector2 {X: newX, Y: newY}
	}

	// Translates the modified shape according to this transform's position field
	for i := range newPos {
		newPos[i] = newPos[i].Add(t.Position)
	}

	return u.MakePolygon(newPos)
}

func (obj *physicsObject) transformCircleCollider(input u.Circle) u.Shape {
	t := obj.transform

	realCenter := input.GetCenter().Add(t.Position)
	scaledRadius := input.GetRadius() * t.Scale

	return u.MakeCircle(realCenter, scaledRadius)
}

func (obj *physicsObject) UpdateAbsoluteColliderPosition() {
	shape := obj.collider.GetShape()

	switch t := shape.(type) {
	case u.Polygon:
		obj.absoluteCollider = MakeCollider(t)
	case u.Circle:
		obj.absoluteCollider = MakeCollider(t)
	}
}