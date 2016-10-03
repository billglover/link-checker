package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

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
		checkURL(u)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}

// Take a URL as a string, query the URL and print the
// status code. Note that this assumes the string represents
// a valid URL.
func checkURL(u string) {
	r, err := http.Get(u)
	if err != nil {
		fmt.Printf("%s,%s\n", u, err.Error())
		return
	}
	fmt.Printf("%s,%d\n", u, r.StatusCode)
}
