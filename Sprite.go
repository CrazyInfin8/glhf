package glhf

import (
	"image/color"
	"math"
)

type (
	Sprite struct {
		iObject
		frame  *Frame
		offset Point
		origin Point
		scale  Point

		angle              float64
		angleUpdated       bool
		sinAngle, cosAngle float64

		frameWidth, frameHeight int
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
	s.scale = Point{1, 1}
	s.SetScrollFactor(Point{1, 1})
	return &s
}

func (s *Sprite) MakeGraphic(width, height int, color color.Color) {
	s.frame = NewFrameWithColor(width, height, color)

	s.frameWidth = width
	s.frameHeight = height
	s.SetSize(float64(width), float64(height))
}

func (s *Sprite) sprite() *Sprite { return s }

func (s *Sprite) Draw() {
	if s.frame == nil {
		return
	}
	for _, c := range g.cameras {
		if !c.Visible() || !c.Exists() || !s.IsOnScreen(c) {
			continue
		}

		point := s.GetScreenPosition(c)
		point.SubPoint(s.offset)

		s.drawComplex(c, point)
	}
}

func (s *Sprite) IsOnScreen(c *Camera) bool {
	if c == nil {
		c = g.GetCamera()
	}
	return c.ContainsRect(s.ScreenBounds(c))
}

func (s *Sprite) ScreenBounds(c *Camera) Rect {
	var rect Rect
	if c == nil {
		c = g.GetCamera()
	}
	rect.SetPosition(s.Position())

	scaledOrigin := s.origin
	scaledOrigin.MultPoint(s.scale)

	scrollFactor := s.ScrollFactor()

	rect.X += c.scroll.X*scrollFactor.X - s.offset.X + s.origin.X - scaledOrigin.X
	rect.Y += c.scroll.Y*scrollFactor.Y - s.offset.Y + s.origin.Y - scaledOrigin.Y

	rect.SetSize(float64(s.frameWidth)*s.scale.X, float64(s.frameHeight)*s.scale.Y)
	return rect.RotatedBounds(s.angle, scaledOrigin)
}

func (s *Sprite) drawComplex(c *Camera, point Point) {
	mat := Identity()

	mat.Translate(-s.origin.X, -s.origin.Y)
	mat.Scale(s.scale.X, s.scale.Y)

	if s.angle != 0 {
		s.updateTrig()
		mat.RotateTrig(s.sinAngle, s.cosAngle)
		// mat.Rotate(s.angle)
	}

	// point.Add(s.origin.X, s.origin.Y)
	mat.Translate(point.X, point.Y)

	c.DrawGraphic(s.frame.graphic, mat)
}

func (s *Sprite) updateTrig() {
	if s.angleUpdated {
		s.sinAngle, s.cosAngle = math.Sincos(s.angle * ToRadians)
		s.angleUpdated = false
	}
}

func (s *Sprite) SetAngle(degrees float64) { 
	s.angle = degrees
	s.angleUpdated = true
}

func (s *Sprite) SetOrigin(p Point) {
	s.origin = p
} 