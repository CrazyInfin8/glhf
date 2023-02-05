package glhf

import (
	"testing"
)

func TestAssert(t *testing.T) {
	warn("Hello world")
}

func TestTypes(t *testing.T) {
	var (
		_ IBasic              = NewBasic()
		_ IObject             = NewObject(0, 0, 0, 0)
		_ ISprite             = NewSprite()
		_ ITypedGroup[IBasic] = NewTypedGroup[IBasic](0)
		_ IGroup              = NewGroup(0)
		_ ISpriteGroup        = NewSpriteGroup(0)
		_ IState              = NewState(0)
		_ ICamera             = NewCamera(0, 0, 0, 0, 0)
	)
}

func expect[T comparable](t *testing.T, v, expected T, message string) {
	if v != expected {
		t.Errorf("%s ~> Failed", message)
	} else {
		t.Logf("%s ~> Passed", message)
	}
}

type NonNillableBasic struct {
	IBasic
}

func NewNonNillableBasic() NonNillableBasic {
	return NonNillableBasic{NewBasic()}
}

func (b NonNillableBasic) String() string {
	return "[NonNillableBasic]"
}

func TestGroup(t *testing.T) {
	var (
		item1  = NewBasic()
		item2  = NewObject(0, 0, 0, 0)
		item3  = NewCamera(0, 0, 0, 0, 0)
		group1 = NewGroup(0)
		group2 = NewGroup(0)
		item4  = NewNonNillableBasic()
	)
	t.Log("item1 is ID: ", item1.ID())
	t.Log("item2 i2 ID: ", item2.ID())
	t.Log("item3 i2 ID: ", item3.ID())
	t.Log("group1 i2 ID: ", group1.ID())
	t.Log("group2 i2 ID: ", group2.ID())

	
	expect(t, group1.Add(item1), true, "Expected to add item1")
	expect(t, group1.Add(item1), false, "Expected to fail. Item1 is already added")
	expect(t, group1.Add(nil), false, "Expected to fail. Cannot add nil")
	expect(t, group1.Add(group2), true, "Expected to add group2")
	expect(t, group2.Add(item2), true, "Expected to add item2")
	expect(t, group2.Add(item3), true, "Expected to add item3")
	expect(t, group2.Add(nil), false, "Expected to fail. Cannot add nil")
	expect(t, group1.Add(item4), true, "Expect to add item4")
	expect(t, group1.Add(item4), false, "Expected to fail. Item4 is already added")

	// TODO: Test SpriteGroup or other specialized TypedGroup to ensure it only works with one type.
	// TODO: Test Groups with max length

	t.Log("ForEach without recursion...")
	group1.ForEach(func(member IBasic) {
		t.Log(member.String(), member.ID())
	}, false)

	t.Log("ForEach with recursion...")
	group1.ForEach(func(member IBasic) {
		t.Log(member.String(), member.ID())

	}, true)
}
