package main

import "fmt"

type bot interface {
	getGreetings() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)

}

func printGreetings(b bot) {
	//Similar function for both english and spnish bots but with different input paramater types
	fmt.Println(b.getGreetings())
}

func (englishBot) getGreetings() string {
	//complex function seperate for english and spanish boits
	return "Hi there"
}

func (spanishBot) getGreetings() string { // if the reciver is not use we dont need to set the value only type defined will work.
	return "Hola"
}