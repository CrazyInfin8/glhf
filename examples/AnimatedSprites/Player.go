package main

import (
	"time"

	"github.com/crazyinfin8/glhf"
)

type Player struct {
	*glhf.Sprite
}

func NewPlayer() *Player {
	p := new(Player)

	p.Sprite = glhf.NewSprite(p)
	err := p.LoadAnimatedGraphics(glhf.NewAssetPath("main:assets/blue.png"), 0, 0)
	if err != nil {
		panic(err)
	}

	p.SetPosition(300, 200)
	p.SetSize(22, 24)
	p.SetOffset(5, 8)
	p.UpdateHitbox()
	p.SetScale(2, 2)

	actrl := p.AnimationController()
	actrl.AddAnimation("idle", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, true)
	actrl.AddAnimation("run", []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}, 20, true)
	actrl.AddAnimation("jump", []int{23}, 20, false)
	actrl.AddAnimation("fall", []int{24}, 20, false)
	actrl.AddAnimation("double_jump", []int{25, 26, 27, 28, 29, 30, 50, 23}, 20, false)
	actrl.AddAnimation("wall_jump", []int{31, 32, 33, 34, 35}, 20, false)

	return p
}

func (p Player) Update(elapsed time.Duration) {
	p.Sprite.Update(elapsed)
	actrl := p.AnimationController()
	actrl.Play("run", false, false, 0)
	print(p.FrameCollection().NumFrames())
}
