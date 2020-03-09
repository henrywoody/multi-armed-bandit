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

func getEpsilonGreedyAllocationWithVariableEpsilon(epsilon float64, state State) Action {
	var lever *Lever
	if state.Time == 0 || BernoulliDistribution(epsilon) {
		leverIndex := int(rand.Int63n(int64(len(state.Levers))))
		lever = &state.Levers[leverIndex]
	} else {
		leverValues := GetLeverSampleAverages(state)
		lever = GetMaxLever(leverValues)
	}

	return Action{
		Lever: lever,
	}
}
