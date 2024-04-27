package main

import "fmt"

func main() {

	//1st type
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	//2nd type
	arr := [3]string{"apple", "banana", "orange"}
	for idx, val := range arr {
		fmt.Printf("%v\t %v\n", idx, val)
	}

	//while loop
	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}
}
