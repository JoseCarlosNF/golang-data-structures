package main

// Queue implementa o comportamento de um fila FIFO.
type Queue struct {
	Head *Node
	Tail *Node
}

// Node representa cada um dos items da fila.
type Node struct {
	Value string
	Next  *Node
}

// Enqueue adiciona um item na calda da fila.
func (q *Queue) Enqueue(name string) {
	node := &Node{Value: name}

	if q.Head == nil {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		q.Tail = node
	}
}

// Dequeue remove um item da cabeça da fila.
func (q *Queue) Dequeue() string {
	if q.Head == nil {
		q.Tail = nil
		return ""
	}
	node := q.Head
	q.Head = q.Head.Next
	return node.Value
}
