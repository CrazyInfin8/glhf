package main

import (
	"time"

	"github.com/crazyinfin8/glhf"
)

type PlayState struct {
	*glhf.State
	player *Player
}

func (s *PlayState) Create() {
	s.State = glhf.NewState(s, 0)

	s.player = NewPlayer()



	s.Add(s.player)
}

func (s *PlayState) Update(elapsed time.Duration) {
	s.State.Update(elapsed)
}

