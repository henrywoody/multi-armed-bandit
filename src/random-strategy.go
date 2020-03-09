package main

import (
	"math/rand"
)

type RandomAgent struct{}

func (agent *RandomAgent) Policy(state State) Action {
	leverIndex := int(rand.Int63n(int64(len(state.Levers))))
	lever := &state.Levers[leverIndex]
	return Action{
		Lever: lever,
	}
}
