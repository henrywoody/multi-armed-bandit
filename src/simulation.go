package main

import (
    "fmt"
)

type Lever struct {
    Pull func() float64
}


type SimulationParameters struct {
    NumLevers int
    NumRounds int
}


func RunSimulation(getAllocation func([][]float64, [][]float64, SimulationParameters) []float64) float64 {
    levers := setUpLevers()
    simParams := SimulationParameters{
        NumLevers: len(levers),
        NumRounds: 1000,
    }

    allocationHistory := [][]float64{}
    resultsHistory := [][]float64{}

    for i := 0; i < simParams.NumRounds; i++ {
        if verbose {
            fmt.Printf("\nRound: %d\n", i)
        }

        allocation := getAllocation(allocationHistory, resultsHistory, simParams)

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
