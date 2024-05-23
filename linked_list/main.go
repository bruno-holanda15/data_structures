package main

import "fmt"

// implementar uma lista encada
// implementar os métodos addLast(content, head) e removeLast(head)
// implementar os métodos count(head) altura(elem, head)elem->end e profundidade(elem, head)elem->start
// implementar os mesmos métodos para array (pode usar slice),
// lista duplamente encadeada (*bonus: explorar addFisrt e removeFirst na lista duplamente encadeada)

type Node struct {
	content string
	next    *Node
}

func (n Node) ToString() {
	fmt.Println(n.content)
	if n.next != nil {
		n.next.ToString()
	}
}

func addLast(content string, nodeHead *Node) {
	if nodeHead.next == nil {
		nodeHead.next = &Node{content: content, next: nil}
		return
	}
	addLast(content, nodeHead.next)
}

func removeLast(nodeHead *Node) {
	if nodeHead.next.next == nil {
		nodeHead.next = nil
		return
	}
	removeLast(nodeHead.next)
}

func count(nodeHead *Node) int{
	qty := 1

	for nodeHead.next != nil {
		nodeHead = nodeHead.next
		qty++
	}

	return qty
}

// elem->end
func altura(elem *Node, nodeHead *Node) int{
	var altura int

	for nodeHead != nil {
		if (elem == nodeHead || altura != 0) && (nodeHead.next != nil) {
			// println(nodeHead.content)
			// PRECISARIA INICIAR A CONTAGEM ATÉ O FINAL do for
			altura++
		}

		nodeHead = nodeHead.next
	}

	return altura
}

//começo->elem
func profundidade(elem *Node, nodeHead *Node) int {
	var profundidade int

	for nodeHead != nil {
		if elem == nodeHead {
			// println(nodeHead.content)
			break
		}
		// CASO SEJA DIFERENTE VAI INCREMENTANDO
		profundidade++
		nodeHead = nodeHead.next
	}

	return profundidade
}

func main() {
	n6 := Node{content: "sexto elemento", next: nil}
	n5 := Node{content: "quinto elemento", next: &n6}
	n4 := Node{content: "quarto elemento", next: &n5}
	n3 := Node{content: "terceiro elemento", next: &n4}
	n2 := Node{content: "segundo elemento", next: &n3}
	n1 := Node{content: "primeiro elemento", next: &n2}

	addLast("last item", &n1)
	n1.ToString()

	fmt.Println("######## REMOVE LAST ##########")

	removeLast(&n1)
	n1.ToString()

	fmt.Println("######## COUNT ##########")

	fmt.Printf("Lista com %d Elementos\n", count(&n1))

	// fmt.Println("######## ALTURA ##########")

	// altura(&n3, &n1)

	// fmt.Println("######## PROFUNDIDADE ##########")

	// profundidade(&n2, &n1)
}
