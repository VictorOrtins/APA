package main

func main(){
	println(fibonacciIterativo(100))
	println(fibonacciRecursivo(100))
}

func fibonacciIterativo(n int) int{
	if n == 0{
		return 0
	}

	var fibonacci [100]int

	fibonacci[0] = 1; fibonacci[1] = 1
	
	for i := 2; i < n; i++{
		fibonacci[i] = fibonacci[i - 1] + fibonacci[i - 2]
	}
	
	return fibonacci[n - 1]

}

func fibonacciRecursivo(n int) int{
	if n == 0{
		return 0
	}

	if n == 1{
		return 1
	}

	return fibonacciRecursivo(n - 1) + fibonacciRecursivo(n - 2)
}