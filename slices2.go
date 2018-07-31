package main

import (
	"fmt"
	"runtime"
)

func printSlice(name string, s []int, comment string, alloc int64) {
	out := fmt.Sprintf("%s = {len=%d, cap=%d, data=%p %v};", name, len(s), cap(s), s, s)
	fmt.Printf("%-25s// %-56s mallocs = +%d\n", comment, out, alloc)
}

var one = []int{1, 2, 3}
var two = []int{4, 5, 6}

func main1() {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)
	k := int64(m.Mallocs)
	a := make([]int, 0, 100)
	runtime.ReadMemStats(&m)
	printSlice("a", a, "a := make([]int, 0, 100)", int64(m.Mallocs)-k)

	runtime.ReadMemStats(&m)
	k = int64(m.Mallocs)
	a = append(a, one...)
	runtime.ReadMemStats(&m)
	printSlice("a", a, "a = append(a, one...)", int64(m.Mallocs)-k)

	runtime.ReadMemStats(&m)
	k = int64(m.Mallocs)
	c := append(a, two...)
	runtime.ReadMemStats(&m)
	printSlice("c", c, "c = append(a, two...)", int64(m.Mallocs)-k)

	runtime.ReadMemStats(&m)
	k = int64(m.Mallocs)
	c[0] = -99
	a[0] = -10
	runtime.ReadMemStats(&m)
	printSlice("a", a, "c[0] = -99, a[0] = -10", int64(m.Mallocs)-k)
	printSlice("c", c, "", int64(m.Mallocs)-k)
}

func main2() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	k := int64(m.Mallocs)

	runtime.ReadMemStats(&m)
	k = int64(m.Mallocs)
	a := make([]int, 0, 0)
	runtime.ReadMemStats(&m)
	printSlice("a", a, "a := make([]int, 0, 0)", int64(m.Mallocs)-k)

	runtime.ReadMemStats(&m)
	k = int64(m.Mallocs)
	a = append(a, one...)
	runtime.ReadMemStats(&m)
	printSlice("a", a, "a = append(a, one...)", int64(m.Mallocs)-k)

	runtime.ReadMemStats(&m)
	k = int64(m.Mallocs)
	c := append(a, two...)
	runtime.ReadMemStats(&m)
	printSlice("c", c, "c = append(a, two...)", int64(m.Mallocs)-k)

	runtime.ReadMemStats(&m)
	k = int64(m.Mallocs)
	c[0] = -99
	a[0] = -10
	runtime.ReadMemStats(&m)
	printSlice("a", a, "c[0] = -99, a[0] = -10", int64(m.Mallocs)-k)
	printSlice("c", c, "", int64(m.Mallocs)-k)
}

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	k := int64(m.Mallocs)
	printSlice("one", one, "one = []int{1, 2, 3}", int64(m.Mallocs)-k)
	printSlice("two", two, "two = []int{4, 5, 6}", int64(m.Mallocs)-k)
	fmt.Println("")

	fmt.Println("main1:")
	main1()
	fmt.Println("main2:")
	main2()
}
