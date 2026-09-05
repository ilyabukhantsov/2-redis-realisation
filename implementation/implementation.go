package implementation

import (
	"errors"
	"sync"
	"time"
)

type Item struct {
	value     any
	expiresAt time.Time
}

type Data struct {
	mu    sync.RWMutex
	store map[any]Item
}

func newData() *Data {
	return &Data{
		store: make(map[any]Item),
	}
}

func (d *Data) Set(key string, value any, alive_minutes int16) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if alive_minutes <= 0 {
		return errors.New("Time cannot be negative!")
	}

	item := Item{
		value:     value,
		expiresAt: time.Now().Add(time.Duration(alive_minutes) * time.Minute),
	}

	d.store[key] = item
	return nil
}

func (d *Data) CheckTTL(value Item) error {

	now := time.Now()

	if now.After(value.expiresAt) {
		return errors.New("TTL Expired!")
	}

	return nil
}

func (d *Data) Get(key string) (value any, err error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	//	err := d.CheckTTL()

	if val, exists := d.store[key]; exists {
		return val, nil
	}
	return nil, errors.New("No key!")
}
