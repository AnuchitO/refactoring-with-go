package main

import (
	"testing"
	"github.com/kr/pretty"
)


func Test_Foo(t *testing.T) {
	var items = []*Item{
		&Item{"foo", 0, 0},
	}

	UpdateQuality(items)

	if items[0].name != "fixme" {
		pretty.Println(items)
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].name)
	}
}
