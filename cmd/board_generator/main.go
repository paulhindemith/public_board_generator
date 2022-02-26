package main

import (
  "log"
  "github.com/paulhindemith/board_generator/cmd"
)

func main() {
  a := cmd.BoardGeneratorNewAdapter()
  if err := a.Start(); err != nil {
    log.Fatal(err)
  }
}
