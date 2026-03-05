package bot

import "sync"

type States string

const (
	StateIdle        States = "idle"
	StateWatingSerie States = "wating_serie_name"
	StateWaitingSet  States = "waiting_set"
)

type State struct {
	mu    sync.Mutex
	state map[int64]States
}

func NewState() State {
	states()
	return State{
		state: make(map[int64]States),
		mu:    sync.Mutex{},
	}
}

var mapper map[States]States

func states() {
	if mapper == nil {
		mapper = make(map[States]States)
	}
	mapper[StateIdle] = StateWatingSerie
	mapper[StateWatingSerie] = StateWaitingSet
	mapper[StateWaitingSet] = StateIdle
}

func (s State) NextState(chatID int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.state[chatID] = mapper[s.state[chatID]]
}
