
"""
	a := make([]int, 0, 100)
	a = append(a, []int{1, 2, 3}...)
	c := append(a, []int{4, 5, 6}...)

	c[0] = -99
	a[0] = -10
	fmt.Println("a =", a)
	fmt.Println("c =", c)
"""
def main1():
    a = []
    a.extend([1, 2, 3])
    c = a # Reference
    c.extend([4, 5, 6])

    a[0] = -99
    c[0] = -10

    print("a =", a)
    print("c =", c)

def main2():
    a = []
    a.extend([1, 2, 3])
    c = a[:] # Copy
    c.extend([4, 5, 6])

    a[0] = -99
    c[0] = -10

    print("a =", a)
    print("c =", c)

print("main1:")
main1()
print("main2:")
main2()
