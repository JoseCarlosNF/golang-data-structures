package main

import "fmt"

// List uma lista encadeada, formada por Nodes.
type List struct {
	Head *Node
	Tail *Node
}

// Node, indica o conteúdo do nó e o apontamento para o próximo.
type Node struct {
	Value People
	Next  *Node
}

// Append adiciona um novo elemento a lista encadeada.
func (l *List) Append(people People) {
	node := &Node{Value: people}
	if l.Head == nil {
		l.Head = node
	}
	if l.Tail != nil {
		l.Tail.Next = node
	}
	l.Tail = node
}

// Display imprime o estado atual da lista encadaada.
func (l *List) Display() {
	curr := l.Head
	for curr != nil {
		fmt.Println(curr.Value.Name)
		curr = curr.Next
	}
}

// Search retorna um item do tipo People.
func (l *List) Search(name string) People {
	curr := l.Head
	for curr != nil {
		if curr.Value.Name == name {
			return curr.Value
		}
		curr = curr.Next
	}
	return People{}
}

// Delete retira um item da lista, com base no valor de referência.
func (l *List) Delete(name string) {
	if l.Head.Value.Name == name {
		l.Head = l.Head.Next
		return
	}
	prev := l.Head
	curr := l.Head.Next
	for curr != nil {
		if curr.Value.Name == name {
			prev.Next = curr.Next
			break
		}
		prev = curr
		curr = curr.Next
	}
}
