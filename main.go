package main

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	println("Welcome to the Go Testing App")
	result := Calculate(2)
	println(result)

}
