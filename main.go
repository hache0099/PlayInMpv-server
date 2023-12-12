package main

import (
  "fmt"
  "net/http"
)

func main(){
  http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request){
    fmt.Fprintf(w, "Hello World")
  })

  
  http.ListenAndServe(":8090",nil)
}


