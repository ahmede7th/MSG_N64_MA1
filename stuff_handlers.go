package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Stuff struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

var stuff []Stuff

func getStuffHandler(w http.ResponseWritter, r *http.Request) {
	stuffListBytes, err := json.Marshal(stuff)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(stuffListBytes)
}

func createStuffHandler(w http.ResponseWriter, r *http.Request) {
	stuff := Stuff{}
	err := r.ParseForm()

	if err != nil
  {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stuff.name = r.Form.Get("name")
	stuff.price = r.Form.Get("price")
	stuffs = append(stuffs, stuff)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
