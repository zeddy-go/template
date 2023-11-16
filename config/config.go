package config

import (
	_ "embed"
	"github.com/spf13/viper"
	"github.com/zeddy-go/core/container"
	"github.com/zeddy-go/core/contract"
	"strings"
)

//go:embed config.yaml
var conf string

type Module struct{}

func (m Module) Register(sub contract.IModule) {
	//TODO implement me
	panic("implement me")
}

func (m Module) Init() {
	container.Register(func() *viper.Viper {
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
