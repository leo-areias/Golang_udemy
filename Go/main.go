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
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//bs := make([]byte, 99999)
	//make is a built-in function that takes a type of a slice and a second element of
	//number of elements or empty spaces that we want our slice to be inicialized with
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	//Simpler way to do the above code

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
