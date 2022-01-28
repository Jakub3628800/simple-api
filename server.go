package main

import "net/http"
import "encoding/json"
import "log"

type resource struct {
		Name        string `json:"name"`
		Otherattr   string `json:"otherattr"`
		Maybeintatr int `json:"maybeintatr"`
}

var testresource *resource = &resource{
        Name: "test",
        Otherattr:  "test",
		Maybeintatr: 8,
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	jsonBytes, _ := json.Marshal(testresource)
	w.Write(jsonBytes)
}

func main() {
	http.HandleFunc("/resource", resourceHandler)
	log.Println("serving!")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
