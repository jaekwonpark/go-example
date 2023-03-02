package main

import (
  "bytes"
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  "log"
)

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/PrismGateway/j_spring_security_check", rcv)

  err := http.ListenAndServeTLS(":9440", "server.cert", "server.key", r)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }

}

func rcv(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("received request to /rcv\n")
  buf := new(bytes.Buffer)
  buf.ReadFrom(r.Body)
  postData := buf.String()

  for name, values := range r.Header {
    // Loop over all values for the name.
    for _, value := range values {
      fmt.Printf("%s:%s\n",name, value)
    }
  }

  fmt.Printf("postData:%s\n", postData)
}