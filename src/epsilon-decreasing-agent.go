package main

type EpsilonDecreasingAgent struct{}

func (agent *EpsilonDecreasingAgent) Policy(state State) Action {
	epsilon := 1.0 / (float64(state.Time) + 1.0)
	return getEpsilonGreedyAllocationWithVariableEpsilon(agent, epsilon, state)
}

func (agent *EpsilonDecreasingAgent) EvaluateActions(state State) ActionValues {
	return GetActionSampleAverages(state)
}
