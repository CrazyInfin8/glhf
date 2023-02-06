package main

import (
	"time"

	glhf "github.com/crazyinfin8/glhf"
)

func main() {
	glhf.NewGame(600, 400).Start(&PlayState{})
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

	if !s.Add(s.sprite) {
		panic("Couldn't add sprite")
	}
	println("my sprite is:", s.sprite)
	println("my sprite is:", s.sprite)
	println("my sprite is:", s.sprite)
}

func (s *PlayState) Update(time.Duration) {
	s.sprite.SetPosition(200, 200)
}