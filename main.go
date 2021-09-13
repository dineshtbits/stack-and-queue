package main

import (
	"container/list"
	"fmt"

	"github.com/dineshtbits/stack-and-queue/queue"
	"github.com/dineshtbits/stack-and-queue/stack"
)

func main() {
	fmt.Println("Testing Stack Using Slices")
	s := &stack.StackUsingSlices{}
	for i := 0; i < 5; i++ {
		fmt.Println("Pushing ", s.Push(i).(int))
	}
	for s.Len() > 0 {
		fmt.Println("Popping ", s.Pop().(int))
	}

	fmt.Println("Testing Stack Using List")
	s2 := list.New()
	for i := 0; i < 5; i++ {
		fmt.Println("Pushing ", s2.PushBack(i).Value)
	}
	for s2.Len() > 0 {
		back := s2.Back()
		s2.Remove(back)
		fmt.Println("Popping ", back.Value)
	}

	fmt.Println("Testing Queue Using Slices")
	q := &queue.QueueUsingSlices{}
	for i := 0; i < 5; i++ {
		fmt.Println("Enqueuing ", q.Push(i).(int))
	}
	for q.Len() > 0 {
		fmt.Println("Dequeuing ", q.Pop().(int))
	}

	fmt.Println("Testing Queue Using List")
	q2 := list.New()
	for i := 0; i < 5; i++ {
		fmt.Println("Enqueuing ", q2.PushBack(i).Value)
	}
	for q2.Len() > 0 {
		front := q2.Front()
		q2.Remove(front)
		fmt.Println("Dequeuing ", front.Value)
	}
}
