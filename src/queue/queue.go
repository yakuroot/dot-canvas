package queue

type Queue[T any] struct {
	Items chan T
}

func New[T any](size int) *Queue[T] {
	q := new(Queue[T])
	q.Items = make(chan T, size)
	return q
}

func (q *Queue[T]) Append(elem T) { q.Items <- elem }
func (q *Queue[T]) Pop() T        { return <-q.Items }
func (q *Queue[T]) Size() int     { return len(q.Items) }
