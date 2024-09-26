package main

import (
  "fmt"
  "time"
)

func DHMUntil(until time.Time) string {
  dur := time.Until(until)
  days := int(dur.Hours() / 24)
  hours := int(dur.Hours()) % 24
  mins := int(dur.Minutes()) % 60
  return fmt.Sprintf("%02d:%02d:%02d", days, hours, mins)
}

