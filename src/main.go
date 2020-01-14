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
    fmt.Println(result)
}

func GetRandomAllocation(numLevers int, allocationHistory, resultsHistory [][]float64) []float64 {
    allocation := make([]float64, numLevers)
    allocation[int(rand.Int63n(int64(numLevers)))] = 1.0
    return allocation
}