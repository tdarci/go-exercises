package bookdetails

import (
	"errors"
	"io"
	"strings"

	"github.com/tdarci/go-exercises/books/data"
	"github.com/tdarci/go-exercises/books/sloth"
)

var (
	ErrNotFound = errors.New("not_found")
)

type Service struct {
	sloth *sloth.Sloth
}

func NewService() *Service {
	return &Service{
		sloth: sloth.New(),
	}
}

// Get returns the contents of a given book file.
func (s *Service) Get(filename string) (io.Reader, error) {
	s.sloth.Wait()
	contents, found := data.Books[filename]
	if !found {
		return nil, ErrNotFound
	}
	s.sloth.Wait() // wait more : )
	return strings.NewReader(contents), nil
}
