package main

import (
	"math/rand"
)

func BernoulliDistribution(probability float64) bool {
	return rand.Float64() < probability
}

func NormalDistribution(mean, stdDev float64) float64 {
	return rand.NormFloat64()*stdDev + mean
}

type ActionValues map[Action]float64

type ActionMetrics map[Action]ActionMetric

type ActionMetric struct {
	TimesChosen int
	RewardTotal float64
}

func GetActionSampleAverages(state State) ActionValues {
	actionValues := make(ActionValues, len(state.ActionSpace))
	for _, action := range state.ActionSpace {
		actionValues[action] = 0
	}

	if state.Time == 0 {
		return actionValues
	}

	actionMetrics := make(ActionMetrics, len(state.ActionSpace))
	for i := 0; i < state.Time; i++ {
		action := state.ActionHistory[i]
		reward := state.RewardHistory[i]
		prevMetrics, _ := actionMetrics[action]
		actionMetrics[action] = ActionMetric{
			TimesChosen: prevMetrics.TimesChosen + 1,
			RewardTotal: prevMetrics.RewardTotal + float64(reward),
		}
	}

	for action, metrics := range actionMetrics {
		actionValues[action] = metrics.RewardTotal / float64(metrics.TimesChosen)
	}

	return actionValues
}

func GetMaxAction(actionValues ActionValues) Action {
	var maxAction Action
	var maxValue float64

	i := 0
	for action, value := range actionValues {
		if i == 0 || value > maxValue {
			maxAction = action
			maxValue = value
		}
		i++
	}

	return maxAction
}

func GetRandomAction(state State) Action {
	actionIndex := int(rand.Int63n(int64(len(state.ActionSpace))))
	action := state.ActionSpace[actionIndex]
	return action
}
