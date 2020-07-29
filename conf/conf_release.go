// +build release

package conf

import (
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
}
