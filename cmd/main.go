package main

import (
	"bufio"
	engine "event-loop-handler"
	"os"
)

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	if input, err := os.Open("test.txt"); err == nil {
		defer func(input *os.File) {
			_ = input.Close()
		}(input)
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := engine.Parse(commandLine) // parse the line to get a Command
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}
