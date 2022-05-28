package engine

import (
	"fmt"
)

type printCommand struct {
	arg string
}

func (p *printCommand) Execute(loop Handler) {
	fmt.Println(p.arg)
}

type palindromCommand struct {
	arg string
}

func (pal *palindromCommand) Execute(loop Handler) {
	// implement palindrome
	loop.Post(&printCommand{arg: pal.arg})
}

func Parse(commandLine string) Command {
	return &printCommand{arg: commandLine}
}
