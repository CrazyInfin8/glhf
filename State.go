package glhf

type (
	State struct {
		iGroup
		persistantUpdate, persistentDraw, destroySubstate bool
		bgColor Color
	}
	iState = IState
	IState interface {
		IGroup
		state() *State
		Create()
	}
)

func (s *State) state() *State {
	checkNil(s, "State")
	s.iGroup.group()
	return s
}

var _ IState = NewState(0)

func NewState(maxLen int) *State {
	s := new(State)
	s.iGroup = NewTypedGroup[IBasic](maxLen)
	return s
}

func (s *State) Create() {}

func (s *State) Draw() {
	println(s)
	s.iGroup.Draw()
}