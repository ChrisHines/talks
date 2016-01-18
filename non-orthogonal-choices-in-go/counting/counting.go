package main

import "fmt"

func main() {
	fmt.Println("How do I loop the,")
	fmt.Println("let me count the ways.")

	forloop(9)
	fmt.Println()

	gotoloop(9)
	fmt.Println()

	recurse(9)
	fmt.Println()
}

func forloop(n int) {
	for i := 1; i <= n; i++ {
		fmt.Print(i, " ")
	}
}

func gotoloop(n int) {
	i := 1
loop:
	if i <= n {
		fmt.Print(i, " ")
		i++
		goto loop
	}
}

func recurse(n int) {
	if n > 1 {
		recurse(n - 1)
	}
	fmt.Print(n, " ")
}
