package glhf

import (
	"time"

	"github.com/crazyinfin8/glhf/math"
)

type AnimationController struct {
	parent         ISprite
	animations     map[string]*Animation
	curAnimation   *Animation
	onComplete     *Listener[*Animation]
	onFrameChanged *Listener[*Animation]
	elapsed        time.Duration
}

func NewAnimationController(parent ISprite) *AnimationController {
	actrl := new(AnimationController)

	actrl.parent = parent

	return actrl
}

func (actrl *AnimationController) AddAnimation(name string, frames []int, framerate float64, loop bool) {
	if actrl.animations == nil {
		actrl.animations = make(map[string]*Animation)
	} else if _, ok := actrl.animations[name]; ok {
		panic("Animation already exists") // TODO: is this unnecessary?
	}

	actrl.animations[name] = newAnimation(actrl, name, frames, framerate, loop)
}

func (actrl *AnimationController) Play(name string, restart, reverse bool, frame int) {
	anim, ok := actrl.animations[name]
	if ok {
		actrl.curAnimation = anim
		anim.play(restart, reverse, frame)
	}
}

func (actrl *AnimationController) Update(elapsed time.Duration) {
	println("AnimationController ~> update called")
	if actrl.curAnimation == nil {
		print("###")
		return
	}
	anim := actrl.curAnimation
	anim.update(elapsed)

	sprite := actrl.parent.sprite()

	println("Frames: ", sprite.frameIndex, anim.curFrameIndex)
	if sprite.frameIndex != anim.curFrameIndex {
		sprite.SetFrameIndex(anim.curFrameIndex)
	}
}

type Animation struct {
	parent *AnimationController

	name    string
	delay   time.Duration
	elapsed time.Duration
	frames  []int

	// curIndex is the index of the current frame in this [Animation]'s
	// timeline.
	curIndex int
	// curFrameIndex is the index of the current frame from the [Sprite]'s
	// [FrameCollection].
	curFrameIndex int

	loop     bool
	finished bool
	paused   bool
	reversed bool
}

func newAnimation(parent *AnimationController, name string, frames []int, framerate float64, loop bool) *Animation {
	anim := new(Animation)

	anim.name = name
	anim.SetFrameRate(framerate)
	anim.frames = frames
	anim.loop = loop

	return anim
}

func (anim *Animation) FrameRate() float64 { return 1 / anim.delay.Seconds() }
func (anim *Animation) SetFrameRate(fps float64) {
	anim.delay = time.Duration(float64(time.Second) / fps)
}

func (anim *Animation) Delay() time.Duration         { return anim.delay }
func (anim *Animation) SetDelay(delay time.Duration) { anim.delay = delay }

func (anim *Animation) Duration() time.Duration { return anim.delay * time.Duration(anim.NumFrames()) }

func (anim *Animation) NumFrames() int { return len(anim.frames) }

// CurrentFrameIndex returns the index of the current frame in the [Animation]'s
// timeline.
func (anim *Animation) CurrentIndex() int { return anim.curIndex }
func (anim *Animation) SetCurrentIndex(frame int) {

}

// CurrentIndex returns the index of the current frame  from the [Sprite]'s
// [FrameCollection].
func (anim *Animation) CurrentFrameIndex() int { return anim.curFrameIndex }

func (anim *Animation) play(force, reverse bool, frame int) {
	if !force && !anim.finished && anim.reversed == reverse {
		anim.paused = false
		return
	}

	anim.reversed = reverse
	anim.paused = false
	maxIndex := anim.NumFrames() - 1
	if frame < 0 {
		anim.curIndex = math.RandRange(0, anim.NumFrames())
	} else {
		if frame > maxIndex {
			frame = maxIndex
		}
		if reverse {
			frame = maxIndex - frame
		}
		anim.curIndex = frame
	}
}

func (anim *Animation) update(elapsed time.Duration) {
	println("Animation ~> update called")
	if anim.delay == 0 || anim.finished || anim.paused {
		return
	}
	anim.elapsed += elapsed

	if anim.elapsed > elapsed {
		frames := int(anim.elapsed / anim.delay)
		anim.elapsed = anim.elapsed % anim.delay
		anim.tickFrame(frames)
	}
}

func (anim *Animation) tickFrame(count int) {
	maxFrame := anim.NumFrames() - 1
	var tempFrame int
	if anim.reversed {
		tempFrame = anim.curIndex - count
	} else {
		tempFrame = anim.curIndex + count
	}
	if anim.loop {
		anim.curIndex = math.WrapInt(tempFrame, 0, anim.NumFrames())
		anim.curFrameIndex = anim.frames[anim.curIndex]
		return
	}
	if tempFrame > maxFrame || tempFrame < 0 {
		anim.finished = true
		anim.curIndex = math.Clamp(tempFrame, 0, maxFrame)
		return
	}
	anim.curIndex = tempFrame
	anim.curFrameIndex = anim.frames[anim.curIndex]
}
