package main

import (
    "math/rand"
)


func NormalDistribution(mean, stdDev float64) float64 {
    return rand.NormFloat64() * stdDev + mean
}
