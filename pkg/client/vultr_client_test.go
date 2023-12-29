package client

import (
	"fmt"
	"testing"

	"vultr-stat/pkg/lib/date"
	"vultr-stat/pkg/lib/string/format"
)

func TestOs(t *testing.T) {
	client := NewVultrClient()
	res, err := client.Os()
	if err != nil {
		t.Fatal(err)
	}
	for _, os := range res.Os {
		json := format.ToJsonPretty(os)
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
