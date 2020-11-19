package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

//从application.yml 获取参数
func GetEnv(env string) string {
	return fmt.Sprintf("%v", viper.Get(env))
}

//设备类型
var DEVICE_TYPES = []string{"MAC", "WINDOWS", "LINUX", "LINUX_CLI", "OPENWRT", "ANDROID", "IOS", "DOCKER", "HELPER"}
var WEB_TYPES = []string{"WEB", "WEB_SUPERADMIN"}
var DEVICE_TYPES_WITH_WEB = []string{"MAC", "WINDOWS", "LINUX", "LINUX_CLI", "OPENWRT", "ANDROID", "IOS", "DOCKER", "HELPER", "WEB", "WEB_SUPERADMIN"}
