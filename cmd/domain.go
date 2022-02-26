package cmd

import (
  "github.com/paulhindemith/pkg-strategist/util"
  "github.com/paulhindemith/board_generator/arrows/board_generator"
  "github.com/paulhindemith/board_generator/arrows/publisher"
  c_generator "github.com/paulhindemith/category/generator"
)

var (
  generator = []interface{}{
    board_generator.NewArrow,
    publisher.NewArrow,
  }
)
func BoardGeneratorNewAdapter() *c_generator.Adapter {
  return &c_generator.Adapter{
    Name: "board_generator",
    ProjectPath: util.ProjectPath("board_generator"),
    NewArrows: generator,
  }
}
