package main

import (
	"time"

	. "github.com/crazyinfin8/glhf"
)

var g *Game

func main() {
	g = NewGame(600, 400)
	g.Start(&PlayState{})
}

type PlayState struct {
	*State
	sprite *Sprite
}

func (s *PlayState) Create() {
	s.State = NewState(s, 0)

	s.sprite = NewSprite(nil)

	err := s.sprite.LoadGraphics(NewAssetPath("glhf:assets/laszlo.png"))
	if err != nil {
		panic(err)
	}

	s.sprite.SetPosition(300, 200)
	s.sprite.SetScale(0.5, 0.5)
	s.sprite.UpdateHitbox()

	if !s.Add(s.sprite) {
		panic("Couldn't add sprite")
	}
}

var i float64 = 0

func (s *PlayState) Update(time.Duration) {
	i++
	s.sprite.SetAngle(i)
}
