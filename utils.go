package glhf

func checkNil(ptr interface{}, name string) {
	if ptr == nil {
		panic(name + " is nil. did you forget to store reference?")
	}
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}

// unwrap takes the return value and error of a function. It either panics on error, or returns the value.
func unwrap[T interface{}](t T, e error) T {
	if e != nil {
		panic(e)
	}
	return t
}