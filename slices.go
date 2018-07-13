package main

import "fmt"

func main1() {
	a := make([]int, 0, 100)
	a = append(a, []int{1, 2, 3}...)
	c := append(a, []int{4, 5, 6}...)

	c[0] = -99
	a[0] = -10
	fmt.Println("a =", a)
	fmt.Println("c =", c)
}
func main2() {
	a := make([]int, 0, 0)
	a = append(a, []int{1, 2, 3}...)
	c := append(a, []int{4, 5, 6}...)

	c[0] = -99
	a[0] = -10
	fmt.Println("a =", a)
	fmt.Println("c =", c)
}

func main() {
	fmt.Println("main1:")
	main1()
	fmt.Println("main2:")
	main2()
}
