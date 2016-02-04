// Service tester
package service

import (
	"fmt"
	"testing"

	"github.com/sdlowrey/go-sandbox/animal"

	"gopkg.in/check.v1"
)

// Test registers the goCheck test harness with the go test framework.
func Test(t *testing.T) {
	check.TestingT(t)
}

// TestError is a dummy error type used for testing.
type TestError struct {
	Message string // Optional message to include in dummy message.
}

// Error satisfies the error interface for Error.
func (e *TestError) Error() string {
	return fmt.Sprintf("Error: %v", e.Message)
}

type ServiceSuite struct{}

var _ = check.Suite(&ServiceSuite{})

func (s *ServiceSuite) TestBasicService(c *check.C) {
	var inputs = []struct {
		name, sound string
	}{
		{"snake", "hiss"},
		{"lizard", ""},
	}

	service := NewReptileService()

	for _, input := range inputs {
		err := service.Add(animal.Reptile{input.name, input.sound})
		c.Assert(err, check.IsNil)
		// service.ListAll()
		rep, err := service.Get(input.name)
		// fmt.Printf("dbg: %+v\n", rep)
		c.Assert(rep.Name, check.Equals, input.name)
	}

	name := "snake"
	service.Remove(name)
	rep, err := service.Get(name)
	fmt.Printf("dbg: %+v\n", rep)
	c.Assert(err, check.NotNil)

}
