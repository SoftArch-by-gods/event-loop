package engine

import (
	"event-loop/commands"
	"log"
	"sync"
)

type messageQueue struct {
	mu   sync.Mutex
	data []commands.Command
	wait bool

	received chan int
}

type EventLoop struct {
	mq *messageQueue

	stopLoop   chan int
	exitSignal bool
}

func (mq *messageQueue) popFromQueue() commands.Command {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	if len(mq.data) == 0 {
		mq.wait = true
		mq.mu.Unlock()
		<-mq.received
		mq.mu.Lock()
	}

	cmd := mq.data[0]
	mq.data[0] = nil
	mq.data = mq.data[1:]
	return cmd
}

// Start create messageQueue and run command reading cycle from queue
func (el *EventLoop) Start() {
	el.mq = &messageQueue{
		received: make(chan int),
	}
	el.stopLoop = make(chan int)

	go func() {
		for !el.exitSignal || !(len(el.mq.data) == 0) {
			cmd := el.mq.popFromQueue()
			cmd.Execute(el)
		}
		el.stopLoop <- 0
	}()
}

// Post add command to queue for processing
func (el *EventLoop) Post(cmd commands.Command) {
	el.mq.mu.Lock()
	defer el.mq.mu.Unlock()

	// If Post is entered after the completion of the AwaitFinish generate error
	if el.exitSignal == true {
		log.Fatalln("Requests stopped!")
	}

	el.mq.data = append(el.mq.data, cmd)
	if el.mq.wait {
		el.mq.wait = false
		el.mq.received <- 0
	}
}

type stopCommand struct{}

func (sc stopCommand) Execute(h commands.Handler) {
	h.(*EventLoop).exitSignal = true
}

// AwaitFinish waiting for everything commands in queue will end
func (el *EventLoop) AwaitFinish() {
	el.Post(stopCommand{})
	<-el.stopLoop
}
