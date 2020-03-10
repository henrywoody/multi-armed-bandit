package main

import (
	"fmt"
)

type Lever struct {
	ID   int
	Pull func() Reward
}

func (lever *Lever) String() string {
	return fmt.Sprintf("<Lever: %d>", lever.ID)
}

type Action struct {
	Lever *Lever
}

func (action Action) String() string {
	return fmt.Sprintf("<Action: pull %s>", action.Lever)
}

type Reward float64

type Agent interface {
	Policy(State) Action
	EvaluateActions(State) ActionValues
}

type State struct {
	Time                 int
	Levers               []Lever
	ActionSpace          []Action
	ActionHistory        []Action
	RewardHistory        []Reward
	SimulationParameters SimulationParameters
}

type SimulationParameters struct {
	NumLevers int
	NumRounds int
}

func RunSimulation(agent Agent) float64 {
	levers := makeLevers()

	state := State{
		Levers:      levers,
		ActionSpace: makeActionSpace(levers),
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
			fmt.Printf("Reward: %f\n", float64(reward))
		}

		state.ActionHistory = append(state.ActionHistory, action)
		state.RewardHistory = append(state.RewardHistory, reward)
	}

	return getRewardTotal(state.RewardHistory)
}

func makeLevers() []Lever {
	return []Lever{
		{
			ID: 1,
			Pull: func() Reward {
				valueMean := 20.0
				valueStdDev := 5.0
				return Reward(NormalDistribution(valueMean, valueStdDev))
			},
		},
		{
			ID: 2,
			Pull: func() Reward {
				valueMean := 90.0
				valueStdDev := 10.0
				return Reward(NormalDistribution(valueMean, valueStdDev))
			},
		},
		{
			ID: 3,
			Pull: func() Reward {
				valueMean := 80.0
				valueStdDev := 20.0
				return Reward(NormalDistribution(valueMean, valueStdDev))
			},
		},
		{
			ID: 4,
			Pull: func() Reward {
				valueMean := 75.0
				valueStdDev := 30.0
				return Reward(NormalDistribution(valueMean, valueStdDev))
			},
		},
		{
			ID: 5,
			Pull: func() Reward {
				valueMean := 85.0
				valueStdDev := 30.0
				return Reward(NormalDistribution(valueMean, valueStdDev))
			},
		},
	}
}

func makeActionSpace(levers []Lever) []Action {
	actionSpace := make([]Action, len(levers))
	for i := range levers {
		lever := levers[i]
		actionSpace[i] = Action{
			Lever: &lever,
		}
	}
	return actionSpace
}

func performAction(action Action) Reward {
	return action.Lever.Pull()
}

func getRewardTotal(rewardHistory []Reward) float64 {
	total := 0.0
	for _, reward := range rewardHistory {
		total += float64(reward)
	}
	return total
}
