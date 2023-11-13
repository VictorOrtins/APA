package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	nome string
	valor float64
	peso  float64
	razao float64
}

type Mochila struct {
	qtdItems   int
	pesoMaximo float64
	pesoAtual float64
	itens []Item
}

func createMochila(pesoMaximo float64) *Mochila {
	itens := []Item{}
	mochila := Mochila{0, pesoMaximo, 0,itens}
	return &mochila
}

func (m* Mochila) printMochila(){
	fmt.Printf("Mochila:\n - qtdItens: %d\n - pesoMaximo: %f\n - pesoAtual: %f\n", m.qtdItems, m.pesoMaximo, m.pesoAtual)
	for _, elemento := range m.itens{
		elemento.prinItem()
	}	
}

func (i* Item) prinItem(){
	fmt.Printf("Item:\n nome: %s\n valor: %f\n peso: %f\n razao: %f\n", i.nome, i.valor, i.peso, i.razao)
}

func (m* Mochila) addItemMochila(item Item) error{
	if m.pesoAtual + item.peso > m.pesoMaximo{
		return errors.New("Não é possível adicionar mais itens")
	}

	m.itens = append(m.itens, item)
	m.qtdItems++
	m.pesoAtual += item.peso

	return nil
}

func registrarItens(nomeTxt string) []Item{
	arquivo, err := os.Open("itens.txt")
	if err != nil{
		fmt.Println("Erro ao abrir o arquivo")
		return nil
	}

	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)

	linhas := []string{}

	for scanner.Scan(){
		linha := scanner.Text()
		linhas = append(linhas, linha)
	}

	itens := []Item{}

	for _, elemento := range linhas{
		pos_virgula := strings.Index(elemento, ",")
		nome := elemento[0:pos_virgula]

		elemento = elemento[pos_virgula + 1:]

		pos_virgula = strings.Index(elemento, ",")
		valor,err := strconv.ParseFloat(elemento[0:pos_virgula], 64)
		if err != nil{
			fmt.Println("Erro no arquivo")
			return nil
		}

		elemento = elemento[pos_virgula + 1:]

		peso, err := strconv.ParseFloat(elemento, 64)
		if err != nil{
			fmt.Println("Erro no arquivo")
			return nil
		}

		item := Item{nome, valor, peso, valor/peso}
		itens = append(itens, item)
	}

	return itens
}

func main() {

	itens_disponiveis := registrarItens("itens.txt")
	if itens_disponiveis == nil{
		fmt.Println("Erro no arquivo!!!")
		return
	}
	mochila := createMochila(9)

	for {
		if len(itens_disponiveis) == 0{
			break
		}

		pos_maior_razao := 0
		for posicao, elemento := range itens_disponiveis{
			if elemento.razao > itens_disponiveis[pos_maior_razao].razao{
				pos_maior_razao = posicao
			}
		}

		err := mochila.addItemMochila(itens_disponiveis[pos_maior_razao])
		if err != nil{
			break
		}

		for i := pos_maior_razao; i < len(itens_disponiveis) - 2; i++{
			itens_disponiveis[i] = itens_disponiveis[i+1]
		}

		itens_disponiveis = itens_disponiveis[:len(itens_disponiveis) - 1]
	}

	mochila.printMochila()
}