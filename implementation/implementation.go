package implementation

import (
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

func (d *Data) Set(key string, value any, ttl time.Duration) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if ttl <= 0 {
		return ErrInvalidTTL
	}

	item := Item{
		value:     value,
		expiresAt: time.Now().Add(ttl * time.Minute),
	}

	d.store[key] = item
	return nil
}

func (d *Data) CheckTTL(value Item) error {

	now := time.Now()

	if now.After(value.expiresAt) {
		return ErrTTLExpired
	}

	return nil
}

func (d *Data) Get(key string) (value any, err error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if val, exists := d.store[key]; exists {

		err = d.CheckTTL(val)

		if err != nil {
			return nil, err
		}

		return val.value, nil
	}
	return nil, ErrKeyNotFound
}

// Delete removes the value associated with key.
func (d *Data) Delete(key string) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, exists := d.store[key]; !exists {
		return ErrKeyNotFound
	}

	delete(d.store, key)
	return nil
}
