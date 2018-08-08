package application

import (
	"github.com/D-Technologies/go-tokentracker/domain/receivedtransaction"
)

type State struct {
	Rts []*receivedtransactiondomain.ReceivedTransaction
	Err error
}

type App struct {
	Rstate chan *State
}

var s *State

func (a *App) SetState(state *State) {
	s = state
}

func (a *App)GetState()*State{
	return s
}



