package main

import (
    "math/rand"
)


func FlipCoin(probability float64) bool {
    return rand.Float64() < probability
}


func NormalDistribution(mean, stdDev float64) float64 {
    return rand.NormFloat64() * stdDev + mean
}


func GetBestLeverIndex(allocationHistory, resultsHistory [][]float64) int {
    numRounds := len(allocationHistory)

    if numRounds == 0 {
        return 0
    }

    numLevers := len(allocationHistory[0])

    allocationTotals := make([]float64, numLevers)
    resultsTotals := make([]float64, numLevers)

    for i := 0; i < numRounds; i++ {
        allocation := allocationHistory[i]
        results := resultsHistory[i]

        for j := 0; j < numLevers; j++ {
            allocationTotals[j] += allocation[j]
            resultsTotals[j] += results[j]
        }
    }

    leverScores := make([]float64, numLevers)

    for i := 0; i < numLevers; i++ {
        if (allocationTotals[i] == 0.0) {
            continue
        }
        leverScores[i] = resultsTotals[i] / allocationTotals[i]
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
