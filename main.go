package main

import "fmt"

// func main() {
// 	// var a = "initial"
// 	// fmt.Println(a)

// 	// var b, c int = 1, 2
// 	// fmt.Println(b, c)

// 	// var d = true
// 	// fmt.Println(d)

// 	// var e int
// 	// fmt.Println(e)

// 	// f := "short"
// 	// fmt.Println(f)

// 	// countries := [2]string{"peru", "argentina"}
// 	// countries[1] = "uruguay"
// 	// fmt.Println(countries)

// 	// var a int = 0x50
// 	// var b float64 = 7.4
// 	// var c float64

// 	// c = float64(a) + float64(b)

// 	// fmt.Println(c)
// 	var sauces [3]string
// 	sauces[0] = "ketchup"
// 	sauces[1] = "bbq"
// 	sauces[2] = "mayo"
// 	fmt.Println(sauces)
// }

// func main() {
// 	f(4)
// }

// func f(x int) {
// 	fmt.Printf(""f(%d)\n"", x+0/x)
// 	defer fmt.Printf(""defer %d\n"", x)
// 	f(x-1)
// }

// func factorial(x uint) uint {
// 	if x == 0 {
// 		return 1
// 	}
// 	return x * factorial(x-1)
// }

// func main() {
// 	fmt.Println(factorial(5))
// }

const MAX int = 3

func main() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* assign the address of integer. */
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of test a[%d] = %d\n", i, *ptr[i])
	}
}
