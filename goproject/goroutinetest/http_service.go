package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world.")
	})
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timestr := fmt.Sprintf("time : %s", t)
		w.Write([]byte(timestr))
	})
	http.ListenAndServe(":8080", nil)
}
