package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
)

// this function is copied from 
// https://medium.com/doing-things-right/pretty-printing-http-requests-in-golang-a918d5aaa000
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string // Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url) // Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host)) // Loop through headers
	for name, headers := range r.Header {
	  name = strings.ToLower(name)
	  for _, h := range headers {
		request = append(request, fmt.Sprintf("%v: %v", name, h))
	  }
	}
	
	// If this is a POST, add post data
	if r.Method == "POST" || r.Method == "PUT" {
	   r.ParseForm()
	   request = append(request, "\n")
	   request = append(request, r.Form.Encode())
	}   // Return the request as a string
	 return strings.Join(request, "\n")
}


func responser (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got " + r.Method + " in path " + r.URL.Path)
	fmt.Fprintf(w, formatRequest(r))
}

func main() {
	fmt.Println("responser started in port 8080")
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.HandlerFunc(responser))
	log.Fatal(http.ListenAndServe(":8080", router))
}