package rt

func RayColor(r Ray) Vec3 {
	t := 0.5 * (r.Dir.Unit().Y + 1.0)
	white := Vec3{X: 1.0, Y: 1.0, Z: 1.0}
	blue := Vec3{X: 0.5, Y: 0.7, Z: 1.0}
	return white.MulS(1.0 - t).Add(blue.MulS(t))
}
