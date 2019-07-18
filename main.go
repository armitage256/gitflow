package main

import (
	"fmt"
	"log"
)

func lSum(x, y int) {
	log.Printf("sum %d + %d = %d\n", x, y, (x + y))
}

func main(){
	fmt.Println("Hello World!")
	lSum(10, 11)
}
