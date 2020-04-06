package main

import (
	"math"
)

type SoftmaxBoltzmannAgent struct {
	Temperature           float64 // Should be greater than 0. Note: low values can lead to overflows. With payouts in ~[0, 100], a value of 1 seems to be the lowest that doesn't often overflow
	ActionValueEstimates  ActionValues
	ActionSelectionCounts map[Action]int
}

func (agent *SoftmaxBoltzmannAgent) Policy(state State) Action {
	actionValues := agent.EvaluateActions(state)

	if state.Time == 0 {
		return GetActionRandomly(state)
	}

	actionProbabilities := agent.ConvertActionsValuesToActionProbabilities(actionValues)
	return GetActionProbablistically(actionProbabilities)
}

func (agent *SoftmaxBoltzmannAgent) ConvertActionsValuesToActionProbabilities(actionValues ActionValues) ActionProbabilities {
	actionNumerators := make(map[Action]float64, len(actionValues))
	denominator := 0.0
	for action, value := range actionValues {
		actionNumerator := math.Exp(value / agent.Temperature)
		actionNumerators[action] = actionNumerator
		denominator += actionNumerator
	}

	actionProbabilities := make(ActionProbabilities, len(actionValues))
	for action := range actionValues {
		actionNumerator := actionNumerators[action]
		actionProbabilities[action] = actionNumerator / denominator
	}

	return actionProbabilities
}

func (agent *SoftmaxBoltzmannAgent) EvaluateActions(state State) ActionValues {
	agent.ActionValueEstimates, agent.ActionSelectionCounts = UpdateActionSampleAverages(
		agent.ActionValueEstimates,
		agent.ActionSelectionCounts,
		state,
	)

	return agent.ActionValueEstimates
}
