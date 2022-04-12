package cache

import (
	"sync"

	"github.com/diamondburned/arikawa/v3/discord"
)

type KeyResolvable interface {
	string | discord.UserID | discord.GuildID | discord.ChannelID | int
}

type Container[T any, K KeyResolvable] struct {
	Mu    sync.RWMutex
	Items map[K]T
}

func (c *Container[T, K]) Set(key K, elem T) {
	defer c.Mu.Unlock()
	c.Mu.Lock()
	c.Items[key] = elem
}

func (c *Container[T, K]) Has(key K) bool {
	defer c.Mu.RUnlock()
	c.Mu.RLock()
	_, ok := c.Items[key]
	return ok
}

func (c *Container[T, K]) Get(key K) (T, bool) {
	defer c.Mu.RUnlock()
	c.Mu.RLock()
	elem, ok := c.Items[key]
	return elem, ok
}

func (c *Container[T, K]) Size() int {
	defer c.Mu.RUnlock()
	c.Mu.RLock()
	return len(c.Items)
}

func (c *Container[T, K]) Remove(key K) {
	defer c.Mu.Unlock()
	c.Mu.Lock()
	delete(c.Items, key)
}
