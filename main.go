package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

// Receive a list of URLs on STDIN, query them to determine
// the HTTP status code. Log the URL and the associate HTTP
// status code to STDOUT. Note: no validation is done on
// the URLs prior to making the request. The order of the
// URLs in the output is not guaranteed to match that the
// input.
func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		u := scanner.Text()
		wg.Add(1)
		go checkURL(u)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	wg.Wait()

}

// Take a URL as a string, query the URL and print the
// status code. Note that this assumes the string represents
// a valid URL.
func checkURL(u string) {

	defer wg.Done()

	r, err := http.Get(u)
	if err != nil {
		fmt.Printf("%s,%s\n", u, err.Error())
		return
	}
	fmt.Printf("%s,%d\n", u, r.StatusCode)
}
