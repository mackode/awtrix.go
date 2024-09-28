package main

import (
  "context"
  "google.golang.org/api/option"
  "google.golang.org/api/youtube/v3"
  "log"
)

const ChannelID = "XYZ"
const ApiKey = "XYZ"

func youtubeStats() (uint64, uint64, error) {
  ctx := context.Background()
  service, err := youtube.NewService(ctx, option.WithAPIKey(ApiKey))
  resp, err := service.Channels.List([]string{"statistics"}).Id(ChannelID).Do()
  if err != nil {
    log.Fatalf("%v", err)
  }
  if len(resp.Items) == 0 {
    log.Fatalf("Channel not found")
  }

  stat := resp.Items[0].Statistics
  return stat.SubscriberCount, stat.VideoCount, nil
}
