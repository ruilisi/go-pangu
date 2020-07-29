package conf

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

var DEVICES = map[string]bool{"WINDOWS": true, "MAC": true, "ANDROID": true, "IOS": true}

func init() {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	// any approach to require this configuration into your program.
	data, err := ioutil.ReadFile("conf/application.yml")
	if err != nil {
		panic(err)
	}

	viper.ReadConfig(bytes.NewBuffer(data))
}

func GetEnv(env string) string {
	if viper.Get("GIN_ENV").(string) != "production" {
		return fmt.Sprintf("%v", viper.Get(env))
	} else {
		return os.Getenv(env)
	}
}
