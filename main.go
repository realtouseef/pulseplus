package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/realtouseef/pulseplus/helpers"
)

var wg sync.WaitGroup

func main() {

	urls := []string{
		"https://github.com",
		"https://www.linkedin.com",
		"https://bloomycorners.com",
	}

	for _, url := range urls {
		wg.Add(1)
		go ping(url)
	}

	wg.Wait()
}

func ping(url string) {
	response, err := http.Get(url)
	helpers.HandleError(err)

	fmt.Printf("status code for %v is %v\n", url, response.StatusCode)
	wg.Done()

	defer response.Body.Close()
}
