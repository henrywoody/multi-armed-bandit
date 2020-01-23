package main

import (
    "fmt"
    "math/rand"
    "time"
    "flag"
)


var verbose bool


func init() {
    rand.Seed(time.Now().UnixNano())
}


func main() {
    flag.BoolVar(&verbose, "v", false, "Decide whether or not to print information about each choice and result on each iteration.")
    flag.Parse()

    result := RunSimulation(GetRandomAllocation)
    fmt.Printf("Random strategy result:\t\t%f\n", result)

    result = RunSimulation(GetEpsilonGreedyAllocation)
    fmt.Printf("Epsilon-greedy strategy result:\t%f\n", result)

    result = RunSimulation(GetEpsilonFirstAllocation)
    fmt.Printf("Epsilon-first strategy result:\t%f\n", result)
}

func GetRandomAllocation(allocationHistory, resultsHistory [][]float64, simParams SimulationParameters) []float64 {
    allocation := make([]float64, simParams.NumLevers)
    allocation[int(rand.Int63n(int64(simParams.NumLevers)))] = 1.0
    return allocation
}