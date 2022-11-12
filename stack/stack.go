package stack

import (
	"fmt"
	"sync"
)

var EmptyStackError = fmt.Errorf("stack is empty")

type Stack[T comparable] struct {
	items []T
	mutex sync.Mutex
}

func New[T comparable]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
		mutex: sync.Mutex{},
	}
}

func (s *Stack[T]) Push(value T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() (T, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.items) <= 0 {
		return *new(T), EmptyStackError
	}
	lastValue := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return lastValue, nil
}

func (s *Stack[T]) Contains(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, v := range s.items {
		if value == v {
			return true
		}
	}
	return false
}

func (s *Stack[T]) Items() []T {
	return s.items
}
