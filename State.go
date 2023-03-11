package glhf

type (
	State struct {
		_group
		parent IState

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

func NewState(parent IState, maxLen int) *State {
	s := new(State)
	s._group = NewTypedGroup[IBasic](maxLen)
	s.parent = parent
	return s
}

func (s *State) Create() {}

func (s *State) Draw() {
	s._group.Draw()
}
