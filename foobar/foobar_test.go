package foobar

import (
	"fmt"
	"testing"

	"github.com/sdlowrey/go-sandbox/service"
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

type FoobarSuite struct{}

var _ = check.Suite(&FoobarSuite{})

func (f *FoobarSuite) TestBasicFoobar(c *check.C) {
	cache := service.NewReptileService()
	foo := NewFoo(cache)
	err := foo.Store("turtle", "sing")
	c.Assert(err, check.IsNil)
	sound, err := foo.Find("turtle")
	c.Assert(sound, check.Equals, "sing")
}
