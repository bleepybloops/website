package main

import (
	"fmt"
	"http"
	// "io"
	// "os"
	// "path"
	// "strings"
)

func main() {
	
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, 世界")
		fmt.Fprintf(w, "Hello, 世界")
	})

	http.ListenAndServe(":8080", nil)

}