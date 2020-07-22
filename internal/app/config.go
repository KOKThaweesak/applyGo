package app

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Configs struct {
	EnablePing bool  `mapstructure:"enable_ping"`
	Mongo      Mongo `mapstructure:"mongo"`

	vn *viper.Viper
}
type Mongo struct {
	Addresses []string `mapstructure: "addresses"`
	Username  string   `mapstructure: "username"`

	Password      string `mapstructure: "password"`
	Database      string `mapstructure: "database"`
	HasConnection bool
}

func (M *Mongo) bindingConnection() error {

	M.HasConnection = true

	fmt.Printf("Bind connection - %+v\n", M)
	return nil
}

func (c *Configs) InitWithState(s *string) error {
	vn := viper.New()
	c.vn = vn
	vn.AddConfigPath("./configs")
	vn.SetConfigName(fmt.Sprintf("config.%s", *s))

	if err := vn.ReadInConfig(); err != nil {
		return err
	}

	if err := c.binding(); err != nil {
		return err
	}

	vn.WatchConfig()
	vn.OnConfigChange(func(in fsnotify.Event) {
		vn.Unmarshal(&c)
		log.Printf("\nnew config: %+v", c)
	})
	return nil
}

func (c *Configs) binding() error {
	if err := c.vn.Unmarshal(&c); err != nil {
		return err
	}

	//  binding addional data
	if err := c.Mongo.bindingConnection(); err != nil {
		return err
	}

	return nil
}
