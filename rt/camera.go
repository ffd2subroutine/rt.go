package rt

type Camera struct {
	Origin          Vec3
	LowerLeftCorner Vec3
	Horizontal      Vec3
	Vertical        Vec3
}

func NewCamera() Camera {
	var (
		aspectRatio    float64 = 16.0 / 9.0
		viewportHeight float64 = 2.0
		viewportWidth  float64 = aspectRatio * viewportHeight
		focalLength    float64 = 1.0
	)

	origin := NewVec3(0.0, 0.0, 0.0)
	horizontal := NewVec3(viewportWidth, 0.0, 0.0)
	vertical := NewVec3(0.0, viewportHeight, 0.0)
	lowerLeftCorner := origin.Sub(horizontal.DivS(2.0)).Sub(vertical.DivS(2.0)).Sub(NewVec3(0.0, 0.0, focalLength))

	return Camera{
		Origin:          origin,
		Horizontal:      horizontal,
		Vertical:        vertical,
		LowerLeftCorner: lowerLeftCorner,
	}
}

func (c Camera) GetRay(u, v float64) Ray {
	return NewRay(c.Origin, c.LowerLeftCorner.Add(c.Horizontal.MulS(u)).Add(c.Vertical.MulS(v)).Sub(c.Origin))
}
