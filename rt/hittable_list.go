package rt

// TODO: rename it to HittableWorld instead or just World?
type HittableList struct {
	Hittables []Hittable
}

func NewHittableList() *HittableList {
	return &HittableList{
		Hittables: make([]Hittable, 0, 1024),
	}
}

func (hl *HittableList) Add(h Hittable) {
	hl.Hittables = append(hl.Hittables, h)
}

func (hl *HittableList) Hit(r Ray, tmin, tmax float64) (*HitRecord, bool) {
	rec := &HitRecord{}
	hitAnything := false
	closestSoFar := tmax

	for _, h := range hl.Hittables {
		tmpRec, hit := h.Hit(r, tmin, closestSoFar)
		if hit {
			hitAnything = true
			closestSoFar = tmpRec.T
			rec = tmpRec
		}
	}

	return rec, hitAnything
}
