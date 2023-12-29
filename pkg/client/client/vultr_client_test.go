package client

import (
	"testing"
)

func TestClient(t *testing.T) {
	a := "hello"
	if a != "hello" {
		t.Fatal("error")
	}
}
