package glhf

import "time"

var idEnumerator = 0

type (
	Basic struct {
		id int

		alive, active, visible, exists bool
	}
	_basic = *Basic
	IBasic interface {
		basic() *Basic
		ID() int
		Alive() bool
		Active() bool
		Visible() bool
		Exists() bool
		SetAlive(bool)
		SetActive(bool)
		SetVisible(bool)
		SetExists(bool)
		Kill()
		Revive()
		Update(time.Duration)
		Draw()
		String() string
	}
)

func NewBasic() *Basic {
	b := new(Basic)
	b.id = idEnumerator
	idEnumerator++
	b.active = true
	b.alive = true
	b.exists = true
	b.visible = true
	return b
}

func (b *Basic) basic() *Basic {
	checkNil(b, "Basic")
	return b
}

func (b *Basic) ID() int { return b.id }

func (b *Basic) Alive() bool { return b.alive }

func (b *Basic) Active() bool { return b.active }

func (b *Basic) Visible() bool { return b.visible }

func (b *Basic) Exists() bool { return b.exists }

func (b *Basic) SetAlive(alive bool) { b.alive = alive }

func (b *Basic) SetActive(active bool) { b.active = active }

func (b *Basic) SetVisible(visible bool) { b.visible = visible }

func (b *Basic) SetExists(exists bool) { b.exists = exists }

func (b *Basic) Kill() {
	b.alive = false
	b.exists = false
}

func (b *Basic) Revive() {
	b.alive = true
	b.exists = true
}

func (b *Basic) Update(time.Duration) {}

func (b *Basic) Draw() {}

func (b *Basic) String() string { return "[Basic]" }
