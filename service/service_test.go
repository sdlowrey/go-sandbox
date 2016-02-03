// Service tester
package service

import "testing"

func TestFoo(t *testing.T) {
	s := NewAnimalService(1, "BARK")
	input := "blarg"
	want := "blah"
	if got := s.Foo(input); got != want {
		t.Errorf("Foo(%q) = %v", input, got)
	}
}
