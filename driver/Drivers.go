package driver

var Drivers = struct {
	DefaultWidth,
	DefaultHeight int

	GraphicProvider
	WindowProvider
}{
	DefaultWidth:  600,
	DefaultHeight: 400,
}
