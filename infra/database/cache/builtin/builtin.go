package builtin

import (
	"errors"
	"sync"
)

type Builtin struct {
	mu    sync.RWMutex
	cache map[uint64][]interface{}
}

func NewBuiltin() *Builtin {
	return &Builtin{
		cache: make(map[uint64][]interface{}),
		mu:    sync.RWMutex{},
	}
}

func (builtin Builtin) Get(seriesID uint64) ([]interface{}, error) {
	builtin.mu.RLock()
	defer builtin.mu.RUnlock()
	defer func() {
		builtin.cache[seriesID] = make([]interface{}, 0)
	}()
	if setDTO, ok := builtin.cache[seriesID]; !ok {
		return setDTO, errors.New("series not found")
	} else {
		return setDTO, nil
	}
}

func (buitin Builtin) Set(seriesID uint64, setDTO interface{}) (interface{}, error) {
	buitin.mu.Lock()
	defer buitin.mu.Unlock()
	buitin.cache[seriesID] = append(buitin.cache[seriesID], setDTO)
	return seriesID, nil
}
