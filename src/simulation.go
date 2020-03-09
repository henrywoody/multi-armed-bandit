package main

import (
	"fmt"
)

type Lever struct {
	Pull func(float64) float64
}

type Action []float64

type Rewards []float64

type Agent interface {
	Policy(State) Action
}

type State struct {
	Time                 int
	Levers               []Lever
	ActionHistory        []Action
	RewardsHistory       []Rewards
	SimulationParameters SimulationParameters
}

type SimulationParameters struct {
	NumLevers int
	NumRounds int
}

func RunSimulation(agent Agent) float64 {
	levers := setUpLevers()

	state := State{
		Levers: levers,
		SimulationParameters: SimulationParameters{
			NumLevers: len(levers),
			NumRounds: 1000,
		},
	}

	for state.Time = 0; state.Time < state.SimulationParameters.NumRounds; state.Time++ {
		if verbose {
			fmt.Printf("\nRound: %d\n", state.Time)
		}

		action := agent.Policy(state)

		if verbose {
			fmt.Printf("Action: %v\n", action)
		}

		rewards := EvaluateAction(action, state.Levers)

		if verbose {
			fmt.Printf("Rewards: %v\n", rewards)
		}

		state.ActionHistory = append(state.ActionHistory, action)
		state.RewardsHistory = append(state.RewardsHistory, rewards)
	}

	return GetTotalRewards(state.RewardsHistory)
}

func setUpLevers() []Lever {
	return []Lever{
		{
			Pull: func(bet float64) float64 {
				valueMean := 20.0
				valueStdDev := 5.0
				return bet * NormalDistribution(valueMean, valueStdDev)
			},
		},
		{
			Pull: func(bet float64) float64 {
				valueMean := 90.0
				valueStdDev := 10.0
				return bet * NormalDistribution(valueMean, valueStdDev)
			},
		},
		{
			Pull: func(bet float64) float64 {
				valueMean := 80.0
				valueStdDev := 20.0
				return bet * NormalDistribution(valueMean, valueStdDev)
			},
		},
		{
			Pull: func(bet float64) float64 {
				valueMean := 75.0
				valueStdDev := 30.0
				return bet * NormalDistribution(valueMean, valueStdDev)
			},
		},
		{
			Pull: func(bet float64) float64 {
				valueMean := 85.0
				valueStdDev := 30.0
				return bet * NormalDistribution(valueMean, valueStdDev)
			},
		},
	}
}

func EvaluateAction(action Action, levers []Lever) []float64 {
	results := Rewards(make([]float64, len(levers)))

	for i, lever := range levers {
		results[i] = lever.Pull(action[i])
	}

	return results
}

func GetTotalRewards(rewardsHistory []Rewards) float64 {
	total := 0.0
	for _, rewards := range rewardsHistory {
		for _, reward := range rewards {
			total += reward
		}
	}
	return total
}
