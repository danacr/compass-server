package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Compass struct {
	Degrees int
}

var D Compass

func hello(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		js, err := json.Marshal(Compass{Degrees: D.Degrees})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(js)
		if err != nil {
			fmt.Printf("Error %q", err)
		}
		return
	case "POST":
		err := json.NewDecoder(r.Body).Decode(&D)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else {
			http.StatusText(200)
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting Compass server\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
