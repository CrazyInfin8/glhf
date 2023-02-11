package main

import (
	"fmt"
	"math"
	"time"

	glhf "github.com/crazyinfin8/glhf"
)

var g *glhf.Game

func main() {
	g = glhf.NewGame(600, 400)
	g.Start(&PlayState{})
}

type PlayState struct {
	*glhf.State
	sprite *glhf.Sprite
}

func (s *PlayState) Create() {
	s.State = glhf.NewState(0)
	s.State.Create()

	s.sprite = glhf.NewSprite()

	err := s.sprite.LoadGraphics(glhf.NewAssetPath("glhf:assets/laszlo.png"))
	if err != nil {
		panic(err)
	}

	w, h := s.sprite.Size()

	s.sprite.SetOrigin( w / 2, h / 2)
	s.sprite.SetPosition(300, 200)
	s.sprite.SetScale(0.5,0.5)

	if !s.Add(s.sprite) {
		panic("Couldn't add sprite")
	}
}

var i float64 = 0

func (s *PlayState) Update(time.Duration) {
	i += 1
	s.sprite.SetAngle(i)
	fmt.Printf("Angle: %f\r", math.Mod(i, 360))
}
