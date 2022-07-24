package cache

import "fmt"

type Cache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, error)
	Delete(key string) error
}

type CacheMemory struct {
	memory map[string]interface{}
}

func New() *CacheMemory {
	return &CacheMemory{
		memory: make(map[string]interface{}),
	}
}

func (c *CacheMemory) Set(key string, value interface{}) {
	c.memory[key] = value
}

func (c *CacheMemory) Get(key string) (interface{}, error) {
	value, ok := c.memory[key]

	if !ok {
		return nil, fmt.Errorf("the key '%s' does not exist", key)
	}

	return value, nil
}

func (c *CacheMemory) Delete(key string) error {
	_, ok := c.memory[key]

	if !ok {
		return fmt.Errorf("the key '%s' does not exist", key)
	}

	delete(c.memory, key)

	return nil
}
