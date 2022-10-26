package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/danwakeem/async-http-api/agent"
	"github.com/danwakeem/async-http-api/extension"
	"github.com/golang-collections/go-datastructures/queue"
)

const INITIAL_QUEUE_SIZE = 5

func main() {
	extensionName := path.Base(os.Args[0])
	queue := queue.New(INITIAL_QUEUE_SIZE)
	ctx, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sigs
		cancel()
		s.String()
	}()

	extensionClient := extension.NewClient(os.Getenv("AWS_LAMBDA_RUNTIME_API"))
	_, err := extensionClient.Register(ctx, extensionName)
	if err != nil {
		panic(err)
	}

	internalAgent, err := agent.NewHttpAgent(queue)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	internalAgent.Init()

	var receivedRuntimeDone bool = false
	flushLogQueue := func(force bool) {
		for !(queue.Empty() && (force || receivedRuntimeDone)) {

		}
	}

	for {
		select {
		case <-ctx.Done():
		default:
			res, err := extensionClient.NextEvent(ctx)
			if err != nil {
				return
			}
			if res.EventType == extension.Shutdown {
				flushLogQueue(true)
				return
			} else {
				flushLogQueue(false)
			}
		}
	}
}
