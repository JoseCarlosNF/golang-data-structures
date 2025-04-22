package main

import "fmt"

// People, tipo simples para exemplificação de uma pessoa.
type People struct {
	Name string
	Age  int
}

func main() {
	p1 := People{"Joaquim", 15}
	p2 := People{"Adagoberto", 25}
	p3 := People{"Jurema", 65}

	linked_list := List{}

	linked_list.Append(p1)
	linked_list.Append(p2)
	linked_list.Append(p3)

	fmt.Println("--- Impressão ---")
	linked_list.Display()

	fmt.Println("\n--- Busca ---")
	fmt.Println(linked_list.Search("Jurema"))

	fmt.Println("\n--- Remoção ---")
	linked_list.Delete("Adagoberto")

	fmt.Println("\n--- Impressão pós remoção ---")
	linked_list.Display()
}
