package cache

import "time"

type cacheValue struct {
	value    string
	deadline time.Time
}

func (v *cacheValue) isExpired() bool {
	if time.Until(v.deadline).Milliseconds() >= 0 || v.deadline.IsZero() {
		return false
	}
	return true
}

type Cache struct {
	Storage map[string]cacheValue
}

func NewCache() Cache {
	return Cache{
		Storage: make(map[string]cacheValue),
	}
}

func (c Cache) Get(key string) (string, bool) {
	item, ok := c.Storage[key]
	if ok == false {
		return "", ok
	}

	if item.isExpired() {
		delete(c.Storage, key)
		return "", false
	}

	return item.value, true
}

func (c Cache) Put(key, value string) {
	c.Storage[key] = cacheValue{
		value: value,
	}
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.Storage[key] = cacheValue{
		value:    value,
		deadline: deadline,
	}
}

func (c Cache) Keys() []string {
	var keys []string

	for key, item := range c.Storage {
		if item.isExpired() {
			continue
		}
		keys = append(keys, key)
	}
	return keys
}
