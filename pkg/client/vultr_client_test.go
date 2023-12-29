package client

import (
	"fmt"
	"testing"

	"vultr-stat/pkg/lib/string/format"
)

func TestInstances(t *testing.T) {
	client := NewVultrClient()
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
	client := NewVultrClient()
	res, err := client.Account()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.Account.ToInfo().ToPretty())
}
