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

func RayColor(r Ray, world Hittable) Vec3 {
	rec, hit := world.Hit(r, 0.0, math.Inf(0))
	if hit {
		return rec.Normal.Add(NewVec3(1.0, 1.0, 1.0)).MulS(0.5)
	}
	t := 0.5 * (r.Dir.Unit().Y + 1.0)
	return NewVec3(1.0, 1.0, 1.0).MulS(1.0 - t).Add(NewVec3(0.5, 0.7, 1.0).MulS(t))
}
