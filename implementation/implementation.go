package implementation

import (
	"errors"
	"sync"
)

type Data struct {
	mu    sync.RWMutex
	store map[any]any
}

func newData() *Data {
	return &Data{
		store: make(map[any]any),
	}
}

func (d *Data) Set(key string, value any) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.store[key] = value
}

func (d *Data) Get(key string) (value any, err error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if val, exists := d.store[key]; exists {
		return val, nil
	}
	return nil, errors.New("No key!")
}
