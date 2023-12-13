package main

import (
  "fmt"
  "net/http"
  //"net/url"
  "PlayInMPV/Exec"
)

func main(){
  
  // Test handler
  http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request){
    fmt.Fprintf(w, "<h1>Hello World</h1>")
  })
  
  http.HandleFunc("/play", func (w http.ResponseWriter, req *http.Request){
    
    options := req.URL.Query()
    
    url, ok := options["url"]
    
    if !ok || len(url) < 1{
		fmt.Fprint(w, "Error: you should specify an url")
		return
	}
	
	if len(url) > 1{
		fmt.Fprint(w, "Error: you should specify only 1 url")
		return
	}

    fmt.Fprint(w,"now playing ", url[0], "\n")

    go Exec.OpenLink(options["url"][0])
  })

  
  fmt.Println("Starting server")
  http.ListenAndServe(":8090",nil)
  fmt.Println("Finishing server")
}


