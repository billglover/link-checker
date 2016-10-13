package main

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestUrlChecker validates that the urlChecker function can fetch a url
// and return the HTTP status code
func TestUrlChecker(t *testing.T) {

	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"http://www.billglover.me/",
			http.StatusOK,
		},
		{
			"http://www.billglover.me/null",
			http.StatusNotFound,
		},
	}

	t.Log("Given the need to test checking a URL.")
	{
		for _, u := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"", u.url, u.statusCode)
			{
				code, err := getStatusCode(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
				}

				t.Log("\t\tShould be able to make the Get call.", checkMark)

				if code == u.statusCode {
					t.Logf("\t\tShould receive a \"%d\" status. %v", u.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould receive a \"%d\" status. %v %v", u.statusCode, ballotX, code)
				}
			}

		}
	}
}
