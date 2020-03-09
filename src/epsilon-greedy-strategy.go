package main

import (
	"math/rand"
)

type EpsilonGreedyAgent struct {
	Epsilon float64
}

func (agent *EpsilonGreedyAgent) Policy(state State) Action {
	return getEpsilonGreedyAllocationWithVariableEpsilon(agent.Epsilon, state)
}

func getEpsilonGreedyAllocationWithVariableEpsilon(epsilon float64, state State) []float64 {
	var leverIndex int
	if state.Time == 0 || BernoulliDistribution(epsilon) {
		leverIndex = int(rand.Int63n(int64(state.SimulationParameters.NumLevers)))
	} else {
		leverIndex = GetBestLeverIndex(state)
	}

	action := Action(make([]float64, state.SimulationParameters.NumLevers))
	action[leverIndex] = 1.0
	return action
}
