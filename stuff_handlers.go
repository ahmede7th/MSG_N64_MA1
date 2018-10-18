package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Stuff struct {
	Name  string `json:"Name"`
	Price string `json:"Price"`
}

var stuff []Stuff

func getStuffHandler(w http.ResponseWriter, r *http.Request) {
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

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stuff.Name = r.Form.Get("Name")
	stuff.Price = r.Form.Get("Price")
	
	stuffs = append(stuffs, stuff)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
