package main

import (
  "github.com/nogizaka-api/router"
)

func main() {
  r := router.New()
  r.Logger.Fatal(r.Start(":8000"))
}