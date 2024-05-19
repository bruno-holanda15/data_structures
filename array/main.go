package main

import "fmt"

// implementar os métodos addLast(content, head) e removeLast(head)
// implementar os métodos count(head) altura(elem, head)elem->end e profundidade(elem, head)elem->start
// implementar os mesmos métodos para array (pode usar slice),

type arrayP struct {
	elements []uint
}

func main() {
	// array uint -> same type elements
	// size 6 -> fixed
	array := []uint{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(array[3:]) // retorna elementos desde o 3 index [3 4 5 6 7 8]
	fmt.Println(array[:3]) // retorna elementos até o 3 index [1 2]
	// fmt.Printf("%T\n",array[:8]) // retorna elementos até o 2 index

	arrayS := arrayP{elements: array}
	// fmt.Println(arrayS.count())

	// arrayS.removeLast()
	// arrayS.toString()

	// arrayS.addLast(10)
	arrayS.toString()

	arrayS.removeAt(2)
	arrayS.toString()

	arrayS.addAt(2, 21)
	arrayS.toString()
}

func(a *arrayP) toString() {
	fmt.Println(a.elements)
}

func(a *arrayP) count() uint {
	var size uint
	for range a.elements {
		size++
	}

	return size
}

func(a *arrayP) removeLast() {
	size := a.count()
	a.elements = a.elements[:size-1]
}

func(a *arrayP) removeAt(position uint) {
	arrayLeft := a.elements[:position-1]
	arrayRight := a.elements[position:]
	a.elements = append(arrayLeft, arrayRight...)
}

// primeira tentativa
// func(a *arrayP) addAt(position, element uint) {
// 	arrayLeft := a.elements[:position]
// 	arrayRight := a.elements[position:]

// 	fmt.Println(arrayLeft, arrayRight)

// 	arrayLeft = append(arrayLeft, element)

// 	fmt.Println(arrayLeft, arrayRight)

// 	a.elements = append(arrayLeft, arrayRight...)
// }

// usando copy
func(a *arrayP) addAt(position, element uint) {
	a.elements = append(a.elements, 0)
	copy(a.elements[position+1:], a.elements[position:])

	a.elements[position] = element
}

func(a *arrayP) addLast(element uint) {
	a.elements = append(a.elements, element)
}