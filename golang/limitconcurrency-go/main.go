package main

import (
	"log"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/router"
)

type workItem struct{ i int }

const maxConcurrency = 5

func doWork(ctx actor.Context) {
	if msg, ok := ctx.Message().(*workItem); ok {
		// this is guaranteed to only execute with a max concurrency level of `maxConcurrency`
		log.Printf("%v got message %d", ctx.Self(), msg.i)
	}
}

func main() {
	pid := actor.Spawn(router.NewRoundRobinPool(maxConcurrency).WithFunc(doWork))
	for i := 0; i < 1000; i++ {
		pid.Tell(&workItem{i})
	}
	console.ReadLine()
}
