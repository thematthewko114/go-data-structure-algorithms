package queues

import "fmt"

type Queue struct {
	item_value []string
}

func NewQueue() *Queue {
	return &Queue{
		item_value: make([]string, 0),
	}
}

func (q *Queue) Enqueue(item string) {
	q.item_value = append(q.item_value, item) //used to add items
}

func (q *Queue) Dequeue() string {
	item := q.item_value[0]
	q.item_value = q.item_value[1:] //used to remove items
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.item_value) == 0
}

func (q *Queue) PrintQueue() {
	fmt.Printf("Front |")
	for i := 0; i < len(q.item_value); i++ {
		fmt.Printf(" %s ", q.item_value[i])
	}
	fmt.Printf("| Back\n")
}
