package main

import (
	"math/rand"
)


type RandomAgent struct {}


func (agent *RandomAgent) Policy(state State) Action {
	action := Action(make([]float64, state.SimulationParameters.NumLevers))
    action[int(rand.Int63n(int64(state.SimulationParameters.NumLevers)))] = 1.0
    return action
}
