package main

import "fmt"

func main() {
	cap := 5
	stack := make([]string, 0, cap)

	stack = append(stack, "A")
	stack = append(stack, "B")
	stack = append(stack, "C")
	stack = append(stack, "D")
	fmt.Println("len\n", len(stack))
	fmt.Println("stack\n", stack[0])

	pr := &stack[3]
	fmt.Println("pr\n", pr)

}
