package components

type Velocity struct {
	Dx    float32
	Dy    float32
	Speed float32
}

func (a *Velocity) Mask() uint64 {
	return MaskVelocity
}

func (a *Velocity) WithDx(x float32) *Velocity {
	a.Dx = x
	return a
}

func (a *Velocity) WithDy(y float32) *Velocity {
	a.Dy = y
	return a
}

func (a *Velocity) WithSpeed(speed float32) *Velocity {
	a.Speed = speed
	return a
}

func NewVelocity() *Velocity {
	return &Velocity{}
}
