package goku

import (
	"fmt"
	"log"
)

type Queue struct {
	Items []interface{}
	MAX_SIZE int
	ERROR_ON_FULL_ENQUEUE, ERROR_ON_EMPTY_DEQUEUE bool
}


func (q *Queue) Size() int {
	return len(q.Items)
}

func (q *Queue) Enqueue(item interface{}) {
	if (q.Size()>q.MAX_SIZE && q.ERROR_ON_FULL_ENQUEUE){
		panic("The queue is full")
	}
	q.Items = append(q.Items, item)
	for ; q.Size()>q.MAX_SIZE;  {
		q.Dequeue()
		log.Printf("Removing '%v' from queue", q.Dequeue())
	}
	log.Printf("added %v to queue. %d total\n", item, q.Size())
}

func (q *Queue) GetItem(i int) interface{} {
	if q.IsEmpty() {
		return nil
	}else if q.Size()<i {
		return nil
	}
	return q.Items[i]
	
}

// remove the least recently added item
func (q *Queue) Dequeue() interface{} {
	if (q.IsEmpty()) {
		if(q.ERROR_ON_EMPTY_DEQUEUE){
			panic("The queue is empty")
		}
		return nil
	}
	removal := q.Items[0]
	q.Items = q.Items[1:]
	return removal
}

func (q *Queue) IsEmpty() bool {
	if (q.Size() == 0) {
		return true
	}
	return false
}

func (q *Queue) String() string {
	out := "queue contents:"
	for i, item := range q.Items {
		out += fmt.Sprintf("\n[%d]: %3v", i, item)
	}
	return out
}

func New(depth int) *Queue {
	q := new(Queue)
	q.MAX_SIZE=depth
	q.ERROR_ON_EMPTY_DEQUEUE = false
	q.ERROR_ON_FULL_ENQUEUE = true
	return q
}

