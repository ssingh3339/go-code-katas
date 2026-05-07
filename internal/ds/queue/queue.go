package queue

// Queue represents a FIFO queue
type Queue struct {
	items []int
}

// NewQueue creates and returns a new Queue
func NewQueue() *Queue {
	return &Queue{
		items: []int{},
	}
}

// Enqueue adds an item to the end of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item at the front of the queue
func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Front returns the item at the front of the queue without removing it
func (q *Queue) Front() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.items[0], true
}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *Queue) Size() int {
	return len(q.items)
}

// Deque represents a double-ended queue
type Deque struct {
	items []int
}

// NewDeque creates and returns a new Deque
func NewDeque() *Deque {
	return &Deque{
		items: []int{},
	}
}

// PushFront adds an item to the front of the deque
func (d *Deque) PushFront(item int) {
	d.items = append([]int{item}, d.items...)
}

// PushBack adds an item to the back of the deque
func (d *Deque) PushBack(item int) {
	d.items = append(d.items, item)
}

// PopFront removes and returns the item at the front of the deque
func (d *Deque) PopFront() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item, true
}

// PopBack removes and returns the item at the back of the deque
func (d *Deque) PopBack() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	n := len(d.items)
	item := d.items[n-1]
	d.items = d.items[:n-1]
	return item, true
}

// Front returns the item at the front of the deque without removing it
func (d *Deque) Front() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	return d.items[0], true
}

// Back returns the item at the back of the deque without removing it
func (d *Deque) Back() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	return d.items[len(d.items)-1], true
}

// IsEmpty returns true if the deque is empty
func (d *Deque) IsEmpty() bool {
	return len(d.items) == 0
}

// Size returns the number of items in the deque
func (d *Deque) Size() int {
	return len(d.items)
}
