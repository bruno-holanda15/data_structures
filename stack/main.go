package main

import "fmt"

// const maxStackSize uint = 100
type Stack struct {
	// items [maxStackSize]uint
	items []uint
}

// adiciona por último
func(s *Stack) push(item uint) {
	tamanho := len(s.items)

	newStack := make([]uint, tamanho + 1)
	for index, content := range s.items {
		newStack[index] = content
	}

	newStack[tamanho] = item
	s.items = newStack
}

// sempre remove(pop) o último
func(s *Stack) pop() uint{
	lenght := len(s.items)

	var emptySlice []uint
	var elementPopped uint = s.items[lenght-1]

	s.items = append(emptySlice, s.items[:lenght-1]...)

	return elementPopped
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

	fmt.Println(pilha.pop()) // remove o 8
	pilha.itemsPrint()

	fmt.Println(pilha.pop()) // remove o 7
	pilha.itemsPrint()

	pilha.push(21)  //adiciona o 21
	pilha.itemsPrint()
}