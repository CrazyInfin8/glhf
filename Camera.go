package glhf

import "github.com/crazyinfin8/glhf/driver"
import "image/color"


type (
	Camera struct {
		iBasic
		x, y          float64
		width, height int
		angle         float64

		scroll    Point
		scale     Point
		viewAngle float64

		frame Graphic

		pixelPerfect      bool
		initialZoom, zoom float64

		viewOffset Rect

		color color.Color
	}
	iCamera = ICamera
	ICamera interface {
		IBasic
		camera() *Camera
	}
)

func (c *Camera) camera() *Camera {
	checkNil(c, "Camera")

	return c
}

func NewCamera(x, y float64, w, h int, zoom float64) *Camera {
	c := Camera{
		iBasic: NewBasic(),
		x:      x, y: y, width: w, height: h,
	}
	// TODO: should width/height < 0 flip camera?
	if c.width <= 0 {
		c.width = g.cfg.StageWidth
	}
	if c.height <= 0 {
		c.height = g.cfg.StageHeight
	}
	// c.pixelPerfect = g.pixelPerfect

	if zoom == 0 {
		zoom = g.defaultZoom
	}

	c.zoom = zoom
	c.initialZoom = zoom

	c.calcViewOffset()
	c.frame = driver.Drivers.NewGraphic(c.width, c.height)

	return &c
}

func (c *Camera) DrawGraphic(src Graphic, matrix Matrix) {
	c.frame.DrawGraphic(src, matrix)
}

func (c *Camera) DrawPixels() {
}

func (c *Camera) ContainsRect(r Rect) bool {
	return (r.Right() > c.viewOffset.X) &&
		(r.X < c.viewOffset.Width) &&
		(r.Bottom() > c.viewOffset.Y) &&
		(r.Y < c.viewOffset.Height)
}

func (c *Camera) calcViewOffset() {
	c.viewOffset.X = 0.5 * float64(c.width) * (c.scale.X - c.initialZoom) / c.scale.X
	c.viewOffset.Width = float64(c.width) - c.viewOffset.X
	// viewWidth = width - 2 * viewOffsetX;

	c.viewOffset.Y = 0.5 * float64(c.height) * (c.scale.Y - c.initialZoom) / c.scale.Y
	c.viewOffset.Height = float64(c.height) - c.viewOffset.Y
	// viewWidth = height - 2 * viewOffsetY;
}

func (c *Camera) Clear() {
	c.frame.Fill(c.color)
}

func (c *Camera) SetPosition(p Point) {
	c.x, c.y = p.X, p.Y
}