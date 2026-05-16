package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:3], c)//7, 2, 8
	go sum(s[3:], c)//-9,4 0
	go sum(s[:], c)//slice inteiro
	x:= <-c // receive from c
    y:= <-c
	z:= <-c
	
	fmt.Println("x:",x, "y:",y, "z:",z)
}
