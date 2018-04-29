package main

import (
"fmt"
"time"
"encoding/json"
)

type Item struct {
      Country      string `json:"country"`
      OfficialName string `json:"official-name"`
      Name         string `json:"name"`
      CitizenNames string `json:"citizen-names"`
}

type ID  struct {
    IndexEntryNumber string    `json:"index-entry-number"`
    EntryNumber      string    `json:"entry-number"`
    EntryTimestamp   time.Time `json:"entry-timestamp"`
    Key              string    `json:"key"`
    Item
}

type Message struct {
  ID
}

func main () {

  blob := []byte(`{"PT":{"index-entry-number":"147","entry-number":"147","entry-timestamp":"2016-04-05T13:23:05Z","key":"PT","item":[{"country":"PT","official-name":"The Portuguese Republic","name":"Portugal","citizen-names":"Portuguese"}]}}`)

  var m Message

  json.Unmarshal(blob, &m)

  fmt.Println("hello")

  fmt.Println(m.Item.Country)
}
