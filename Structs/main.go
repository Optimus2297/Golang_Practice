package main

import "fmt"

type address struct{
	city string
	state string
}

type person struct {
	firstname string
	lastname  string
	contact address
}

func main() {
	// bhaskar := person{firstname: "Bhaskar", lastname: "Haranale"}
	// fmt.Println(bhaskar)
	// fmt.Printf("%+v", bhaskar)
	bhaskar := person{
				firstname: "Bhaskar",
				lastname: "Haranale",
				contact: address{
					city: "sangli",
					state: "Maharashtra",
				},
	}
	// bhaskarPointer := &bhaskar
	// bhaskarPointer.updatename("Bhagyaraj")
	bhaskar.updatename("Bhausaheb")
	bhaskar.print()
}
func (pointerToPerson *person) updatename( name string){
	(*pointerToPerson).firstname = name
}

func (p person) print(){
	fmt.Printf("%+v", p)
}