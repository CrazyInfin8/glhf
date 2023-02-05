package glhf

import "github.com/hajimehoshi/ebiten"

type (
	Camera struct {
		iBasic
		x, y float64
		width, height int

		frame *ebiten.Image

		pixelPerfect bool
		initialZoom, zoom float64

		color Color
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
		c.width = g.Width()
	}
	if c.height <= 0 {
		c.height = g.Height()
	}
	// c.pixelPerfect = g.pixelPerfect
	var filter ebiten.Filter = ebiten.FilterLinear
	if c.pixelPerfect {
		filter = ebiten.FilterNearest
	}
	// if zoom == 0 {
		// zoom = g.defaultZoom
	// }
	c.zoom = zoom
	c.initialZoom = zoom
	c.frame = unwrap(ebiten.NewImage(c.width, c.height, filter))
	return &c
}

func (c *Camera)Draw() {}