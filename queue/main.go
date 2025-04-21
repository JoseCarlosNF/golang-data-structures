package main

import "fmt"

func main() {
	queue := Queue{}

	queue.Enqueue("Adagoberto")
	queue.Enqueue("Jurema")
	queue.Enqueue("Adalgiza")
	queue.Enqueue("Jumencio")

  fmt.Println(queue.Dequeue())
  fmt.Println(queue.Dequeue())
  fmt.Println(queue.Dequeue())
  fmt.Println(queue.Dequeue())
}
