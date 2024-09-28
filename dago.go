package main

import (
	"bufio"
	"os"
	"os/user"
	"path"
	"regexp"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func mon() string {
  usr, err := user.Current()
  if err != nil {
    panic(err)
  }

  logf := path.Join(usr.HomeDir, "data/monlog.txt")
  file, err := os.Open(logf)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  var lastLine string
  for scanner.Scan() {
    lastLine = scanner.Text()
  }
  if err := scanner.Err(); err != nil {
    panic(err)
  }

  re := regexp.MustCompile(`\d+`)
  match := re.FindString(lastLine)
  n, err := strconv.ParseInt(match, 10, 64)
  if err != nil {
    panic(err)
  }

  n = n / 1000
  p := message.NewPrinter(language.English)
  return p.Sprintf("%d", n)
}
