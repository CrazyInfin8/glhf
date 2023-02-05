package glhf

type (
	Sprite struct {
		iObject
		frame *Frame
	}
	iSprite = ISprite
	ISprite interface {
		IObject
		sprite() *Sprite
	}
)

func NewSprite() *Sprite {
	s := Sprite{}
	s.iObject = NewObject(0, 0, 0, 0)
	return &s
}

func (s *Sprite) sprite() *Sprite { return s }
