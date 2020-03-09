package main

type RandomAgent struct{}

func (agent *RandomAgent) Policy(state State) Action {
	return GetRandomAction(state)
}

func (agent *RandomAgent) EvaluateActions(state State) ActionValues {
	actionValues := make(ActionValues, len(state.ActionSpace))
	for _, action := range state.ActionSpace {
		actionValues[action] = 0.0
	}
	return actionValues
}
