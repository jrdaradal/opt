package ds

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

func QueueFrom[T any](items []T) *Queue[T] {
	q := NewQueue[T]()
	for _, item := range items {
		q.Enqueue(item)
	}
	return q
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
	front := q.items[0]
	q.items = q.items[1:]
	return front
}

func (q Queue[T]) Front() T {
	return q.items[0]
}

func (q Queue[T]) Len() int {
	return len(q.items)
}

func (q Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}
