package main

import (
	"time"

	glhf "github.com/crazyinfin8/glhf"
)

func main() {
	glhf.NewGame(600, 400).Start(PlayState{glhf.NewState(0)})
}

type PlayState struct {
	*glhf.State
}

func (s PlayState) Create() {

}

func (s PlayState) Update(time.Duration) {

}