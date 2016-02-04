package animal

import (
	"fmt"
)

// Error from Animal method
type Error struct {
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error: %v", e.Message)
}

// Animal is a simple representation of an animal
type Animal interface {
	MakeSound() (string, error)
}

// Reptile conforms to th Animal interface
type Reptile struct {
	Name  string
	Sound string
}

// MakeSound treats the Sound attribute as if it were an event
func (r *Reptile) MakeSound() (string, error) {
	if r.Sound != "" {
		return r.Sound, nil
	}
	return "", &Error{Message: "no sound"}
}
