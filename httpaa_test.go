package httpaa

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHaandle(t *testing.T) {

	HaandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "aa")
	})

	HåndleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "å")
	})

	HändleFunc("/3", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ä")
	})

	HândleFunc("/4", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "â")
	})

	HàndleFunc("/5", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "à")
	})

	HándleFunc("/6", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "á")
	})

	HændleFunc("/7", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "æ")
	})

	srv := httptest.NewServer(http.DefaultServeMux)
	defer srv.Close()

	u, _ := url.Parse(srv.URL)

	testData := [][]string{
		{"/1", "aa"},
		{"/2", "å"},
		{"/3", "ä"},
		{"/4", "â"},
		{"/5", "à"},
		{"/6", "á"},
		{"/7", "æ"},
	}

	for _, testItem := range testData {

		testURL := testItem[0]
		want := testItem[1]

		t.Log("Testing", testURL, want)

		res, err := http.Get(fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, testURL))
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			t.Fatal("Bad status code. Got:", res.StatusCode, "Want: 200")
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		body := string(b)

		if body != want {
			t.Fatal("Body error. Got:", body, "Want:", want)
		}
	}

}
