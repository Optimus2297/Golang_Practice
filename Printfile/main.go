package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// data, err := os.ReadFile(os.Args[1])
	// if err != nil {
	// 	fmt.Println("Error :", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(data))

	f, err := os.Open("demo.txt") // f is of type pointer to file
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	fmt.Println(f, *f)
	io.Copy(os.Stdout, f)
}