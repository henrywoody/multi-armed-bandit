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

func GetBestLeverIndex(state State) int {
	if state.Time == 0 {
		return 0
	}

	actionTotals := make([]float64, state.SimulationParameters.NumLevers)
	rewardsTotals := make([]float64, state.SimulationParameters.NumLevers)

	for i := 0; i < state.Time; i++ {
		action := state.ActionHistory[i]
		rewards := state.RewardsHistory[i]

		for j := 0; j < state.SimulationParameters.NumLevers; j++ {
			actionTotals[j] += action[j]
			rewardsTotals[j] += rewards[j]
		}
	}

	leverScores := make([]float64, state.SimulationParameters.NumLevers)

	for i := 0; i < state.SimulationParameters.NumLevers; i++ {
		if actionTotals[i] == 0.0 {
			continue
		}
		leverScores[i] = rewardsTotals[i] / actionTotals[i]
	}

	return getMaxIndex(leverScores)
}

func getMaxIndex(array []float64) int {
	var maxIndex int
	var maxValue float64

	for i, e := range array {
		if i == 0 || e > maxValue {
			maxIndex = i
			maxValue = e
		}
	}

	return maxIndex
}
