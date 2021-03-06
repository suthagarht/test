package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "log"
  _"encoding/json"
)


func main () {

  // Get the http request data
  var netClient = http.Client{}
  url := "https://country.register.gov.uk/records.json"

  request, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
    log.Fatal(err)
  }

  response, err := netClient.Do(request)
  if err != nil {
    log.Fatal(err)
  }

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(body))
}
