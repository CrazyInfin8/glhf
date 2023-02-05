package driver

import "image"

type Drivers[G GraphicsDriver[Img], Img image.Image] struct {
	Graphics GraphicsDriver[Img]
}