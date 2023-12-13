package main

import (
  "fmt"
  "net/http"
  //"net/url"
  "PlayInMPV/Exec"
)

func main(){
  http.HandleFunc("/play", func (w http.ResponseWriter, req *http.Request){
    fmt.Fprintf(w, "<h1>Hello World</h1>")
    options := req.URL.Query()

    fmt.Println(options)

    err := Exec.OpenLink(options["url"][0])
    if err != nil{
      panic(err)
    }
  })

  
  
  http.ListenAndServe(":8090",nil)
}


