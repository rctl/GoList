package GoList

//Queue is a basic queue implementation
type Queue struct {
	count  int
	top    *node
	bottom *node
}

//Enqueue puts an item ino the queue
func (queue *Queue) Enqueue(value interface{}) {
	queue.count++

	//Initialize of not initialized
	if queue.top == nil {
		queue.top = &node{
			Value: value,
			Child: nil,
		}
		queue.bottom = queue.top
		return
	}

	queue.bottom.Child = &node{
		Value: value,
		Child: nil,
	}

	queue.bottom = queue.bottom.Child
}

//Dequeue returns first item in the queue and removes it
func (queue *Queue) Dequeue() interface{} {
	value := queue.top.Value
	queue.top = queue.top.Child
	queue.count--
	return value
}

//Peek returns first item in the queue without removing it
func (queue *Queue) Peek() interface{} {
	return queue.top.Value
}

//Count returns the number of items in the queue
func (queue *Queue) Count() int {
	return queue.count
}
