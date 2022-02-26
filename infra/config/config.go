// Code generated .* DO NOT EDIT.

package config

import (
  "reflect"
  "time"
  c_adfapter "github.com/paulhindemith/category/pkg/adapter"
)

type Config struct {
  Github_com_paulhindemith_board_generator_arrows_board_generator_newarrow *ArrowConfig
  Github_com_paulhindemith_board_generator_arrows_publisher_newarrow *ArrowConfig
  Github_com_paulhindemith_board_generator_arrows_publisher_terminal *ArrowConfig
}
type ArrowConfig struct {
  Host string
  Delay time.Duration
}

var filedIdMap = map[string]string{
  "Github_com_paulhindemith_board_generator_arrows_board_generator_newarrow": "github.com/paulhindemith/board_generator/arrows/board_generator.NewArrow",
  "Github_com_paulhindemith_board_generator_arrows_publisher_newarrow": "github.com/paulhindemith/board_generator/arrows/publisher.NewArrow",
  "Github_com_paulhindemith_board_generator_arrows_publisher_terminal": "github.com/paulhindemith/board_generator/arrows/publisher.Terminal",
}
type FuncId = string
func (c *Config) GetArrows() map[FuncId]c_adfapter.ArrowConfigI {
  res := map[FuncId]c_adfapter.ArrowConfigI{}
  t_config := reflect.TypeOf(c).Elem()
  v_config := reflect.ValueOf(c).Elem()
  for i:=0;i<t_config.NumField();i++ {
		id := filedIdMap[t_config.Field(i).Name]
    ac := v_config.Field(i).Interface().(*ArrowConfig)
    if ac == nil {
      ac = &ArrowConfig{}
    }
    ac.updateDefault()
    res[id] = ac
	}
  return res
}
func (ac *ArrowConfig) GetHost() string {
  return ac.Host
}
func (ac *ArrowConfig) updateDefault() {
  if ac.Host == "" {
    ac.Host = "server1"
  }
}
func (ac *ArrowConfig) GetDelay() time.Duration {
  return ac.Delay
}
