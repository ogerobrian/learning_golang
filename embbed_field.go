package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	specialty string
}

func main() {
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His specialty is ", mark.specialty)

	mark.specialty = "AI"
	fmt.Println("Mark changed his specialty")
	fmt.Println("His specialty is ", mark.specialty)

	fmt.Println("Mark become old. He is not an athlete anymore")
	mark.age = 46
	mark.weight += 60
	fmt.Println("His age is", mark.age)
	fmt.Println("His weight is", mark.weight)
}
