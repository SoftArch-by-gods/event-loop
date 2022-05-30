package commands

import (
	"fmt"
)

// Command represents actions that can be performed
// in a single event loop iteration.
type Command interface {
	Execute(handler Handler)
}

// Handler allows to sent commands to an event loop
// it’s associated with.
type Handler interface {
	Post(cmd Command)
}

type printCommand string

func (p printCommand) Execute(_ Handler) {
	fmt.Println(string(p))
}

type palindromCommand string

func (pal palindromCommand) Execute(loop Handler) {
	res := string(pal)
	for i := len(pal) - 1; i >= 0; i-- {
		res += string(pal[i])
	}
	loop.Post(printCommand(res))
}
