package service

import (
	"fmt"

	"github.com/sdlowrey/go-sandbox/animal"
)

// Error in service call
type Error struct {
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error: %v", e.Message)
}

// CacheService is a simple interface for stashing things
type CacheService interface {
	Add(animal.Animal) error
	Get(name string) (animal.Animal, error)
	ListAll()
	Remove(name string) error
}

// ReptileService is crude CacheService that stores objects in a map
type ReptileService struct {
	reptiles map[string]*animal.Reptile
}

// NewReptileService creates a new Reptile cache
func NewReptileService() *ReptileService {
	return &ReptileService{}
}

// Add stores a Reptile object. The Name is the key.
func (r *ReptileService) Add(rep animal.Reptile) error {
	if r.reptiles == nil {
		// fmt.Println("dbg Add: creating new cache")
		r.reptiles = make(map[string]*animal.Reptile)
	}
	// fmt.Printf("dbg Add: adding %+v to cache\n", rep)
	r.reptiles[rep.Name] = &rep
	return nil
}

// Get retrieves a Reptile by name. Returns Error if name does not exist.
func (r *ReptileService) Get(name string) (rep *animal.Reptile, e error) {
	rep, ok := r.reptiles[name]
	if !ok {
		return nil, &Error{Message: fmt.Sprintf("%s not found", name)}
	}
	//fmt.Printf("dbg Get: found %+v\n", rep)
	return rep, nil
}

// ListAll prints everything in the cache to stdout
func (r *ReptileService) ListAll() {
	for name, rep := range r.reptiles {
		fmt.Printf("%s\t%+v\n", name, *rep)
	}
}

// Remove deletes a Reptile by name
func (r *ReptileService) Remove(name string) error {
	delete(r.reptiles, name)
	return nil
}
