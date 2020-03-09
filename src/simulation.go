package main

import (
	"fmt"
)

type Lever struct {
	Pull func() Reward
}

type Action struct {
	Lever *Lever
}

type Reward float64

type Agent interface {
	Policy(State) Action
}

type State struct {
	Time                 int
	Levers               []Lever
	ActionHistory        []Action
	RewardHistory        []Reward
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

		reward := performAction(action)

		if verbose {
			fmt.Printf("Reward: %d\n", float64(reward))
		}

		state.ActionHistory = append(state.ActionHistory, action)
		state.RewardHistory = append(state.RewardHistory, reward)
	}

	return GetRewardTotal(state.RewardHistory)
}

func setUpLevers() []Lever {
	return []Lever{
		{
			Pull: func() Reward {
				valueMean := 20.0
				valueStdDev := 5.0
				return Reward(NormalDistribution(valueMean, valueStdDev))
			},
		},
		{
			Pull: func() Reward {
				valueMean := 90.0
				valueStdDev := 10.0
				return Reward(NormalDistribution(valueMean, valueStdDev))
			},
		},
		{
			Pull: func() Reward {
				valueMean := 80.0
				valueStdDev := 20.0
				return Reward(NormalDistribution(valueMean, valueStdDev))
			},
		},
		{
			Pull: func() Reward {
				valueMean := 75.0
				valueStdDev := 30.0
				return Reward(NormalDistribution(valueMean, valueStdDev))
			},
		},
		{
			Pull: func() Reward {
				valueMean := 85.0
				valueStdDev := 30.0
				return Reward(NormalDistribution(valueMean, valueStdDev))
			},
		},
	}
}

func performAction(action Action) Reward {
	return action.Lever.Pull()
}

func GetRewardTotal(rewardHistory []Reward) float64 {
	total := 0.0
	for _, reward := range rewardHistory {
		total += float64(reward)
	}
	return total
}
