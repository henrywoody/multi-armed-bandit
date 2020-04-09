package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var verbose bool

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.BoolVar(&verbose, "v", false, "Decide whether or not to print information about each choice and result on each iteration.")
	flag.Parse()

	stationarySimulationOptions := SimulationOptions{StationaryLevers: true}

	randomAgent := &RandomAgent{}
	result := RunSimulation(randomAgent, stationarySimulationOptions)
	fmt.Printf("Random agent result:\t\t\t%.2f\n", result)

	epsilonGreedyAgent := &EpsilonGreedyAgent{Epsilon: 0.1}
	result = RunSimulation(epsilonGreedyAgent, stationarySimulationOptions)
	fmt.Printf("Epsilon-greedy agent result:\t\t%.2f\n", result)

	epsilonFirstAgent := &EpsilonFirstAgent{Epsilon: 0.1}
	result = RunSimulation(epsilonFirstAgent, stationarySimulationOptions)
	fmt.Printf("Epsilon-first agent result:\t\t%.2f\n", result)

	epsilonDecreasingAgent := &EpsilonDecreasingAgent{}
	result = RunSimulation(epsilonDecreasingAgent, stationarySimulationOptions)
	fmt.Printf("Epsilon-decreasing agent result:\t%.2f\n", result)

	softmaxBoltzmannAgent := &SoftmaxBoltzmannAgent{Temperature: 10}
	result = RunSimulation(softmaxBoltzmannAgent, stationarySimulationOptions)
	fmt.Printf("Softmax Boltzmann agent result:\t\t%.2f\n", result)

	vDBEAgent := &VDBEAgent{InverseSensitivity: 1.0}
	result = RunSimulation(vDBEAgent, stationarySimulationOptions)
	fmt.Printf("VDBE agent result:\t\t\t%.2f\n", result)
}
