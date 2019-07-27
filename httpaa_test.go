package httpaa

import (
	"fmt"
	"net/http"
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

}
