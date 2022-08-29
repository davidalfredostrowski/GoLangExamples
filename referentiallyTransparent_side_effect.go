// referentially transparent but has a 'side effect'

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
	return i + j
}



//3
//3
//3
//3
//3
//3
//
//Program exited.