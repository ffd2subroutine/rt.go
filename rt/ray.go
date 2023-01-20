package rt

import "math"

type Ray struct {
	Orig Vec3
	Dir  Vec3
}

func NewRay(origin, direction Vec3) Ray {
	return Ray{
		Orig: origin,
		Dir:  direction,
	}
}

func (r Ray) At(t float64) Vec3 {
	return r.Orig.Add(r.Dir.MulS(t))
}

func RayColor(r Ray) Vec3 {
	t := HitSphere(NewVec3(0.0, 0.0, -1.0), 0.5, r)
	if t > 0.0 {
		n := r.At(t).Sub(NewVec3(0.0, 0.0, -1.0)).Unit()
		return NewVec3(n.X+1, n.Y+1, n.Z+1).MulS(0.5)
	}
	t = 0.5 * (r.Dir.Unit().Y + 1.0)
	return NewVec3(1.0, 1.0, 1.0).MulS(1.0 - t).Add(NewVec3(0.5, 0.7, 1.0).MulS(t))
}

func HitSphere(center Vec3, radius float64, r Ray) float64 {
	oc := r.Orig.Sub(center)
	a := r.Dir.Dot(r.Dir)
	halfb := oc.Dot(r.Dir)
	c := oc.Dot(oc) - radius*radius
	discriminant := halfb*halfb - a*c
	if discriminant < 0 {
		return -1.0
	}
	return (-halfb - math.Sqrt(discriminant)) / a
}
