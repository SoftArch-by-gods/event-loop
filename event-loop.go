package engine

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

type EventLoop struct{}

// Start create messageQueue and run command reading cycle from queue
func (el *EventLoop) Start() {}

// Post add command to queue for processing
func (el *EventLoop) Post(command Command) {}

// AwaitFinish waiting for everything commands in queue will end
func (el *EventLoop) AwaitFinish() {}
