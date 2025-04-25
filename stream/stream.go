package stream

import "fmt"

type Stream[T any] interface {
	// ForEach applies a functions to each item in the stream.
	ForEach(fn func(item T))
	// ToSlice converts the stream to a slice.
	ToSlice() ([]T, error)
	// ToMap converts the stream to a map.
	ToMap() (map[string]T, error)
	// ToString converts the stream to a string.
	ToString() (string, error)
	// Filter filters the stream based on a predicate.
	Filter(predicate func(item T) bool) Stream[T]
	// Map applies a functions to each item in the stream and returns a new stream.
	Map(mapper func(item T) T) Stream[T]
	// Reduce reduces the stream to a single value based on an accumulator functions.
	Reduce(reducer func(acc T, item T) T, initial T) T
	// Sort sorts the stream based on a comparator functions.
	Sort(comparator func(a T, b T) int) Stream[T]
	// Limit limits the number of items in the stream.
	Limit(limit int) Stream[T]
	// Skip skips the first n items in the stream.
	Skip(n int) Stream[T]
	// Peek applies a functions to each item in the stream without modifying the stream.
	Peek(peeker func(item T) error) Stream[T]
	// Find finds the first item in the stream that matches a predicate.
	Find(predicate func(item T) bool) (T, bool)
	// Any checks if any item in the stream matches a predicate.
	Any(predicate func(item T) bool) bool
	// All checks if all items in the stream match a predicate.
	All(predicate func(item T) bool) bool
	IfPresent(predicate func(item T) bool, consumer func(item T))
}

type stream[T any] struct {
	items []T
}

func (s stream[T]) ForEach(fn func(item T)) {
	for _, item := range s.items {
		fn(item)
	}
}

func (s stream[T]) ToSlice() ([]T, error) {
	return s.items, nil
}

func (s stream[T]) ToMap() (map[string]T, error) {
	m := make(map[string]T)
	for _, item := range s.items {
		m[fmt.Sprint(item)] = item
	}
	return m, nil
}

func (s stream[T]) ToString() (string, error) {
	var str string
	for _, item := range s.items {
		str += fmt.Sprint(item) + ", "
	}
	if len(str) > 0 {
		str = str[:len(str)-2]
	}
	return str, nil
}

func (s stream[T]) Filter(predicate func(item T) bool) Stream[T] {
	var filtered []T
	for _, item := range s.items {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}
	return &stream[T]{items: filtered}
}

func (s stream[T]) Map(mapper func(item T) T) Stream[T] {
	var mapped []T
	for _, item := range s.items {
		mapped = append(mapped, mapper(item))
	}
	return &stream[T]{items: mapped}
}

func (s stream[T]) Reduce(reducer func(acc T, item T) T, initial T) T {
	for _, item := range s.items {
		initial = reducer(initial, item)
	}
	return initial
}

func (s stream[T]) Sort(comparator func(a T, b T) int) Stream[T] {
	// quick sort
	if len(s.items) < 2 {
		return s
	}

	pivot := s.items[len(s.items)/2]
	left := make([]T, 0)
	right := make([]T, 0)
	for _, item := range s.items {
		if comparator(item, pivot) < 0 {
			left = append(left, item)
		} else if comparator(item, pivot) > 0 {
			right = append(right, item)
		}
	}
	left = append(left, pivot)
	right = append(right, right...)
	s.items = append(left, right...)
	return &stream[T]{items: s.items}
}

func (s stream[T]) Limit(limit int) Stream[T] {
	if limit > len(s.items) {
		limit = len(s.items)
	}
	s.items = s.items[:limit]
	return s
}

func (s stream[T]) Skip(n int) Stream[T] {
	s.items = s.items[n:]
	return s
}

func (s stream[T]) Peek(peeker func(item T) error) Stream[T] {
	for _, item := range s.items {
		if err := peeker(item); err != nil {
			return nil
		}
	}
	return s
}

func (s stream[T]) Find(predicate func(item T) bool) (T, bool) {
	for _, item := range s.items {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

func (s stream[T]) Any(predicate func(item T) bool) bool {
	for _, item := range s.items {
		if predicate(item) {
			return true
		}
	}
	return false
}

func (s stream[T]) All(predicate func(item T) bool) bool {
	for _, item := range s.items {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func (s stream[T]) IfPresent(predicate func(item T) bool, consumer func(item T)) {
	v, ok := s.Find(predicate)
	if ok {
		consumer(v)
	}
}

func Of[T any](items ...T) Stream[T] {
	return &stream[T]{items: items}
}
