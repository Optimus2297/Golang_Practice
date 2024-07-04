// program to find and print odd and even numbers froma sllice
package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("The Number %v is Even\n", num)
		} else {
			fmt.Printf("The number %v is ODD\n", num)
		}
	}
}