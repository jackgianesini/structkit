package structkit

import (
	"sync"
)

var cacheInstance *cache

type cache struct {
	sync.Mutex
	data map[string]map[string][]int
}

func init() {
	cacheInstance = new(cache)
	cacheInstance.data = make(map[string]map[string][]int)
}

func getCache() *cache {
	return cacheInstance
}

func (c *cache) GetIdentifier(identifier string) map[string][]int {
	c.Lock()
	defer c.Unlock()

	value := c.data[identifier]
	if value == nil {
		c.data[identifier] = make(map[string][]int)
	}

	return c.data[identifier]
}

func (c *cache) GetField(identifier string, field string) []int {
	c.Lock()
	defer c.Unlock()

	if c.data[identifier] == nil {
		return nil
	}

	return c.data[identifier][field]
}

func (c *cache) SetField(identifier string, field string, value []int) {
	c.Lock()
	defer c.Unlock()

	c.data[identifier][field] = value
}
