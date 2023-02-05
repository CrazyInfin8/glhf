package glhf

type (
	Camera struct {
		iBasic
		x, y float64
		width, height int

		scroll Point

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
		c.width = g.cfg.StageWidth
	}
	if c.height <= 0 {
		c.height = g.cfg.StageHeight
	}
	// c.pixelPerfect = g.pixelPerfect
	
	// if zoom == 0 {
		// zoom = g.defaultZoom
	// }
	c.zoom = zoom
	c.initialZoom = zoom

	return &c
}

func (c *Camera)Draw() {}