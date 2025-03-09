package client

import (
	"fmt"
	"testing"

	"github.com/rwiv/vultr-stat/pkg/common"
	"github.com/rwiv/vultr-stat/pkg/lib/string/format"
)

func TestInstances(t *testing.T) {
	env, err := common.ReadEnv()
	if err != nil {
		t.Fatal(err)
	}
	client := NewVultrClient(env.VultrApiKey)
	res, err := client.Instances()
	if err != nil {
		t.Fatal(err)
	}
	for _, instance := range res.Instances {
		json := format.ToJsonPretty(instance)
		fmt.Println(json)
	}
}

func TestAccount(t *testing.T) {
	env, err := common.ReadEnv()
	if err != nil {
		t.Fatal(err)
	}
	client := NewVultrClient(env.VultrApiKey)
	res, err := client.Account()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.Account.ToInfo().ToPretty())
}
