package main

type EpsilonGreedyAgent struct {
	Epsilon float64
}

func (agent *EpsilonGreedyAgent) Policy(state State) Action {
	return getEpsilonGreedyAllocationWithVariableEpsilon(agent, agent.Epsilon, state)
}

func getEpsilonGreedyAllocationWithVariableEpsilon(agent Agent, epsilon float64, state State) Action {
	actionValues := agent.EvaluateActions(state)
	if state.Time == 0 || BernoulliDistribution(epsilon) {
		return GetActionRandomly(state)
	}
	return GetMaxAction(actionValues)
}

func (agent *EpsilonGreedyAgent) EvaluateActions(state State) ActionValues {
	return GetActionSampleAverages(state)
}
