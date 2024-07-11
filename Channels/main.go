package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://youtube.com",
		"http://instagram.com",
		"http://flipkart.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i:=0;i< len(links);i++{
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkLink(<-c, c)
	// }

	// for l := range c{
	// 	go checkLink(l, c) // Repeating routines
	// }

	for l := range c{
		go func( link string){
			time.Sleep(5 * time.Second)   // using some pause before making repated routine calls. Also one important thing is not to share main routine variables in child routine. Always pass the variable from main routine to child routine. 
											//otherwise the variable value in main routine may change.
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " is not up!")
		// c <- "Nope its not up!"
		c <- link
		return
	}

	fmt.Println(link , " is up!")
	// c <- "Yep its up!"
	c <- link
}