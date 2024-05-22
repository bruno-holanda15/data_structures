package main

import "fmt"

type Stack struct {
	items []uint
}

// adiciona por último
func(s *Stack) add(item uint) {
	s.items = append(s.items, item)
}

// sempre remove o último
func(s *Stack) remove() {
	lenght := len(s.items)

	var emptySlice []uint
	s.items = append(emptySlice, s.items[:lenght-1]...)
}

func(s *Stack) itemsPrint() {
	fmt.Println(s.items)
}

// exibe o último elemento da pilha/ próximo a ser removido
func(s *Stack) peek() {
	lenght := len(s.items)

	fmt.Println(s.items[lenght-1])
}


func main() {
	pilha := Stack{items: []uint{1,3,4,5,6,7,8}}
	pilha.itemsPrint()
	pilha.peek()

	pilha.remove() // remove o 8
	pilha.itemsPrint()
	pilha.peek()


	pilha.remove() // remove o 7
	pilha.itemsPrint()
	pilha.peek()


	pilha.add(21)  //adiciona o 21
	pilha.itemsPrint()
	pilha.peek()

}