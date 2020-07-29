// +build !release

package conf

import (
	"bytes"
	"io/ioutil"

	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	// any approach to require this configuration into your program.
	data, err := ioutil.ReadFile("application.yml")
	if err != nil {
		panic(err)
	}

	viper.ReadConfig(bytes.NewBuffer(data))
}
