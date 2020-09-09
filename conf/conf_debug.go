// +build !release

package conf

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	// any approach to require this configuration into your program.
	str, _ := os.Getwd()
	str = strings.Replace(str, "/test", "", 1)
	str = strings.Replace(str, "/controller", "", 1)
	url := str + "/application.yml"
	data, err := ioutil.ReadFile(url)
	if err != nil {
		panic(err)
	}

	viper.ReadConfig(bytes.NewBuffer(data))
}
