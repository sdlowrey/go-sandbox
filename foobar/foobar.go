package foobar

import (
	"github.com/sdlowrey/go-sandbox/animal"
	"github.com/sdlowrey/go-sandbox/service"
)

// Foobar interface
type Foobar interface {
	Find(string) (string, error)
	Store(string, string) error
}

// FooThang conforms to Foobar i/f and is backed by a ReptileService
type FooThang struct {
	cache *service.ReptileService
}

// NewFoo creates a FooThang
func NewFoo(c *service.ReptileService) *FooThang {
	return &FooThang{
		cache: c,
	}
}

// Find gets Animal stuff by name
func (f *FooThang) Find(name string) (string, error) {
	rep, err := f.cache.Get(name)
	return rep.Sound, err
}

// Store saves reptile stuff
func (f *FooThang) Store(name string, sound string) error {
	thing := animal.Reptile{name, sound}
	err := f.cache.Add(thing)
	return err
}
