package bookdetails

import (
	"io"
	"os"
	"path/filepath"
)

type Service struct {
	dataDirectory string
}

func NewService(dataDirectory string) *Service {
	return &Service{dataDirectory: dataDirectory}
}

func (s *Service) Get(filename string) (io.ReadCloser, error) {
	f, err := os.Open(filepath.Join(s.dataDirectory, filename))
	if err != nil {
		return nil, err
	}
	return f, nil
}
