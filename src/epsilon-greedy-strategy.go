package main

import (
    "math/rand"
)


func GetEpsilonGreedyAllocation(numLevers int, allocationHistory, resultsHistory [][]float64) []float64 {
    const epsilon = 0.1

    roundNumber := len(allocationHistory)

    var leverIndex int
    if roundNumber == 0 || FlipCoin(epsilon) {
        leverIndex = int(rand.Int63n(int64(numLevers)))
    } else {
        leverIndex = GetBestLeverIndex(allocationHistory, resultsHistory)
    }

    allocation := make([]float64, numLevers)
    allocation[leverIndex] = 1.0
    return allocation
}
