int x = 0;

package main

import "fmt"

var globalvariable int = 1

func main() {
	fmt.Println(referentiallyTransparent(1))
	fmt.Println(referentiallyTransparent(1))
	
	fmt.Println(hasNoSideEffect(1))
	fmt.Println(hasNoSideEffect(1))
	
}



referentiallyTransparent(int y) int
{
    return y + 1;
}


// not referentially transparent
hasNoSideEffects(int y) int
{
    return globalvariable + y;
}