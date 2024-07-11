package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

// 	bs := make([]byte, 999) // Reader interface has a function read which accepts byte slice and updates the response body in the provided byte slice.

// 	resp.Body.Read(bs) // the read function inreader interface is called with byte slice

// 	fmt.Println(string(bs)) // we get the response in byte slice

	// io.Copy(os.Stdout, resp.Body) 	// io.Copy function takes two input arguments any  value of type that implements the writer interface and any value of type that implements the reader interface. 
									//So because of this os.stdout is value of type that implements Writer interface and resp.Body is value of type that implements Reader interface.

	lw := logWriter{}
							// Custom implementation of Write interface which we are using as an input argument fro io.Copy()
	io.Copy(lw, resp.Body)
}

// custom implementation of Write interface which recieves logWriter type struct
func (logWriter) Write(bs []byte) (int, error){
	fmt.Println((string(bs)))
	fmt.Println("logged this much bytes : ", len(bs))
	return len(bs), nil
}