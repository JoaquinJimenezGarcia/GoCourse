package main

import "fmt"

func main() {
	var ns []int

	for i := 0; i <= 10; i++ {
		ns = append(ns, i)
	}
	
	for _,value := range ns {
		if value % 2 == 0 {
			fmt.Println(value, "It's even")
		} else {
			fmt.Println(value, "It's odd")
		}
	}
}