package glhf

import (
	"fmt"
	"runtime"
)

// var logger = log.New(os.Stderr, "glhf: ", log.Lshortfile)

func warn(message string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		frame := runtime.FuncForPC(pc)
		fname := frame.Name()
		fmt.Printf("WARN (%s:%d in %s)\n  ~>\t%s\n", file, line, fname, message)
	} else {
		fmt.Printf("WARN: %s", message)
	}
}

func shortFile() {}
