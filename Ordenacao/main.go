package main

import (
	"fmt"
	"math/rand"
	"time"
)

type heap[T int] struct{
	heap_array []T
	heap_len int
}

func (h *heap[T])create_heap_array(size int){
	h.heap_array = make([]T, size)
}

func (h *heap[T])set_heap_array(array []T){
	h.heap_array = array
}

func (h *heap[T])right_son(node_index int) int{
	return 2*node_index + 2
}

func (h *heap[T])left_son(node_index int) int{
	return 2*node_index + 1
}

func (h *heap[T])max_heapify(node_index int){

	if node_index >= h.heap_len{
		return
	}

	maior := node_index

	esquerdo := h.left_son(node_index)
	direito := h.right_son(node_index)

	if esquerdo < h.heap_len && h.heap_array[esquerdo] > h.heap_array[maior]{
		maior = esquerdo
	}

	if direito < h.heap_len && h.heap_array[direito] > h.heap_array[maior]{
		maior = direito
	}

	if maior != node_index{
		swap(&h.heap_array[node_index], &h.heap_array[maior])
		h.max_heapify(maior)
	}
}

func (h *heap[T])heap_sort(array []T){
	h.set_heap_array(array)
	h.heap_len = len(array)

	for i := h.heap_len/2 - 1; i >= 0; i--{
		h.max_heapify(i)
	}

	for i := h.heap_len - 1; i >= 0; i--{
		swap(&array[0], &array[i])
		h.heap_len--
		h.max_heapify(0)
	}
}

func main() {
	const maxNumero = 100
	const tamList = 10

	// list := createList(tamList, maxNumero)
	// fmt.Println(list2)
	// selectionSort(list)
	// fmt.Println(list2)


	// list2 := createList(tamList, maxNumero)
	// fmt.Println(list2)
	// insertionSort(list2)
	// fmt.Println(list2)

	// list3 := []int{6,5,4,3,2,1}
	// fmt.Println(list3)
	// bubbleSort(list3)
	// fmt.Println(list3)

	// list4 := createList(tamList, maxNumero)
	// fmt.Println(list4)
	// posicao, str := buscaLinear(list4, 1)
	// fmt.Println(str, posicao)

	// posicao, str = buscaBinaria(list4, 1, 0, len(list4) - 1)
	// fmt.Println(str, posicao)

	// list5 := []int{1, 3, 5, 7, 2, 4, 6, 8, 10, 12, 14}
	// fmt.Println(list5)
	// merge(list5, 0, 3, 4, 10)
	// fmt.Println(list5)

	// list6 := createList(tamList, maxNumero)
	// fmt.Println(list6)
	// mergeSort(list6, 0, len(list6) - 1)
	// fmt.Println(list6)

	// list7 := []int{7,8,1,2,90,4,64,32}
	// fmt.Println(list7)
	// partition(list7, 0, 7)
	// fmt.Println(list7)

	// list8 := createList(tamList, maxNumero)
	// fmt.Println(list8)
	// quickSort(list8, 0, len(list8) - 1)
	// fmt.Println(list8)

	// list9 := createList(tamList, maxNumero)
	// fmt.Println(list9)
	// countingSort(list9)
	// fmt.Println(list9)

	heap := heap[int]{}

	list10 := createList(tamList, maxNumero)
	fmt.Println(list10)
	heap.heap_sort(list10)
	fmt.Println(list10)

}

func selectionSort[T int](list []T) (string){
	start := time.Now()

	for i := 0; i < len(list) - 1; i++{
		atual := &list[i]

		menor := &list[i + 1]

		for j := i + 2; j < len(list); j++{
			if list[j] < *menor{
				menor = &list[j]
			}
		}

		if *menor < *atual{
			swap(atual, menor)
		} 
	}
	percorrido := time.Since(start)
	return fmt.Sprintln("Tempo de Ordenação:", percorrido.String())
}

func insertionSort[T int](list []T) (string){

	start := time.Now()

	for i := 1; i < len(list); i++{
		j := i - 1
		atual := list[i]

		for j >= 0 && list[j] > atual{
			list[j + 1] = list[j]
			j--
		}

		list[j + 1] = atual
	}

	percorrido := time.Since(start)
	return fmt.Sprintln("Tempo de Ordenação:", percorrido.String())
}

func bubbleSort[T int](list []T) (string){
	start := time.Now()

	houveTroca := true

	contador := 0
	for houveTroca{
		houveTroca = false
		for i := 0; i < len(list) - 1; i++{
			contador++
			if list[i] > list[i + 1]{
				swap(&list[i], &list[i + 1])
				houveTroca = true
			}
		}
	}

	fmt.Println(contador)

	percorrido := time.Since(start)
	return fmt.Sprintln("Tempo de Ordenação:", percorrido.String())
}

func buscaLinear[T int](list []T, buscar T) (int, string){
	start := time.Now()

	i := 0
	for ; i < len(list); i++{
		if list[i] == buscar{
			break
		}
	}

	percorrido := time.Since(start)
	return i, fmt.Sprintln("Tempo de Execução:", percorrido.String())
}

func buscaBinaria[T int](list []T, buscar T, inicio, fim int) (int, string){
	if inicio == fim{
		return -1, "Falhou"
	}

	meio := (inicio + fim)/2
	fmt.Println(meio)
	if buscar == list[meio]{
		return meio, "Deu certo"
	}

	if buscar < list[meio]{
		return buscaBinaria(list, buscar, inicio, meio - 1)
	}

	return buscaBinaria(list, buscar, meio + 1, fim)
}

func mergeSort[T int](list []T, inicio, fim int){
	if inicio < fim{
		meio := (inicio + fim)/2
		mergeSort(list, inicio, meio)
		mergeSort(list, meio + 1, fim)
		merge(list, inicio, meio, meio + 1, fim)
	}
}

func merge[T int](list []T, inicio1, fim1, inicio2, fim2 int){
	// lenParte1 := fim1 - inicio1 + 1
	// lenParte2 := fim2 - inicio2 + 1

	novoArray := make([]T, len(list))

	i := inicio1
	j := inicio2

	contador := 0

	for i <= fim1 && j <= fim2{
		if list[i] <= list[j]{
			novoArray[contador] = list[i]
			i++
		}else{
			novoArray[contador] = list[j]
			j++
		}
		contador++
	}

	if i <= fim1{
		for i <= fim1{
			novoArray[contador] = list[i]
			i++
			contador++
		}
	}else if j <= fim2{
		for j <= fim2{
			novoArray[contador] = list[j]
			j++
			contador++
		}
	}

	contador = 0
	for i := inicio1; i <= fim2; i++{
		list[i] = novoArray[contador]
		contador++
	}
}

func quickSort[T int](list []T, inicio, fim int){
	if inicio < fim{
		meio := partition(list, inicio, fim)
		quickSort(list, inicio, meio - 1)
		quickSort(list, meio + 1, fim)		
	}
}

func partition[T int](list []T, inicio, fim int) int{
	na_frente := inicio + 1
	for i := inicio; i <= fim; i++{
		if list[inicio] > list[i]{
			swap(&list[na_frente],&list[i])
			na_frente++
		}
	}

	swap(&list[inicio], &list[na_frente - 1])
	return na_frente - 1
}

func countingSort[T int](list []T){
	newArray := make([]T, len(list))
	maior := list[0]

	for i := 1; i < len(list); i++{
		if list[i] > maior{
			maior = list[i]
		}
	}

	arrayDistribuicao := make([]T, maior + 1)

	for i := 0; i < len(list); i++{
		arrayDistribuicao[list[i]]++
	}

	for i := 1; i < len(arrayDistribuicao); i++{
		arrayDistribuicao[i] = arrayDistribuicao[i] + arrayDistribuicao[i - 1]
	}

	for i := len(list) - 1; i >= 0; i--{
		newArray[arrayDistribuicao[list[i]] - 1] = list[i]
	}

	for i := 0; i < len(list); i++{
		list[i] = newArray[i]
	}
}

func createList(tamanho, maxNumero int) []int{
	lista := make([]int, tamanho)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < tamanho; i++{
		lista[i] = rand.Intn(maxNumero)
	}

	return lista
}

func swap[T comparable](a* T, b* T) {
	temp := *a
	*a = *b
	*b = temp
}