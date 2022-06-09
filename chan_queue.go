package queue

// a non-threadsafe linked queue.
type chanQueue struct {
	ch chan interface{}
}

func NewChanQueue() *chanQueue {
	return &chanQueue{
		ch: make(chan interface{}, 10000),
	}
}

// Enqueue puts the given value v at the tail of the queue.
func (q *chanQueue) Enqueue(v interface{}) {
	q.ch <- v
}

// Dequeue removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *chanQueue) Dequeue() interface{} {
	v := <-q.ch
	return v
}
