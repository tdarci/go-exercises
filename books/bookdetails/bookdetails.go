package bookdetails

import (
	"errors"
	"io"
	"strings"

	"github.com/tdarci/go-exercises/books/data"
)

var (
	ErrNotFound = errors.New("not_found")
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// Get returns the contents of a given book file. Make sure to close the returned
// reader when you are done with it.
func (s *Service) Get(filename string) (io.Reader, error) {
	contents, found := data.Books[filename]
	if !found {
		return nil, ErrNotFound
	}
	return strings.NewReader(contents), nil
}
