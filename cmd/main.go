package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("=== starting logger on localhost:3000")
	http.ListenAndServe("localhost:3000",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("Request: %s %s\n", r.Method, r.URL)
			fmt.Printf("Header: %s\n", spew.Sdump(r.Header))

			bodyData, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			fmt.Printf("Body: %s\n", bodyData)
		}),
	)
}
