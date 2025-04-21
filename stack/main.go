package main

import "fmt"

func main() {
  stack := Stack{Size: 3}
	stack.Push("papa-capim")
	stack.Push("dos")
	stack.Push("meus")
	stack.Push("sonhos")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
