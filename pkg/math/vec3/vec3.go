package vec3

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (a Vec3) String() string {
	return fmt.Sprintf("%f %f %f", a.X, a.Y, a.Z)
}

func Neg(a Vec3) Vec3 {
	return Vec3{
		X: -a.X,
		Y: -a.Y,
		Z: -a.Z,
	}
}

func Add(a, b Vec3) Vec3 {
	return Vec3{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func Sub(a, b Vec3) Vec3 {
	return Vec3{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func Mul(a, b Vec3) Vec3 {
	return Vec3{
		X: a.X * b.X,
		Y: a.Y * b.Y,
		Z: a.Z * b.Z,
	}
}

func MulS(a Vec3, b float64) Vec3 {
	return Vec3{
		X: a.X * b,
		Y: a.Y * b,
		Z: a.Z * b,
	}
}

func DivS(a Vec3, b float64) Vec3 {
	return MulS(a, 1.0/b)
}

func Cross(a, b Vec3) Vec3 {
	return Vec3{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

func Dot(a, b Vec3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Length(a Vec3) float64 {
	return math.Sqrt(Dot(a, a))
}

func Unit(a Vec3) Vec3 {
	return DivS(a, Length(a))
}
