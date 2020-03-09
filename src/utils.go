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

type LeverValue struct {
	Lever *Lever
	Value float64
}

type LeverMetric struct {
	Lever       *Lever
	PullCount   int
	RewardTotal float64
}

func GetLeverSampleAverages(state State) []LeverValue {
	leverValues := make([]LeverValue, len(state.Levers))
	for i, lever := range state.Levers {
		leverValues[i] = LeverValue{Lever: &lever, Value: 0}
	}

	if state.Time == 0 {
		return leverValues
	}

	leverMetrics := make(map[*Lever]LeverMetric, len(state.Levers))
	for i := 0; i < state.Time; i++ {
		action := state.ActionHistory[i]
		reward := state.RewardHistory[i]
		prevMetrics, _ := leverMetrics[action.Lever]
		leverMetrics[action.Lever] = LeverMetric{
			Lever:       action.Lever,
			PullCount:   prevMetrics.PullCount + 1,
			RewardTotal: prevMetrics.RewardTotal + float64(reward),
		}
	}

	leverValuesMap := make(map[*Lever]float64, len(state.Levers))
	for lever, metrics := range leverMetrics {
		leverValuesMap[lever] = metrics.RewardTotal / float64(metrics.PullCount)
	}

	for i, lever := range state.Levers {
		leverValues[i] = LeverValue{
			Lever: &lever,
			Value: leverValuesMap[&lever],
		}
	}

	return leverValues
}

func GetMaxLever(values []LeverValue) *Lever {
	var maxLever *Lever
	var maxValue float64

	for i, leverValue := range values {
		if i == 0 || leverValue.Value > maxValue {
			maxLever = leverValue.Lever
			maxValue = leverValue.Value
		}
	}

	return maxLever
}
