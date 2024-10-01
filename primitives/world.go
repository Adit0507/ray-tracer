package primitives

type World []Hitable

func (w *World) Add(h Hitable) {
	*w = append(*w, h)
}

func (w *World) AddAll(hitables ...Hitable) {
	for _, h := range hitables {
		w.Add(h)
	}
}

func (w *World) Count()  int{
	return len(*w)
}

func (w *World) Hit(r Ray, tMin float64, tMax float64) (bool, Hit) {
	hitAnything := false
	closest := tMax
	record := Hit{}

	for _, element := range *w {
		if element != nil {
			hit, tempRecord := element.Hit(r, tMin, closest)
			if hit {
				hitAnything = true
				closest = tempRecord.T
				record = tempRecord
			}
		
		}
	}

	return hitAnything, record
}
