// 'not' referentially transparent w/ 'side effect'
// "not" referentially transparent

package main

import "fmt"

var globalvariable int = 1

func main() {
	fmt.Println(a(1, 2))
	var ans = a(1, 2)
	fmt.Println(ans)
	fmt.Println(a(1, 2))
	fmt.Println(3)
	fmt.Println(a(1, 2))
	fmt.Println(3)
}

func a(i, j int) int {
	globalvariable = globalvariable + 1
	return globalvariable + i + j
}
