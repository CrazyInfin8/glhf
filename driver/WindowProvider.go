package driver

type WindowProvider interface {
	Init(cfg *WindowProviderConfig)

	Start()

	WindowSize() (width, height int)
}

type WindowProviderConfig struct {
	StageWidth, StageHeight int

	WindowMode WindowMode

	ResetTimeDeltaFn func()
	UpdateFn         func()
	RenderFn         func(target Graphic)
	ResizeFn         func(width, height int) (newWidth, newHeight int)
}

type WindowMode byte

const (
	WindowModeDefault WindowMode = iota
	WindowModeMaximized
	WindowModeMinimized
	WindowModeFullscreen
	WindowModeBorderlessFullscreen
)
