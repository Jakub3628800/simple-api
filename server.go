package main

import "net/http"
import "encoding/json"

type Resource struct {
	name        string
	otherattr   string
	maybeintatr int
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	var resource Resource
	resource.name = "test"
	resource.otherattr = "test2"
	resource.maybeintatr = 8
	jsonBytes, err := json.Marshal(resource)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write(jsonBytes)
}

func main() {
	http.HandleFunc("/resource", resourceHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
