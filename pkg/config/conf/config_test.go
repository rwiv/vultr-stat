package conf

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	conf, err := GetConf()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(conf.VultrApiKey)
}
