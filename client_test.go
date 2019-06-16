package rachio

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientSpec(t *testing.T) {
	ts := serveHTTP(t)
	defer ts.Close()
	previousURL := BaseURL
	BaseURL = ts.URL

	client, err := NewClient("the-rachio-api-key")

	Convey("Should not return an error", t, func() {
		So(err, ShouldBeNil)
	})

	Convey("Should set the HTTP headers", t, func() {
		req, _ := http.NewRequest("GET", "http://foo.com", nil)
		client.setHeaders(req)
		So(req.Header.Get("Authorization"), ShouldEqual, "Bearer the-rachio-api-key")
		So(req.Header.Get("Accept"), ShouldEqual, "application/json")
		So(req.Header.Get("Content-Type"), ShouldEqual, "application/json")
	})

	BaseURL = previousURL
}

func serveHTTP(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		fmt.Println(body)
		switch req.URL.String() {
		case "/person/info":
			checkHeaders(t, req)
			w.WriteHeader(200)
			w.Write([]byte(PersonInfoJSON))
		case "/person/3c59a593-04b8-42df-91db-758f4fe4a97f":
			checkHeaders(t, req)
			w.WriteHeader(200)
			w.Write([]byte(PersonJSON))
		}
	}))
}

func checkHeaders(t *testing.T, req *http.Request) {
	Convey("HTTP headers should be present", t, func() {
		So(req.Header["Accept"][0], ShouldEqual, "application/json")
		So(req.Header["Content-Type"][0], ShouldEqual, "application/json")
	})
}
