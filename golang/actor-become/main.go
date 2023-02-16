package main

import (
	"fmt"

	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
)

type Hello struct {
	Who string
}
type SetBehaviorActor struct {
	actor.Behavior
}

func (state *SetBehaviorActor) Receive(context actor.Context) {
	state.Behavior.Receive(context)
}

func (state *SetBehaviorActor) Happy(context actor.Context) {
	switch msg := context.Message().(type) {
	case *Hello:
		fmt.Printf("Hello %v, I'm so happy\n", msg.Who)
		state.Become(state.Angry)

	}
}

func (state *SetBehaviorActor) Angry(context actor.Context) {
	switch msg := context.Message().(type) {
	case *Hello:
		fmt.Printf("%v, Now I am angry!\n", msg.Who)
	}
}

func NewSetBehaviorActor() actor.Actor {
	a := &SetBehaviorActor{}
	a.Become(a.Happy)
	return a
}

func main() {
	system := actor.NewActorSystem()
	props := actor.PropsFromProducer(NewSetBehaviorActor)

	pid := system.Root.Spawn(props)
	system.Root.Send(pid, &Hello{Who: "Roger"})
	system.Root.Send(pid, &Hello{Who: "Roger"})
	_, _ = console.ReadLine()
}
