package main


func GetEpsilonDecreasingAllocation(allocationHistory, resultsHistory [][]float64, simParams SimulationParameters) []float64 {
    roundNumber := len(allocationHistory)
    epsilon := 1.0 / (float64(roundNumber) + 1.0)
    return getEpsilonGreedyAllocationWithVariableEpsilon(epsilon, allocationHistory, resultsHistory, simParams)
}
