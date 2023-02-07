package main

import (
	"fmt"
	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
)

type (
	hello      struct{ Who string }
	helloActor struct{}
)

func (state *helloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		fmt.Println("1. Started, initialize actor here")
	case *actor.Stopping:
		fmt.Println("2. Stopping, actor is about shut down")
	case *actor.Stopped:
		fmt.Println("3. Stopped, actor and its children are stopped")
	case *actor.Restarting:
		fmt.Println("4. Restarting, actor is about restart")
	case *hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

func main() {
	system := actor.NewActorSystem()
	props := actor.PropsFromProducer(func() actor.Actor { return &helloActor{} })
	pid := system.Root.Spawn(props)
	system.Root.Send(pid, &hello{Who: "Roger"})

	system.Root.Poison(pid)

	_, _ = console.ReadLine()
}
