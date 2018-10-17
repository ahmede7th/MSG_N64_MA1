package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestHandler(t *testing.T) {
  req, err := http.NewRequest("GET", "", nil)

  if err != nil {
    t.Fatal(err)
  }


  recorder := httptest.NewRecorder()

  hf := http.HandlerFunc(handler)

  hf.ServeHTTP(recorder, req)

  if status := recorder.Code; status != http.StatusOk {
    t.Errorf("handler returned wrong status code: got %v want %v")
  }



  //Check for expected vs actual
  /////////////////////////////
  ////////////////////////////
  //=)
}
