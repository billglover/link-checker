package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
)

const conc = 10 // concurrency factor

func main() {
	checkUrls(bufio.NewScanner(os.Stdin))
}

// Scan through a bufio.Scanner for a list of URLs, and 
// query them to determine the HTTP status code. Log the 
// URL and the associate HTTP status code to STDOUT. 
// Note: no validation is done on the URLs prior to 
// making the request. The order of the URLs in the 
// output is not guaranteed to match that the input.
func checkUrls(s *bufio.Scanner) {
	var wg sync.WaitGroup

	c := make(chan string)

	for i := 0; i < conc; i++ {
		wg.Add(1)
		go urlChecker(c, &wg)
	}

	for s.Scan() {
		u := s.Text()
		c <- u
	}

	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	close(c)

	wg.Wait()
}

// Read URLs off a channel and perform an HTTP GET to determine the
// HTTP status code. When there is nothing left to read off the
// channel signal to the WaitGroup that processing is complete and
// return. Note: this assumes that all strings passed on the
// channel are valid URLs.
func urlChecker(c chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	for u := range c {
		c := getStatusCode(u)
		fmt.Printf("%s, %d\n", u, c)
	}
}

// Perform an HTTP GET on a URL and return the status code.
// Note that this wraps errors and doesn't pass them up
// the stack.
func getStatusCode(u string) int {
	r, err := http.Get(u)

	if err != nil {
		fmt.Printf("%s, %s\n", u, err.Error())
		return 0
	}

	return r.StatusCode
}
