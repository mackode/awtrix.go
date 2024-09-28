package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "net/http"
)

const baseURL = "http://192.168.0.1/api/custom"

type apiPayload struct {
  Text string `json:"text"`
  Rainbow bool `json:"rainbow"`
  Duration int `json:"duration"`
  Icon int `json:"icon"`
}

func postToAPI(name string, p apiPayload) error {
  url := baseURL + "?name=" + name
  jsonBytes, err := json.Marshal(p)
  if err != nil {
    return err
  }

  resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("%v", resp.StatusCode)
  }
  return nil
}
