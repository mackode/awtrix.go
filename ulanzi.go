package main

import (
  "fmt"
  "time"
)

func main() {
  // Youtube
  f, v, err := youtubeStats()
  if err != nil {
    panic(err)
  }
  p := apiPayload{
    Text: fmt.Sprintf("%d/%d", f, v),
    Icon: 974,
    Duration: 4,
    Rainbow: true
  }
  err = postToAPI("youtube", p)
  if err != nil {
    panic(err)
  }

  // Odliczanie
  loc, err := time.LoadLocation("Europe/Warsaw")
  if err != nil {
    panic(err)
  }

  timerVal := DHMUntil(time.Date(2024, time.September, 1, 0, 0, 0, 0, loc))
  p = apiPayload{
    Text: timerVal,
    Duration: 4,
    Rainbow: true
  }
  err = postToAPI("countdown", p)
  if err != nil {
    panic(err)
  }

  p = apiPayload{
    Text: mon(),
    Icon: 23003,
    Duration: 4,
    Rainbow: true
  }
  err = postToAPI("dago", p)
  if err != nil {
    panic(err)
  }
}
