package queue

type QueueUsingSlices struct {
	items []interface{}
}

func (sus *QueueUsingSlices) Push(element interface{}) interface{} {
	sus.items = append(sus.items, element)
	return element
}

func (sus *QueueUsingSlices) Pop() interface{} {
	element := sus.items[0]
	sus.items = sus.items[1:]
	return element
}

func (sus *QueueUsingSlices) Peek() interface{} {
	return sus.items[0]
}

func (sus *QueueUsingSlices) Len() int {
	return len(sus.items)
}
