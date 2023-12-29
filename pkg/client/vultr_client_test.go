package client

import (
	"fmt"
	"testing"

	"vultr-stat/pkg/lib/date"
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
	json := format.ToJsonPretty(res.Account)
	fmt.Println(json)

	tm, err := date.ByRFC3339String(res.Account.LastPaymentDate)
	if err != nil {
		t.Fatal(err)
	}
	dateStr := date.ToPrettyString(tm)
	fmt.Println(dateStr)
	fmt.Println(date.ToRFC3339String(tm))
}
