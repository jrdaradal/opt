package ds

import "github.com/jrdaradal/opt/internal/fn"

type Set[T comparable] struct {
	items map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]bool),
	}
}

func SetFrom[T comparable](items []T) *Set[T] {
	set := NewSet[T]()
	for _, item := range items {
		set.Add(item)
	}
	return set
}

func (s *Set[T]) Add(item T) {
	s.items[item] = true
}

func (s *Set[T]) AddItems(items []T) {
	for _, item := range items {
		s.Add(item)
	}
}

func (s *Set[T]) Delete(item T) {
	if fn.HasKey(s.items, item) {
		delete(s.items, item)
	}
}

func (s Set[T]) Contains(item T) bool {
	return s.items[item]
}

func (s Set[T]) Len() int {
	return len(s.items)
}

func (s Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s Set[T]) Items() []T {
	return fn.MapKeys(s.items)
}

func (s1 *Set[T]) Diff(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for item := range s1.items {
		if !fn.HasKey(s2.items, item) {
			s3.Add(item)
		}
	}
	return s3
}

func (s1 *Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for item := range s1.items {
		if fn.HasKey(s2.items, item) {
			s3.Add(item)
		}
	}
	return s3
}

func AllUnique[T comparable](items []T) bool {
	return len(items) == SetFrom(items).Len()
}

func AllSame[T comparable](items []T) bool {
	return SetFrom(items).Len() == 1
}
