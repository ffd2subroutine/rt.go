package rt

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}

func (a Vec3) String() string {
	return fmt.Sprintf("%f %f %f", a.X, a.Y, a.Z)
}

func (a Vec3) Neg() Vec3 {
	return Vec3{
		X: -a.X,
		Y: -a.Y,
		Z: -a.Z,
	}
}

func (a Vec3) Add(b Vec3) Vec3 {
	return Vec3{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func (a Vec3) Sub(b Vec3) Vec3 {
	return Vec3{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func (a Vec3) Mul(b Vec3) Vec3 {
	return Vec3{
		X: a.X * b.X,
		Y: a.Y * b.Y,
		Z: a.Z * b.Z,
	}
}

func (a Vec3) MulS(b float64) Vec3 {
	return Vec3{
		X: a.X * b,
		Y: a.Y * b,
		Z: a.Z * b,
	}
}

func (a Vec3) DivS(b float64) Vec3 {
	return a.MulS(1.0 / b)
}

func (a Vec3) Cross(b Vec3) Vec3 {
	return Vec3{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

func (a Vec3) Dot(b Vec3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vec3) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

func (a Vec3) Unit() Vec3 {
	return a.DivS(a.Length())
}
