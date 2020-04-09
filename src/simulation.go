package main

import (
	"fmt"
)

type LeverParams map[string]float64

type Lever struct {
	ID             int
	Params         LeverParams
	GenerateReward func(LeverParams) Reward
	UpdateParams   func(LeverParams) LeverParams
}

func (lever *Lever) String() string {
	return fmt.Sprintf("<Lever: %d>", lever.ID)
}

func (lever *Lever) Pull() Reward {
	return lever.GenerateReward(lever.Params)
}

func (lever *Lever) Update() {
	lever.Params = lever.UpdateParams(lever.Params)
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

type SimulationOptions struct {
	StationaryLevers bool
}

func RunSimulation(agent Agent, options SimulationOptions) float64 {
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

		if !options.StationaryLevers {
			updateLevers(&state)
		}
	}
	return getRewardTotal(state.RewardHistory)
}

func makeLevers() []Lever {
	return []Lever{
		{
			ID: 1,
			Params: LeverParams{
				"valueMean":   20.0,
				"valueStdDev": 5.0,
			},
			GenerateReward: func(params LeverParams) Reward {
				return Reward(NormalDistribution(params["valueMean"], params["valueStdDev"]))
			},
			UpdateParams: func(params LeverParams) LeverParams {
				params["valueMean"] *= NormalDistribution(1, 0.01)
				return params
			},
		},
		{
			ID: 2,
			Params: LeverParams{
				"valueMean":   90.0,
				"valueStdDev": 10.0,
			},
			GenerateReward: func(params LeverParams) Reward {
				return Reward(NormalDistribution(params["valueMean"], params["valueStdDev"]))
			},
			UpdateParams: func(params LeverParams) LeverParams {
				params["valueMean"] *= NormalDistribution(1, 0.01)
				return params
			},
		},
		{
			ID: 3,
			Params: LeverParams{
				"valueMean":   80.0,
				"valueStdDev": 20.0,
			},
			GenerateReward: func(params LeverParams) Reward {
				return Reward(NormalDistribution(params["valueMean"], params["valueStdDev"]))
			},
			UpdateParams: func(params LeverParams) LeverParams {
				params["valueMean"] *= NormalDistribution(1, 0.01)
				return params
			},
		},
		{
			ID: 4,
			Params: LeverParams{
				"valueMean":   75.0,
				"valueStdDev": 30.0,
			},
			GenerateReward: func(params LeverParams) Reward {
				return Reward(NormalDistribution(params["valueMean"], params["valueStdDev"]))
			},
			UpdateParams: func(params LeverParams) LeverParams {
				params["valueMean"] *= NormalDistribution(1, 0.01)
				return params
			},
		},
		{
			ID: 5,
			Params: LeverParams{
				"valueMean":   85.0,
				"valueStdDev": 30.0,
			},
			GenerateReward: func(params LeverParams) Reward {
				return Reward(NormalDistribution(params["valueMean"], params["valueStdDev"]))
			},
			UpdateParams: func(params LeverParams) LeverParams {
				params["valueMean"] *= NormalDistribution(1, 0.01)
				return params
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

func updateLevers(state *State) {
	for i := range state.Levers {
		lever := state.Levers[i]
		lever.Update()
	}
}
