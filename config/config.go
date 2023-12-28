package config

import (
	_ "embed"
	"strings"

	"github.com/spf13/viper"
	"github.com/zeddy-go/zeddy/container"
)

//go:embed config.yaml
var conf string

func init() {
	container.Bind[*viper.Viper](func() *viper.Viper {
		c := viper.New()
		c.SetConfigType("yaml")
		err := c.ReadConfig(strings.NewReader(conf))
		if err != nil {
			panic(err)
		}
		c.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
		c.AutomaticEnv()
		return c
	})
}
