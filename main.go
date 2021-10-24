package main

//package main

import (
"fmt"
"log"
"net/http"
)
//func
func main() {
	c := 0
	fmt.Println("started serving ...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		c += 1
		fmt.Fprintf(w, "pong")
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "total pings: %v", c)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
