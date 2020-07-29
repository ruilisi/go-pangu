package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

var DEVICES = map[string]bool{"WINDOWS": true, "MAC": true, "ANDROID": true, "IOS": true}

func GetEnv(env string) string {
	return fmt.Sprintf("%v", viper.Get(env))
}
