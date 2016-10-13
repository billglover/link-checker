package main

import "testing"

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestUrlChecker validates that the urlChecker function can fetch a url
// and return the HTTP status code
func TestUrlChecker(t *testing.T) {
	url := "http://www.billglover.me/"
	statusCode := 200

	t.Log("Given the need to test checking a URL.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			code, err := getStatusCode(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}

			t.Log("\t\tShould be able to make the Get call.", checkMark)

			if code == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, code)
			}
		}
	}
}
