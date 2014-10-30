package core

type Gossiper interface {
	NextRound()
	Execute()
}

type CyclicGossiper interface {
	NextRound()
}

type ReactiveGossiper interface {
	Execute()
}
