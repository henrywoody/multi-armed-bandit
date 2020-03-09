package main

type EpsilonFirstAgent struct {
	Epsilon float64
}

func (agent *EpsilonFirstAgent) Policy(state State) Action {
	if state.Time < int(agent.Epsilon*float64(state.SimulationParameters.NumRounds)) {
		return GetRandomAction(state)
	}
	actionValues := agent.EvaluateActions(state)
	return GetMaxAction(actionValues)
}

func (agent *EpsilonFirstAgent) EvaluateActions(state State) ActionValues {
	return GetActionSampleAverages(state)
}
