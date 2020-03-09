package main

import (
	"math/rand"
)

type EpsilonFirstAgent struct {
	Epsilon float64
}

func (agent *EpsilonFirstAgent) Policy(state State) Action {
	var lever *Lever
	if state.Time < int(agent.Epsilon*float64(state.SimulationParameters.NumRounds)) {
		leverIndex := int(rand.Int63n(int64(state.SimulationParameters.NumLevers)))
		lever = &state.Levers[leverIndex]
	} else {
		lever = GetBestLever(state)
	}

	return Action{
		Lever: lever,
	}
}
