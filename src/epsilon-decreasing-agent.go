package main

type EpsilonDecreasingAgent struct {
	ActionValueEstimates  ActionValues
	ActionSelectionCounts map[Action]int
}

func (agent *EpsilonDecreasingAgent) Policy(state State) Action {
	epsilon := 1.0 / (float64(state.Time) + 1.0)
	return getEpsilonGreedyAllocationWithVariableEpsilon(agent, epsilon, state)
}

func (agent *EpsilonDecreasingAgent) EvaluateActions(state State) ActionValues {
	agent.ActionValueEstimates, agent.ActionSelectionCounts = UpdateActionSampleAverages(
		agent.ActionValueEstimates,
		agent.ActionSelectionCounts,
		state,
	)

	return agent.ActionValueEstimates
}
