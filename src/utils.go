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

type LeverMetric struct {
	PullCount   int
	RewardTotal float64
}

func GetBestLever(state State) *Lever {
	if state.Time == 0 {
		return nil
	}

	leverMetrics := make(map[*Lever]LeverMetric, len(state.Levers))

	for i := 0; i < state.Time; i++ {
		action := state.ActionHistory[i]
		reward := state.RewardHistory[i]
		prevMetrics, _ := leverMetrics[action.Lever]
		leverMetrics[action.Lever] = LeverMetric{
			PullCount:   prevMetrics.PullCount + 1,
			RewardTotal: prevMetrics.RewardTotal + float64(reward),
		}
	}

	leverScores := make(map[*Lever]float64, len(state.Levers))

	for lever, metrics := range leverMetrics {
		leverScores[lever] = metrics.RewardTotal / float64(metrics.PullCount)
	}

	return getMaxLeverKey(leverScores)
}

func getMaxLeverKey(scores map[*Lever]float64) *Lever {
	var maxLever *Lever
	var maxValue float64

	for lever, value := range scores {
		if lever == nil || value > maxValue {
			maxLever = lever
			maxValue = value
		}
	}

	return maxLever
}
