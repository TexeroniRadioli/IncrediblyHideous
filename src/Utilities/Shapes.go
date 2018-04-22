package Utilities

import "math"

type Shape interface {
	// Support function used in the GJK algorithm
	SupportFunc (param Vector2) Vector2
	// Returns the geometric center of this shape
	GetGeometricCenter () Vector2
	// Returns the moment of inertia of this shape
	GetMomentOfInertia () float64
}

type polygon struct {
	// The points that define this polygon
	//	The polygon is closed, so the first point in the slice counts as the last point
	points []Vector2
	// The moment of inertia for this polygon
	momentOfInertia float64
	// The geometric center for this polygon
	geometricCenter Vector2
}

type Polygon interface {
	SupportFunc (param Vector2) Vector2
	GetGeometricCenter () Vector2
	GetMomentOfInertia() float64
	GetPoints() []Vector2
}

// Returns the vertex of this polygon that yields the highest dot product with the input vector
func (p polygon) SupportFunc(param Vector2) Vector2 {
	var theMeme Vector2
	var d = -math.MaxFloat32

	for _, v := range p.points {
		dot := param.Dot(v)
		if dot > d {
			d = dot
			theMeme = v
		}
	}

	return theMeme
}

func (p polygon) GetGeometricCenter() Vector2 {
	return p.geometricCenter
}

func (p polygon) GetMomentOfInertia() float64 {
	return p.momentOfInertia
}

func (p polygon) GetPoints() []Vector2 {
	return p.points
}

// sets the moment of inertia for this polygon
func (p *polygon) setMomentOfInertia() {
	var upperSum, lowerSum, crossTemp float64
	var p0, p1 Vector2

	for i := 0; i < len(p.points) - 1; i++ {
		p0 = p.points[i].Subtract(p.geometricCenter)
		p1 = p.points[i + 1].Subtract(p.geometricCenter)

		crossTemp = CrossMagnitude(p1, p0)
		upperSum += crossTemp * (p0.Dot(p0) + p0.Dot(p1) + p1.Dot(p1))
		lowerSum += crossTemp;
	}
	p.momentOfInertia = upperSum / lowerSum / 6
}

// Sets the geometric center for this polygon
func (p *polygon) setGeometricCenter() {
	var center Vector2
	for _, v := range p.points {
		center = center.Add(v)
	}

	center = center.ScalarDivide(float64(len(p.points)))
	p.geometricCenter = center
}

type circle struct {
	center Vector2
	radius float64
}

type Circle interface {
	SupportFunc (param Vector2) Vector2
	GetGeometricCenter () Vector2
	GetMomentOfInertia () float64
	GetCenter () Vector2
	GetRadius () float64
}

func (c circle) SupportFunc(param Vector2) Vector2 {
	multiplier := c.radius / param.GetMagnitude()
	return Vector2 {param.X * multiplier, param.Y * multiplier}
}

func (c circle) GetGeometricCenter() Vector2 {
	return c.center
}

func (c circle) GetMomentOfInertia() float64 {
	return c.radius * c.radius
}

func (c circle) GetCenter() Vector2 {
	return c.center
}

func (c circle) GetRadius() float64 {
	return c.radius
}

// Factory function that creates a new polygon from a slice of points
func MakePolygon(points []Vector2) Polygon {
	var p = polygon{points: points}
	p.setGeometricCenter()
	p.setMomentOfInertia()

	return p
}

// Factory function that creates a new circle from a center point, and radius
func MakeCircle(center Vector2, radius float64) Circle {
	return circle {center, radius}
}
