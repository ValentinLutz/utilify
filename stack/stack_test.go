package stack_test

import (
	"github.com/ValentinLutz/utilify/stack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Push(t *testing.T) {
	// GIVEN
	s := stack.New[string]()

	// WHEN
	s.Push("hello")
	s.Push("world")

	// THEN
	assert.Equal(t, s.Items(), []string{"hello", "world"})
}

func Test_Pop(t *testing.T) {
	// GIVEN
	s := stack.New[string]()
	s.Push("hello")
	s.Push("world")

	// WHEN
	value, err := s.Pop()

	// THEN
	assert.Nil(t, err)
	assert.Equal(t, value, "world")
	assert.Equal(t, s.Items(), []string{"hello"})
}

func Test_Pop_StringEmptyStack(t *testing.T) {
	// GIVEN
	s := stack.New[string]()

	// WHEN
	value, err := s.Pop()

	// THEN
	assert.Equal(t, err, stack.EmptyStackError)
	assert.Equal(t, value, "")
}

func Test_Pop_IntEmptyStack(t *testing.T) {
	// GIVEN
	s := stack.New[int]()

	// WHEN
	value, err := s.Pop()

	// THEN
	assert.Equal(t, err, stack.EmptyStackError)
	assert.Equal(t, value, 0)
}

func Test_Contains(t *testing.T) {
	// GIVEN
	s := stack.New[string]()
	s.Push("hello")
	s.Push("world")

	// WHEN
	contains := s.Contains("hello")

	// THEN
	assert.True(t, contains)
}

func Test_Contains_Not(t *testing.T) {
	// GIVEN
	s := stack.New[string]()
	s.Push("hello")
	s.Push("world")

	// WHEN
	contains := s.Contains("hello2")

	// THEN
	assert.False(t, contains)
}
