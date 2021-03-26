package main

import "github.com/kr/pretty"

type Item struct {
	kind            string
	sellIn, quality int
}

type Operator interface {
	Apply(item Item) Item
	isMatch(kind string) bool
}

type Operation struct {
	Operators []Operator
}

func (op Operation) Operate(item Item) Item {
	for _, o := range op.Operators {
		if o.isMatch(item.kind) {
			item = o.Apply(item)
		}
	}
	return item
}

type AgedBrie struct{}

func (AgedBrie) isMatch(kind string) bool {
	return kind == "AgedBrie"
}
func (AgedBrie) Apply(item Item) Item {
	item.quality++
	item.sellIn--
	return item
}

type Backstage struct{}

func (Backstage) isMatch(kind string) bool {
	return kind == "Backstage"
}
func (Backstage) Apply(item Item) Item {
	item.quality += 3
	item.sellIn--
	return item
}

func main() {
	ab := AgedBrie{}
	bs := Backstage{}
	ops := Operation{[]Operator{ab, bs}}

	items := []Item{{kind: "Backstage", sellIn: 10, quality: 5}, {kind: "AgedBrie", sellIn: 8, quality: 12}}

	var newItems []Item
	pretty.Println(items)
	for _, item := range items {
		newItems = append(newItems, ops.Operate(item))
	}
	pretty.Println(newItems)
}
