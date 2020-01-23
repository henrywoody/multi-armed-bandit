package main

import (
    "math/rand"
)


func GetEpsilonFirstAllocation(allocationHistory, resultsHistory [][]float64, simParams SimulationParameters) []float64 {
    const epsilon = 0.1

    var leverIndex int
    roundNumber := len(allocationHistory)
    if roundNumber < int(epsilon * float64(simParams.NumRounds)) {
        leverIndex = int(rand.Int63n(int64(simParams.NumLevers)))
    } else {
        leverIndex = GetBestLeverIndex(allocationHistory, resultsHistory)
    }

    allocation := make([]float64, simParams.NumLevers)
    allocation[leverIndex] = 1.0
    return allocation
}
