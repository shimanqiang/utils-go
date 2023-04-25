package utils

import "fmt"

/**
自增ID生成器
*/

type AutoInc struct {
	step int
	max  int
	ch   chan int
}

func NewAutoInc(step, max int) (*AutoInc, error) {
	if step <= 0 || step >= max {
		return nil, fmt.Errorf("step is zero or step greater than max")
	}
	_seqId := &AutoInc{
		step: step,
		max:  max,
		ch:   make(chan int, 64),
	}
	go _seqId.autoIncId()
	return _seqId, nil
}

func (s *AutoInc) autoIncId() {
	defer func() { recover() }()
	id := 0
	for {
		id = id + s.step
		if s.max > 0 && id > s.max {
			id = 0
			continue
		}
		s.ch <- id
	}
}

func (s *AutoInc) NextId() int {
	select {
	case id := <-s.ch:
		return id
	}
}
