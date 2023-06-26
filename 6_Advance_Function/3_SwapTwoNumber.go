package main

import "fmt"

func swapByPointer(a, b *int) {
	*b, *a = *a, *b
}

func main(){
	a := 10
	b := 20

	swapByPointer(&a, &b)
	fmt.Println(a,b)
}



