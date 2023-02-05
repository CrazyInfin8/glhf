package glhf

import "github.com/crazyinfin8/glhf/driver"

var g *Game

type Game struct {
	cfg driver.WindowProviderConfig

	cameras    []*Camera
	state      IState
	resizeMode ResizeMode

	pixelPerfect bool
	defaultZoom  float64
}

type ResizeMode byte

const (
	ResizeModeDefault ResizeMode = iota
	ResizeModeScale
	ResizeModeResize
)

func NewGame(width, height int) *Game {
	g := new(Game)
	if width < 1 {
		width = driver.Drivers.DefaultWidth
	}
	if height < 0 {
		height = driver.Drivers.DefaultHeight
	}
	g.cfg = driver.WindowProviderConfig{
		StageWidth: width, StageHeight: height,
		WindowMode:       driver.WindowModeDefault,
		ResetTimeDeltaFn: g.resetTimeDelta,
		UpdateFn:         g.update,
		RenderFn:         g.render,
		ResizeFn:         g.resize,
	}

	driver.Drivers.WindowProvider.Init(&g.cfg)

	return g
}

func (g *Game) resetTimeDelta() {}

func (g *Game) update() {

}

func (g *Game) render(target driver.Graphic) {

}

func (g *Game) resize(width, height int) (newWidth, newHeight int) {
	switch g.resizeMode {
	default:
		fallthrough
	case ResizeModeDefault, ResizeModeScale:
		return g.cfg.StageWidth, g.cfg.StageHeight
	case ResizeModeResize:
		g.cfg.StageWidth, g.cfg.StageHeight = width, height
		return width, height
	}
}

func (g *Game) Start(state IState) {
	g.state = state
	state.Create()

	driver.Drivers.Start()
}
