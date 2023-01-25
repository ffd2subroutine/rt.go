package rt

import "math"

type Sphere struct {
	Center Vec3
	Radius float64
}

func NewSphere(center Vec3, radius float64) Sphere {
	return Sphere{
		Center: center,
		Radius: radius,
	}
}

func (s Sphere) Hit(r Ray, tmin, tmax float64) (*HitRecord, bool) {
	oc := r.Orig.Sub(s.Center)
	a := r.Dir.Dot(r.Dir)
	halfb := oc.Dot(r.Dir)
	c := oc.Dot(oc) - s.Radius*s.Radius

	discriminant := halfb*halfb - a*c
	if discriminant < 0 {
		return nil, false
	}
	sqrtd := math.Sqrt(discriminant)

	root := (-halfb - sqrtd) / a
	if root < tmin || root > tmax {
		root = (-halfb + sqrtd) / a
		if root < tmin || root > tmax {
			return nil, false
		}
	}

	rec := &HitRecord{}
	rec.T = root
	rec.P = r.At(rec.T)
	outwardNormal := rec.P.Sub(s.Center).DivS(s.Radius)
	rec.SetFaceNormal(r, outwardNormal)
	return rec, true
}
