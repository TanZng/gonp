package components

type TypeComponent uint8

const (
	TypeNone TypeComponent = iota
	TypePlayer1
	TypePlayer2
	TypeBall
)

// Type is the actual ECS component type
type Type struct {
	Name TypeComponent
}

// Mask Satisfys the ECS Component interface
func (c *Type) Mask() uint64 {
	return MaskType
}
