package main

type EpsilonFirstAgent struct {
	Epsilon               float64
	ActionValueEstimates  ActionValues
	ActionSelectionCounts map[Action]int
}

func (agent *EpsilonFirstAgent) Policy(state State) Action {
	actionValues := agent.EvaluateActions(state)
	if state.Time < int(agent.Epsilon*float64(state.SimulationParameters.NumRounds)) {
		return GetActionRandomly(state)
	}
	return GetMaxAction(actionValues)
}

func (agent *EpsilonFirstAgent) EvaluateActions(state State) ActionValues {
	agent.ActionValueEstimates, agent.ActionSelectionCounts = UpdateActionSampleAverages(
		agent.ActionValueEstimates,
		agent.ActionSelectionCounts,
		state,
	)

	return agent.ActionValueEstimates
}
