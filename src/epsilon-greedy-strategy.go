package main

import (
    "math/rand"
)


func GetEpsilonGreedyAllocation(allocationHistory, resultsHistory [][]float64, simParams SimulationParameters) []float64 {
    const epsilon = 0.1
    return getEpsilonGreedyAllocationWithVariableEpsilon(epsilon, allocationHistory, resultsHistory, simParams)
}


func getEpsilonGreedyAllocationWithVariableEpsilon(epsilon float64, allocationHistory, resultsHistory [][]float64, simParams SimulationParameters) []float64 {
    roundNumber := len(allocationHistory)

    var leverIndex int
    if roundNumber == 0 || FlipCoin(epsilon) {
        leverIndex = int(rand.Int63n(int64(simParams.NumLevers)))
    } else {
        leverIndex = GetBestLeverIndex(allocationHistory, resultsHistory)
    }

    allocation := make([]float64, simParams.NumLevers)
    allocation[leverIndex] = 1.0
    return allocation
}