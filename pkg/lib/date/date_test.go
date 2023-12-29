package date

import (
	"fmt"
	"testing"
)

func TestDate(t *testing.T) {
	data := "2023-08-06T18:39:15.145945500"

	time, err := ByIsoString(data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(time)

	str := ToIsoString(time)
	fmt.Println(str)

	str2 := ToPrettyString(time)
	fmt.Println(str2)
}
