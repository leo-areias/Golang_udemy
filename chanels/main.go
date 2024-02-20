package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://uniceub.br",
		"http://amazon.com.br",
		"http://stackoverflow.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

// this alternative might be a little bad because you have to wait
// for each response to finish to start another one
// we can fix this by implementing routines
func checkLink(link string) {
	_, err := http.Get(link) // blocking function
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")

}
