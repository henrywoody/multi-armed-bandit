package main

import (
	"math"
)

type VDBEAgent struct {
	Epsilon               float64
	InverseSensitivity    float64 // sigma
	ActionValueEstimates  ActionValues
	ActionSelectionCounts map[Action]int
}

func (agent *VDBEAgent) Policy(state State) Action {
	agent.UpdateEpsilon(state)
	return getEpsilonGreedyAllocationWithVariableEpsilon(agent, agent.Epsilon, state)
}

func (agent *VDBEAgent) UpdateEpsilon(state State) {
	if state.Time == 0 {
		agent.Epsilon = 1.0
		return
	}

	delta := 1 / float64(len(state.ActionSpace))
	agent.Epsilon = delta*agent.GetTemporalDifferenceWeight(state) + (1.0-delta)*agent.Epsilon
}

func (agent *VDBEAgent) GetTemporalDifferenceWeight(state State) float64 {
	prevAction := state.ActionHistory[len(state.ActionHistory)-1]
	prevReward := state.RewardHistory[len(state.RewardHistory)-1]

	stepSize := 1.0 / float64(1.0+agent.ActionSelectionCounts[prevAction])
	temporalDifferenceError := float64(prevReward) - agent.ActionValueEstimates[prevAction]

	eTerm := math.Exp(-1.0 * math.Abs(stepSize*temporalDifferenceError) / agent.InverseSensitivity)
	numerator := 1.0 - eTerm
	denominator := 1.0 + eTerm
	return numerator / denominator
}

func (agent *VDBEAgent) EvaluateActions(state State) ActionValues {
	if state.Time == 0 {
		agent.ActionValueEstimates = make(ActionValues, len(state.ActionSpace))
		agent.ActionSelectionCounts = make(map[Action]int, len(state.ActionSpace))

		return agent.ActionValueEstimates
	}

	prevAction := state.ActionHistory[len(state.ActionHistory)-1]
	prevReward := state.RewardHistory[len(state.RewardHistory)-1]

	stepSize := 1.0 / (1.0 + float64(agent.ActionSelectionCounts[prevAction]))
	temporalDifferenceError := float64(prevReward) - agent.ActionValueEstimates[prevAction]
	agent.ActionValueEstimates[prevAction] += stepSize * temporalDifferenceError
	agent.ActionSelectionCounts[prevAction]++

	return agent.ActionValueEstimates
}
