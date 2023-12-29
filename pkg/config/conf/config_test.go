package conf

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	conf := GetConf()
	fmt.Println(conf.VultrApiKey)
}
