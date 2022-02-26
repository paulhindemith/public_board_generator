// Code generated .* DO NOT EDIT.

package main

import (
  "context"
  "github.com/paulhindemith/pkg-strategist/adapter"
  "github.com/paulhindemith/board_generator/infra/server1"
)
func main() {
  adapter.Main(context.Background(), server1.NewAdapter)
}
