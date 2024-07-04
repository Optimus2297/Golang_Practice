package main

import "fmt"

func main(){
	// var colors map[string]string // creates an empty map[]
	// colors := make(map[string]string) // creates empty map[]
	colors := map[string]string{
		"red": "#hkjh",
		"blue": "#kjhkjh",
	}

	colors["white"] = "#dadsdsa"

	printMap(colors)

	delete(colors, "red")
	fmt.Println(colors)
}

func printMap(c map[string]string){
	for color, hex := range c{
		fmt.Println(color, hex)
	}
}