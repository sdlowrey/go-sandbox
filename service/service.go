package service

// Service interface
type Service interface {
	Foo(op string) (string, error)
	Bar(baz int) error
}

// AnimalService combines a database and a manager
type AnimalService struct {
	quack int
	bark  string
}

// NewAnimalService creates a new service object
func NewAnimalService(q int, b string) *AnimalService {
	return &AnimalService{
		quack: q,
		bark:  b,
	}
}

// Foo does foo
func (s *AnimalService) Foo(op string) (resp string) {
	return "blah"
}

// Bar does bar
func (s *AnimalService) Bar(baz int) (err error) {
	return nil
}
