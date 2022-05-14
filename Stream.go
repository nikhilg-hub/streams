package streams

import "fmt"

type numeric interface {
	int | int32 | int64 | float32 | float64
}

type Stream[T any] struct {
	list []T
}

func (stream *Stream[T]) String() string {
	return fmt.Sprint(stream.list)
}

func New[T any](array []T) *Stream[T] {
	return &Stream[T]{array}
}

// Filter: used to filter the values.
func (stream *Stream[T]) Filter(f func(element T) bool) *Stream[T] {
	newList := []T{}
	for _, v := range stream.list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return &Stream[T]{newList}
}

// ForEach : perform an operation on each element from the list.
func (stream *Stream[T]) ForEach(f func(element T)) {
	for _, v := range stream.list {
		f(v)
	}
}

func (stream *Stream[T]) Map(f func(element T) any) *Stream[any] {
	newList := []any{}

	for _, v := range stream.list {
		newList = append(newList, f(v))
	}

	return &Stream[any]{newList}
}
