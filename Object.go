package glhf

type (
	Object struct {
		iBasic
		x, y, width, height float64

		scrollFactor Point

		angle float64
	}
	iObject = IObject
	IObject interface {
		IBasic
		object() *Object
		X() float64
		Y() float64
		Width() float64
		Height() float64
		Position() (x, y float64)
		Size() (w, h float64)
		SetX(x float64)
		SetY(y float64)
		SetPosition(x, y float64)
		SetSize(w, h float64)
		String() string
		GetScreenPosition(c *Camera) Point
		ScrollFactor() Point
		SetScrollFactor(p Point)
	}
)

var (
	_ IObject = NewObject(0, 0, 0, 0)
)

func NewObject(x, y, w, h float64) *Object {
	o := Object{x: x, y: y, width: w, height: h}
	o.iBasic = NewBasic()
	return &o
}

func (o *Object) object() *Object {
	checkNil(o, "Object")
	return o
}

func (o *Object) X() float64 { return o.x }

func (o *Object) Y() float64 { return o.y }

func (o *Object) Width() float64 { return o.width }

func (o *Object) Height() float64 { return o.height }

func (o *Object) Position() (x, y float64) { return o.x, o.y }

func (o *Object) Size() (w, h float64) { return o.width, o.height }

func (o *Object) SetX(x float64) { o.x = x }

func (o *Object) SetY(y float64) { o.y = y }

func (o *Object) SetPosition(x, y float64) { o.x, o.y = x, y }

func (o *Object) SetSize(w, h float64) { o.width, o.height = w, h }

func (o *Object) String() string { return "[Object]" }

func (o *Object) GetScreenPosition(c *Camera) Point {
	if c == nil {
		c = g.GetCamera()
	}
	var p Point = Point{c.x, c.y}
	p.Sub(c.scroll.X*o.scrollFactor.X, c.scroll.Y*o.scrollFactor.Y)
	return p
}

func (o *Object) ScrollFactor() Point     { return o.scrollFactor }
func (o *Object) SetScrollFactor(p Point) { o.scrollFactor = p }
