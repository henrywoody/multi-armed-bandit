package main

import (
    "fmt"
)

type Lever struct {
    Pull func() float64
}


func RunSimulation(getAllocation func(int, [][]float64, [][]float64) []float64) float64 {
    levers := setUpLevers()
    numLevers := len(levers)

    allocationHistory := [][]float64{}
    resultsHistory := [][]float64{}

    const NUM_ROUNDS = 1000

    for i := 0; i < NUM_ROUNDS; i++ {
        if verbose {
            fmt.Printf("\nRound: %d\n", i)
        }

        allocation := getAllocation(numLevers, allocationHistory, resultsHistory)

        if verbose {
            fmt.Printf("Allocation: %v\n", allocation)
        }

        results := EvaluateAllocation(allocation, levers)

        if verbose {
            fmt.Printf("Results: %v\n", results)
        }

        allocationHistory = append(allocationHistory, allocation)
        resultsHistory = append(resultsHistory, results)
    }

    return GetTotalResults(resultsHistory)
}


func setUpLevers() []Lever {
    return []Lever{
        {
            Pull: func() float64 {
                valueMean := 20.0
                valueStdDev := 5.0
                return NormalDistribution(valueMean, valueStdDev)
            },
        },
        {
            Pull: func() float64 {
                valueMean := 90.0
                valueStdDev := 10.0
                return NormalDistribution(valueMean, valueStdDev)
            },
        },
        {
            Pull: func() float64 {
                valueMean := 80.0
                valueStdDev := 20.0
                return NormalDistribution(valueMean, valueStdDev)
            },
        },
        {
            Pull: func() float64 {
                valueMean := 75.0
                valueStdDev := 30.0
                return NormalDistribution(valueMean, valueStdDev)
            },
        },
        {
            Pull: func() float64 {
                valueMean := 85.0
                valueStdDev := 30.0
                return NormalDistribution(valueMean, valueStdDev)
            },
        },
    }
}


func EvaluateAllocation(allocation []float64, levers []Lever) []float64 {
    results := make([]float64, len(levers))

    for i, lever := range levers {
        results[i] = lever.Pull() * allocation[i]
    }

    return results
}


func GetTotalResults(resultsHistory [][]float64) float64 {
    total := 0.0
    for _, results := range resultsHistory {
        for _, result := range results {
            total += result
        }
    }
    return total
}
