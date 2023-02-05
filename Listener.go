package glhf

type (
	Handler[T any] interface {
		Dispatch(T)
		Destroy()
	}
	Listener[T any] struct {
		handlers   map[Handler[T]]bool
		toDestroy   []Handler[T]
		dispatching bool
	}
)

func NewListener[T any]() *Listener[T] {
	l := new(Listener[T])
	l.handlers = make(map[Handler[T]]bool)
	return l
}

func (l *Listener[T])Add(handler Handler[T], once bool) {
	if handler == nil {
		return
	}
	
}