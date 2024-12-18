package util

import "container/heap"

type prioritisedElement[T any] struct {
	priority int
	element T
}

type PriorityQueue[T any] []prioritisedElement[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(prioritisedElement[T])
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue[T]) IsEmpty() bool {
	return pq.Len() == 0
}

func (pq *PriorityQueue[T]) Insert(element T, priority int) {
	heap.Push(pq, prioritisedElement[T]{priority, element})
}

func (pq *PriorityQueue[T]) PullPriority() (T, int) {
	item := heap.Pop(pq)
	entry := item.(prioritisedElement[T])
	return entry.element, entry.priority
}

func (pq *PriorityQueue[T]) PullValue() T {
	value, _ := pq.PullPriority()
	return value
}

func PriorityQueueOf[T any](item T, priority int) PriorityQueue[T] {
	pq := make(PriorityQueue[T], 0)
	pq.Insert(item, priority)
	return pq
}

