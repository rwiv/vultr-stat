package client

import (
	"fmt"
	"testing"

	"vultr-stat/pkg/lib/string/format"
)

func TestClient(t *testing.T) {
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
