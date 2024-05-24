package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	listNames := generateRandomNames()
	// fmt.Println(listNames)

	nomeProcurado := "Juliana"
	
	inicio := 0
	fim := len(listNames) - 1

	var numOperacoes uint

	for inicio <= fim {
		numOperacoes++

		meio := (fim + inicio) / 2
		nomeAtual := listNames[meio]
		
		fmt.Println(inicio, fim, meio)
		if nomeAtual == nomeProcurado {
			fmt.Printf("NOME %s ENCONTRADO NA LISTA, COM %d OPERAÇÕES\n", nomeProcurado, numOperacoes)
			break
		}

		if nomeAtual > nomeProcurado {
			fim = meio -1
		}

		if nomeAtual < nomeProcurado {
			inicio = meio + 1
		}

		if inicio == fim {
			fmt.Printf("NOME %s NÃO ENCONTRADO", nomeProcurado)
			break
		}
	}

}	

func generateRandomNames() []string {
	names := "João, Maria, Ana, Pedro, Lucas, Juliana, Carlos, Mariana, Gabriel, Rafael, Aline, Fernanda, Bruno, Amanda, José, Daniel, Rodrigo, Larissa, Paulo, Roberta, Eduardo, Beatriz, Ricardo, Patrícia, Luiz, Eliane, Henrique, Sofia, Gustavo, Renata"
	splited := strings.Split(names, ",")
	var listNames []string

	for _, name := range splited {
		listNames = append(listNames, strings.TrimLeft(name, " "))
	}
	sort.Strings(listNames)

	return listNames
}