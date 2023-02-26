package glhf

type (
	Handler[T any] interface {
		Dispatch(T)
		DestroyHandler()
	}
	Listener[T any] struct {
		handlers    map[Handler[T]]bool
		toDestroy   []Handler[T]
		dispatching bool
	}
)

func NewListener[T any]() *Listener[T] {
	l := new(Listener[T])
	l.handlers = make(map[Handler[T]]bool)
	return l
}

func (l *Listener[T]) Add(handler Handler[T], once bool) {
	if handler == nil {
		return
	}

	if l.handlers == nil {
		l.handlers = make(map[Handler[T]]bool)
	}
	_once, ok := l.handlers[handler]
	if ok {
		if once != _once {
			panic("Cannot add handler as both once and indefinite")
		}
		return
	}
	l.handlers[handler] = once
}

func (l *Listener[T]) Remove(handler Handler[T]) bool {
	if handler == nil {
		return false
	}
	_, ok := l.handlers[handler]
	if ok {
		if l.dispatching {
			l.toDestroy = append(l.toDestroy, handler)
		} else {
			handler.DestroyHandler()
		}
		delete(l.handlers, handler)
	}
	return ok
}

func (l *Listener[T]) RemoveAll() {
	for h := range l.handlers {
		h.DestroyHandler()
		delete(l.handlers, h)
	}
}

func (l *Listener[T]) Has(handler Handler[T]) bool {
	if handler == nil {
		return false
	}
	_, ok := l.handlers[handler]
	return ok
}

func (l *Listener[T]) IsOnce(handler Handler[T]) bool {
	if handler == nil {
		return false
	}
	once, ok := l.handlers[handler]
	if !ok {
		return false
	}
	return once
}

func (l *Listener[T]) Dispatch(t T) {
	l.dispatching = true
	for handler, once := range l.handlers {
		if once {
			l.Remove(handler)
		}
		handler.Dispatch(t)
	}
	l.dispatching = false
	if len(l.toDestroy) > 0 {
		for i, h := range l.toDestroy {
			h.DestroyHandler()
			l.toDestroy[i] = nil
		}
		l.toDestroy = nil
	}
}

func (l *Listener[T]) IsEmpty() bool { return len(l.handlers) == 0 }

func (l *Listener[T]) Destroy() {
	l.RemoveAll()
	l.handlers = nil
	l.toDestroy = nil
}
