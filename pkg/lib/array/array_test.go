package array

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func f1(ps []Person) {
	ps[0].SetName("hello")
}

func TestArrayPointer(t *testing.T) {
	p1 := Person{"foo"}
	p2 := Person{"bar"}

	ps := make([]Person, 0)
	ps = append(ps, p1)
	ps = append(ps, p2)

	f1(ps)

	// 결론: 매개변수로 들어온 배열의 모든 원자들은 전부 복사된다
	fmt.Println(p1)
	fmt.Println(p2)
}
