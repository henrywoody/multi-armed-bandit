package main

import (
    "math/rand"
)


type EpsilonFirstAgent struct {
    Epsilon float64
}


func (agent *EpsilonFirstAgent) Policy(state State) Action {
    var leverIndex int
    if state.Time < int(agent.Epsilon * float64(state.SimulationParameters.NumRounds)) {
        leverIndex = int(rand.Int63n(int64(state.SimulationParameters.NumLevers)))
    } else {
        leverIndex = GetBestLeverIndex(state)
    }

    action := Action(make([]float64, state.SimulationParameters.NumLevers))
    action[leverIndex] = 1.0
    return action
}
