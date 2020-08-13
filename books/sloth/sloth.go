package sloth

import (
	"math/rand"
	"time"
)

const (
	minWait time.Duration = time.Millisecond * 500
	maxWait time.Duration = time.Second * 3
)

type Sloth struct {
	randomizer *rand.Rand
}

func New() *Sloth {
	return &Sloth{randomizer: rand.New(rand.NewSource(time.Now().Unix()))}
}

func (s *Sloth) Wait() {
	time.Sleep(minWait + time.Duration(s.randomizer.Int63n(int64(maxWait-minWait))))
}
