package internal

import (
	"fmt"

	"golang.org/x/term"
)

type Printer interface {
	Printf(format string, a ...any)
	Println(text string)
}

type TerminalWrapper struct {
	T *term.Terminal
}

func (tw TerminalWrapper) Printf(format string, args ...interface{}) {
	formatted := fmt.Sprintf(format, args...)
	tw.T.Write([]byte(formatted))
}

func (tw TerminalWrapper) Println(text string) {
	tw.T.Write([]byte(text + "\n"))
}
