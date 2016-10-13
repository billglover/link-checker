package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// mockServer returns a pointer to a server to handle the HTTP Get call
func mockServer(c int) *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(c)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

// TestUrlChecker validates that the urlChecker function can fetch a url
// and return the HTTP status code
func TestUrlChecker(t *testing.T) {

	serverOk := mockServer(http.StatusOK)
	serverNotFound := mockServer(http.StatusNotFound)
	defer serverOk.Close()
	defer serverNotFound.Close()

	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			serverOk.URL,
			http.StatusOK,
		},
		{
			serverNotFound.URL,
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

// BenchmarkCheckUrls provides performance numbers for the checkUrls function
func BenchmarkCheckUrls(b *testing.B) {

	s := mockServer(http.StatusOK)
	defer s.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = getStatusCode(s.URL)
	}

}
