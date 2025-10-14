package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "Idle",
	StateConnected: "Connected",
	StateError:     "Error",
	StateRetrying:  "Retrying",
}

func (stateServer ServerState) String() string {
	return stateName[stateServer]
}

func main() {
	serverNetwork := verificacionDeRed(StateError)
	fmt.Println("Estado del servidor:", serverNetwork)
	secondCheck := verificacionDeRed(serverNetwork)
	fmt.Println("Estado del servidor:", secondCheck)
}

func verificacionDeRed(servidor ServerState) ServerState {
	switch servidor {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateRetrying
	default:
		panic(fmt.Errorf("estado desconocido: %v", servidor))
	}
}
