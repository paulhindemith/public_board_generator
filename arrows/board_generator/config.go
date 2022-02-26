package board_generator

import(
  "time"
  "go.uber.org/zap"
  "github.com/spf13/viper"
)

type Config struct {
  Targets []*Target `mapstructure:"board_generator_targets"`
  Wait time.Duration `mapstructure:"board_generator_wait"`
  OandaSecret string `mapstructure:"oanda_api_secret"`
}
type Target struct {
  PriceName string `mapstructure:"price_name"`
  Timeout time.Duration  `mapstructure:"timeout"`
  ReadLimit int `mapstructure:"read_limit"`
  WriteLimigt int `mapstructure:"write_limit"`
}

func (g *Generator) getConfig() *Config {
  var config Config
  viper.BindEnv("oanda_api_secret")
  if err := viper.Unmarshal(&config); err != nil {
    g.logger.Fatal("config file Unmarshal error", zap.String("error", err.Error()))
  }
  return &config
}
