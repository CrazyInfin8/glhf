package ebiten

import (
	"github.com/crazyinfin8/glhf/driver"
	"github.com/hajimehoshi/ebiten/v2"
)

type WindowProvider struct {
	Config *driver.WindowProviderConfig
}

func init() {
	driver.Drivers.WindowProvider = &WindowProvider{}
}

// --- Implement [driver.WindowProvider]

func (w *WindowProvider) Init(cfg *driver.WindowProviderConfig) {
	w.Config = cfg

	if cfg.StageWidth != 0 || cfg.StageHeight != 0 {
		ebiten.SetWindowSize(cfg.StageWidth, cfg.StageHeight)
	}
}

func (w *WindowProvider) Start() {
	ebiten.RunGame(w)
}

// --- Implement [ebiten.Game]

func (w *WindowProvider) Update() error {
	w.Config.UpdateFn()
	return nil
}

func (w *WindowProvider) Draw(screen *ebiten.Image) {
	w.Config.RenderFn(Graphic{screen, w.Config.PixelPerfect})
}

func (w *WindowProvider) Layout(width, height int) (newWidth, newHeight int) {
	return w.Config.ResizeFn(width, height)
}

func (WindowProvider) WindowSize() (width, height int) { return ebiten.WindowSize() }
