package rt

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
