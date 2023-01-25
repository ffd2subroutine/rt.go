package rt

// TODO: split these into separate files: hittable.go and hitrecord.go ?
type Hittable interface {
	Hit(r Ray, tmin, tmax float64) (*HitRecord, bool)
}

type HitRecord struct {
	P         Vec3 // hit point
	Normal    Vec3
	T         float64
	FrontFace bool
}

func (rec *HitRecord) SetFaceNormal(r Ray, outwardNormal Vec3) {
	rec.FrontFace = r.Dir.Dot(outwardNormal) < 0
	if rec.FrontFace {
		rec.Normal = outwardNormal
	} else {
		rec.Normal = outwardNormal.Neg()
	}
}
