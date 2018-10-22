package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("=== starting logger on localhost:3000")
	http.ListenAndServe("localhost:3000",
		http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
			fmt.Printf("%s %s\n", r.Method, r.URL)
		}),
	)
}
