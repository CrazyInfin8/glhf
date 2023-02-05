package glhf

import (
	"time"
)

type (
	TypedGroup[T IBasic] struct {
		iBasic
		maxSize int
		// Cannot compare type T to nil, so keep array as IBasics
		members []IBasic
	}
	iTypedGroup[T IBasic] interface {
		comparable
		ITypedGroup[T]
	}
	ITypedGroup[T IBasic] interface {
		IBasic
		group() *TypedGroup[T]
		Len() int
		ForEach(fn func(member T), recurse bool)
	}

	SpriteGroup  = TypedGroup[ISprite]
	iSpriteGroup = ISpriteGroup
	ISpriteGroup = ITypedGroup[ISprite]

	Group  = TypedGroup[IBasic]
	iGroup = IGroup
	IGroup = ITypedGroup[IBasic]
)

func NewTypedGroup[T IBasic](maxLen int) *TypedGroup[T] {
	g := new(TypedGroup[T])
	g.iBasic = NewBasic()
	if maxLen > 0 {
		g.maxSize = maxLen
		g.members = make([]IBasic, 0, maxLen)
	}
	return g
}

func NewGroup(maxLen int) *Group {
	return NewTypedGroup[IBasic](maxLen)
}

func NewSpriteGroup(maxLen int) *SpriteGroup {
	return NewTypedGroup[ISprite](maxLen)
}

func (g *TypedGroup[T]) group() *TypedGroup[T] {
	checkNil(g, "TypedGroup")
	return g
}

func (g *TypedGroup[T]) Len() int {
	return len(g.members)
}

func (g *TypedGroup[T]) ForEach(fn func(member T), recurse bool) {
	for _, member := range g.members {
		fn(member.(T))
		if recurse {
			if g, ok := member.(ITypedGroup[T]); ok {
				g.ForEach(fn, recurse)
			}
		}
	}
}

func (g *TypedGroup[T]) Update(dt time.Duration) {
	for _, member := range g.members {
		if member != nil && member.Exists() && member.Active() {
			member.Update(dt)
		}
	}
}

func (g *TypedGroup[T]) Draw() {
	for _, member := range g.members {
		if member.Exists() && member.Visible() {
			member.Draw()
		}
	}
}

func (g *TypedGroup[T]) Add(newMember T) bool {
	// This appears to work even for non-nillable items like plain structs.
	if interface{}(newMember) == nil {
		return false
	}

	// Panics if basic is unset (Which is supposed to be invalid by design)
	basic := newMember.basic()

	for _, member := range g.members {
		// Cannot compare type T so compare their Basics (which should be unique)
		if basic == member.basic() {
			return false
		}
	}

	i := g.FirstNil()
	if i != -1 {
		g.members[i] = newMember
		goto onAdded
	}

	if g.maxSize > 0 && len(g.members) >= g.maxSize {
		return false
	}

	g.members = append(g.members, newMember)
onAdded:
	// if g.addedListener != nil {
	//     g.addedListeber.Dispatch(newMember)
	// }
	return true
}

func (g *TypedGroup[T]) FirstNil() int {
	for i, member := range g.members {
		if member == nil {
			return i
		}
	}
	return -1
}

func (g *TypedGroup[T]) String() string { return "[Group]" }