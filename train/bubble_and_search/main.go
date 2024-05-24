package main

import "fmt"

func main() {
	items := []int{785, 4, 2, 123, 31, 32, 55, 23, 12, 64, 76, 65, 54, 344, 12345, 879, 786, 56, 777, 890, 21, 78, 1}

	bubbleSort(items)
	fmt.Println(items)

	binarySearch(items, 2)
}

func binarySearch(items []int, item int) {
	inicio := 0
	fim := len(items) - 1
	var nomeEncontrado bool

	for inicio <= fim {
		meio := (inicio + fim) / 2
		// fmt.Println(items[inicio], items[meio], items[fim], inicio, fim)

		if item == items[meio] {
			fmt.Printf("O número %d foi encontrado", item)
			nomeEncontrado = true
			break
		}

		if items[meio] > item {
			fim = meio - 1
		}

		if items[meio] < item {
			inicio = meio + 1
		}

		// isso causa o erro de não achar o número na ultima iteração por somar antes da verificação do inicio==fim
		// if inicio == fim {
		// 	fmt.Println(inicio, meio, fim, items[meio])
		// 	fmt.Printf("O número %d NÃO foi encontrado", item)
		// 	break
		// }
	}

	if !nomeEncontrado {
		fmt.Printf("O número %d NÃO foi encontrado", item)
	}
}

func bubbleSort(items []int) {
	for range items {
		fmt.Println(items)
		for i := 0; i < len(items)-1; i++ {
			// fmt.Println(i, items[i], items[i+1])
			if items[i] > items[i+1] {
				swap(items, i, i+1)
			}
			// fmt.Println(items)
		}
	}
}

// troca de posição dos valores
func swap(items []int, i, j int) {
	items[i], items[j] = items[j], items[i]
}
