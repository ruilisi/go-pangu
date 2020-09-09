package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetEnv(env string) string {
	return fmt.Sprintf("%v", viper.Get(env))
}

var DEVICE_TYPES = []string{"MAC", "WINDOWS", "LINUX", "LINUX_CLI", "OPENWRT", "ANDROID", "IOS", "DOCKER", "HELPER"}
var WEB_TYPES = []string{"WEB", "WEB_SUPERADMIN"}
var DEVICE_TYPES_WITH_WEB = []string{"MAC", "WINDOWS", "LINUX", "LINUX_CLI", "OPENWRT", "ANDROID", "IOS", "DOCKER", "HELPER", "WEB", "WEB_SUPERADMIN"}
