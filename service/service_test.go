// Service tester
package service

import (
	"fmt"
	"testing"

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

func (s *ServiceSuite) TestFoo(c *check.C) {
	service := NewAnimalService(1, "BARK")
	input := "blarg"
	want := "blah"
	foo := service.Foo(input)
	c.Assert(foo, check.Equals, want)
	// if got := s.Foo(input); got != want {
	// 	t.Errorf("Foo(%q) = %v", input, got)
	// }
}
