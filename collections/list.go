package collections

import (
	"fmt"
	"github.com/misenkashari/goutils/stream"
	"reflect"
)

type list[T any] struct {
	// items is the underlying slice that stores the items in the list.
	items []T
}

func EmptyList[T any]() Collection[T] {
	return &list[T]{items: []T{}}
}

func List[T any](items ...T) Collection[T] {
	return &list[T]{items: items}
}

func (l *list[T]) Add(item T) {
	l.items = append(l.items, item)
}

func (l *list[T]) Remove(item T, equals func(a, b T) bool) {
	for i, v := range l.items {
		if equals(v, item) {
			l.items = append(l.items[:i], l.items[i+1:]...)
			break
		}
	}
}

func (l *list[T]) Contains(item T, check func(a, b T) bool) bool {
	for _, v := range l.items {
		if check(v, item) {
			return true
		}
	}
	return false
}

func (l *list[T]) Size() int {
	return len(l.items)
}

func (l *list[T]) Clear() {
	l.items = nil
}

func (l *list[T]) IsEmpty() bool {
	return len(l.items) == 0
}

func (l *list[T]) Set() Collection[T] {
	uniqueItems := make([]T, 0, len(l.items))

outer:
	for _, item := range l.items {
		for _, u := range uniqueItems {
			if reflect.DeepEqual(u, item) {
				continue outer
			}
		}
		uniqueItems = append(uniqueItems, item)
	}

	return &list[T]{items: uniqueItems}
}

func (l *list[T]) ToSlice() []T {
	return l.items
}

func (l *list[T]) ToMap() (map[string]T, error) {
	m := make(map[string]T)
	for _, item := range l.items {
		m[fmt.Sprint(item)] = item
	}
	return m, nil
}

func (l *list[T]) ToString() (string, error) {
	var str string
	for _, item := range l.items {
		str += fmt.Sprint(item) + ", "
	}
	if len(str) > 0 {
		str = str[:len(str)-2]
	}
	return str, nil
}

func (l *list[T]) Stream() stream.Stream[T] {
	return stream.Of(l.items...)
}
