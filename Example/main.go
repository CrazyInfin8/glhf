package main

import (
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
	
	
	s.sprite.MakeGraphic(64, 64, glhf.Color(0xFFABCDEF).RGBA())
	s.sprite.SetOrigin(glhf.Point{32, 32})
	s.sprite.SetPosition(300, 200)
	if !s.Add(s.sprite) {
		panic("Couldn't add sprite")
	}
}

var i float64 = 0 

func (s *PlayState) Update(time.Duration) {
	i += 1
	s.sprite.SetAngle(i)
}