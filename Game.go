package glhf

import (
	"GLHF/driver"
	"image"

	"github.com/hajimehoshi/ebiten"
)

type Game[Img image.Image] struct {
	width, height int
	cameras []*Camera
	state IState

	drivers struct {
		driver.GraphicsDriver[Img]
	}

	pixelPerfect bool
	defaultZoom float64
}

type game = Game

func (g *game)Update(screen *ebiten.Image) error {
	g.state.Update(0)
	g.state.Draw()

	for _, camera := range g.cameras {
		opt := ebiten.DrawImageOptions{}
		opt.GeoM.Translate(camera.x, camera.y)
		screen.DrawImage(camera.frame, &opt)
	}
	return nil
}

func (g *game)Layout(w, h int) (int, int) {
	return w, h
}

// func (g *game)Draw(screen *ebiten.Image) {
//
// }

type exit byte

func (exit) Error() string { return "graceful exit" }

const exit_success = exit(0)

func (g *Game)Run(s IState) error {
	if err := ebiten.RunGame(g); err == exit_success {
		return nil
	} else {
		return err
	}
}

func createGame(width, height int) {
	if g != nil {
		panic("Game is already created")
	}
	g = new(Game)
	g.pixelPerfect = false
	g.defaultZoom = 1
}

func (g *game) Width() int {
	if g == nil {
		return 0
	}
	return g.width
}

func (g *game) Height() int {
	if g == nil {
		return 0
	}
	return g.width
}