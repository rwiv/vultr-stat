package apath

import (
	"fmt"
	"testing"
)

func TestPath(t *testing.T) {
	p1, err := GetProjectAbsPath()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(p1)
	p2, err := GetProjectAbsPathBy("/foo/bar/vultr-stat/haha/hello.txt", "/")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(p2)
}

func TestIsAbsPath(t *testing.T) {
	p1 := "/foo/bar.txt"
	if IsAbsPath(p1) != true {
		t.Fatal("error")
	}

	p2 := "../bar.txt"
	if IsAbsPath(p2) != false {
		t.Fatal("error")
	}
}
