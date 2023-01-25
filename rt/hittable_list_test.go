package rt

import "testing"

func TestAdd(t *testing.T) {
	s := NewSphere(NewVec3(3.0, 3.0, 3.0), 5.0)
	hl := NewHittableList()
	hl.Add(s)
	if got, want := len(hl.Hittables), 1; got != want {
		t.Errorf("hittable list length: got %d, want %d", got, want)
	}
	//s2, _ := hl.Hittables[0].(Sphere)
	s.Radius = 1.0
	if s != hl.Hittables[0] {
		t.Errorf("spheres not equal: have %+v, want %+v", s, hl.Hittables[0])
	}
}
