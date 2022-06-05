package main

import (
	"bufio"
	engine "event-loop"
	"event-loop/commands"
	"flag"
	"io"
	"os"
)

var inputFile = flag.String("f", "", "File with commands")

func scanner(input io.Reader, el engine.EventLoop) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		commandLine := scanner.Text()
		cmd := commands.Parse(commandLine) // parse the line to get a Command
		el.Post(cmd)
	}
}

func main() {
	flag.Parse()
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	if *inputFile != "" {
		input, err := os.Open(*inputFile)
		if err != nil {
			panic(err)
		}
		defer func(input *os.File) {
			_ = input.Close()
		}(input)
		scanner(input, *eventLoop)
	} else {
		scanner(os.Stdin, *eventLoop)
	}

	eventLoop.AwaitFinish()
}
