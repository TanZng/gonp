package components

type Velocity struct {
	Dx    float32
	Dy    float32
	Speed float32
}

func (a *Velocity) Mask() uint64 {
	return MaskSize
}

func (a *Velocity) IncreaseSpeed() {
	if a.Speed < 8 {
		a.Speed += 0.5
	}
}
