package queue

type MyCircularQueue struct {
	Head, Tail, Length int
	Data               []int
	Max                bool
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Head:   0,
		Tail:   0,
		Length: 0,
		Data:   make([]int, k),
		Max:    false,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.Data) == this.Length {
		return false
	}
	if this.Length > 0 {
		this.Tail = (this.Tail + 1) % len(this.Data)
	}
	this.Data[this.Tail] = value
	this.Length += 1
	//入队后检查队列情况
	if len(this.Data) == this.Length {
		this.Max = true
	} else {
		this.Max = false
	}
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	length := len(this.Data)
	if this.Length == 0 {
		return false
	}
	if this.Tail == this.Head {
		this.Tail = (this.Tail + 1) % length
	}
	if this.Head == length {
		this.Head = 0
	} else {
		this.Head += 1
	}
	this.Length--
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.Length == 0 {
		return -1
	}
	return this.Data[this.Head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.Length == 0 {
		return -1
	}
	return this.Data[this.Tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if this.Length == 0 {
		return true
	}
	return false
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if this.Max {
		return true
	}
	return false
}
