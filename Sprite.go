package glhf

import (
	"image/color"
	"math"
	"time"
)

type (
	Sprite struct {
		_object
		parent ISprite

		offset Point
		origin Point
		scale  Point

		angle              float64
		angleUpdated       bool
		sinAngle, cosAngle float64

		frameCollection *FrameCollection
		frame           *Frame
		frameIndex      int
		graphic         *Graphic

		animations *AnimationController

		frameWidth, frameHeight int
	}
	_sprite = *Sprite
	ISprite interface {
		IObject
		sprite() *Sprite
	}
)

func NewSprite(parent ISprite) *Sprite {
	s := new(Sprite)

	s._object = NewObject(0, 0, 0, 0)
	s.parent = parent
	s.scale = Point{1, 1}
	s.SetScrollFactor(1, 1)

	return s
}

func (s *Sprite) Angle() float64 { return s.angle }

func (s *Sprite) SetAngle(degrees float64) {
	s.angleUpdated = s.angle != degrees
	s.angle = degrees
}

func (s *Sprite) Origin() (x, y float64) { return s.origin.XY() }
func (s *Sprite) SetOrigin(x, y float64) { s.origin.Set(x, y) }
func (s *Sprite) Scale() (x, y float64)  { return s.scale.XY() }
func (s *Sprite) SetScale(x, y float64)  { s.scale.Set(x, y) }
func (s *Sprite) Offset() (x, y float64) { return s.offset.XY() }
func (s *Sprite) SetOffset(x, y float64) { s.offset.Set(x, y) }

func (s *Sprite) UpdateHitbox() {
	s.width = s.scale.X() * float64(s.frameWidth)
	s.height = s.scale.Y() * float64(s.frameHeight)

	w, h := s.Size()
	s.offset.Set(
		-(w - float64(s.frameWidth)),
		-(h - float64(s.frameHeight)),
	)
	s.CenterOrigin()
}

func (s *Sprite) CenterOrigin() {
	s.origin.Set(
		0.5*float64(s.frameWidth),
		0.5*float64(s.frameHeight),
	)
}

func (s *Sprite) CenterOffset(adjustPosition bool) {
	s.offset.Set(
		(float64(s.frameWidth)-s.Width())*0.5,
		(float64(s.frameHeight)-s.Height())*0.5,
	)
	if adjustPosition {
		x, y := s.Position()

		x += s.offset.X()
		y += s.offset.Y()
		s.SetPosition(x, y)
	}
}

func (s *Sprite) MakeGraphic(width, height int, color color.Color) {
	s.frame = NewFrameWithColor(width, height, color)

	s.frameWidth = width
	s.frameHeight = height
	s.SetSize(float64(width), float64(height))
}

func (s *Sprite) LoadGraphics(path AssetPath) error {
	frame, err := NewFrameFromImage(path, true, false)
	if err != nil {
		return err
	}
	s.SetFrameCollection(NewFrameCollection(frame))
	return nil
}

func (s *Sprite) LoadAnimatedGraphics(path AssetPath, width, height int) error {
	frame, err := NewFrameFromImage(path, true, false)
	if err != nil {
		return err
	}
	s.SetFrameCollection(NewFrameCollectionFromAnimation(frame, width, height))
	return nil
}

func (s *Sprite) sprite() *Sprite {
	checkNil(s, "Sprite")
	checkNil(s._object, "Object")
	return s
}

func (s *Sprite) ResetFrameSize() {
	if s.frame != nil {
		s.frameWidth = s.frame.Width()
		s.frameHeight = s.frame.Height()
	}
}

func (s *Sprite) Update(elapsed time.Duration) {
	if s.animations != nil {
		s.animations.Update(elapsed)
	}
}

func (s *Sprite) Draw() {
	if s.frame == nil {
		return
	}
	for _, c := range g.cameras {
		if !c.Visible() || !c.Exists() || !s.IsOnScreen(c) {
			continue
		}

		point := s.GetScreenPosition(c)
		scaledOffset := s.offset
		scaledOffset.MultPoint(s.scale)
		point.SubPoint(scaledOffset)

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

	rect.SetPosition(
		c.scroll.X()*scrollFactor.X()-s.offset.X()+s.origin.X()-scaledOrigin.X(),
		c.scroll.Y()*scrollFactor.Y()-s.offset.Y()+s.origin.Y()-scaledOrigin.Y(),
	)

	rect.SetSize(
		float64(s.frameWidth)*s.scale.X(),
		float64(s.frameHeight)*s.scale.Y(),
	)
	return rect.RotatedBounds(s.angle, scaledOrigin)
}

func (s *Sprite) drawComplex(c *Camera, point Point) {
	mat := NewMatrixIdentity()

	mat.Translate(-s.origin.X(), -s.origin.Y())
	mat.Scale(s.scale.XY())

	if math.Mod(s.angle, 90) == 0 {
		mat.Rotate(s.angle)
	} else {
		s.updateTrig()
		mat.RotateTrig(s.sinAngle, s.cosAngle)
	}

	point.Add(s.origin.X()*s.scale.X(), s.origin.Y()*s.scale.Y())
	mat.Translate(point.XY())

	c.DrawGraphic(s.frame._graphic, mat)
}

func (s *Sprite) updateTrig() {
	if s.angleUpdated {
		s.sinAngle, s.cosAngle = math.Sincos(s.angle * ToRadians)
		s.angleUpdated = false
	}
}

func (s *Sprite) SetFrameCollection(collection *FrameCollection) {
	// s.animations.destroy()
	// s.animations = nil
	if collection == nil {
		s.frameCollection = nil
		s.frame = nil
		s.graphic = nil
		return
	}
	s.frameCollection = collection
	s.SetFrameIndex(0)
}

func (s *Sprite) FrameCollection() *FrameCollection {
	return s.frameCollection
}

func (s *Sprite) NumFrames() int { return s.frameCollection.NumFrames() }

func (s *Sprite) SetFrameIndex(index int) {
	frame, ok := s.frameCollection.GetFrame(index)
	if !ok {
		return
	}
	s.frame = frame
	s.frameIndex = index
	s.ResetFrameSize()
}

func (s *Sprite) AnimationController() *AnimationController {
	if s.animations == nil {
		if s.parent != nil {
			s.animations = NewAnimationController(s.parent)
		} else {
			s.animations = NewAnimationController(s)
		}
	}
	return s.animations
}
