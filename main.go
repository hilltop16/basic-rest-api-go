package main

import "net/http"

func main() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:8000", nil)
}

func serveRest(w http.ResponseWriter, r *http.Request){

}
