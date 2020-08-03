package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetEnv(env string) string {
	return fmt.Sprintf("%v", viper.Get(env))
}
