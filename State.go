package glhf

type (
	State struct {
		_group
		persistantUpdate, persistentDraw, destroySubstate bool

		bgColor Color
	}
	_state = *State
	IState interface {
		IGroup
		state() *State
		Create()
	}
)

func (s *State) state() *State {
	checkNil(s, "State")
	checkNil(s._group, "Group")
	return s
}

func NewState(maxLen int) *State {
	s := new(State)
	s._group = NewTypedGroup[IBasic](maxLen)
	return s
}

func (s *State) Create() {}

func (s *State) Draw() {
	s._group.Draw()
}
