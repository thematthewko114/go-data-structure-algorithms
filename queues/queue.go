package queues

import "fmt"

type Queue struct {
	item_value []int
}

func NewQueue() *Queue {
	return &Queue{
		item_value: make([]int, 0),
	}
}

func (q *Queue) Enqueue(item int) {
	q.item_value = append(q.item_value, item) //used to add items
}

func (q *Queue) Dequeue() int {
	item := q.item_value[0]
	q.item_value = q.item_value[1:] //used to remove items
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.item_value) == 0
}

func (q *Queue) PrintQueue() {
	fmt.Printf("Front ")
	for i := 0; i < len(q.item_value); i++ {
		fmt.Printf("%d ", q.item_value[i])
	}
	fmt.Printf("Back\n")
}
