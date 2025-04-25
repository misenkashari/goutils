package collections

import "goutils/stream"

type Collection[T comparable] interface {
	// Add adds an item to the collections.
	Add(item T)
	// Remove removes an item from the collections.
	Remove(item T)
	// Contains checks if the collections contains an item.
	Contains(item T) bool
	// Size returns the number of items in the collections.
	Size() int
	// Clear clears the collections.
	Clear()
	// IsEmpty checks if the collections is empty.
	IsEmpty() bool
	// Set returns a new collections with distinct items.
	Set() Collection[T]
	// ToSlice converts the collections to a slice.
	ToSlice() []T
	// ToMap converts the collections to a map.
	ToMap() (map[string]T, error)
	// ToString converts the collections to a string.
	ToString() (string, error)
	// Stream returns a stream of items in the collections.
	Stream() stream.Stream[T]
}
