package main

import (
	"github.com/kr/pretty"
	"reflect"
	"testing"
)

func TestFixture(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []*Item
		outputs []*Item
	}{
		{name: "#1",inputs: []*Item{{"+5 Dexterity Vest", 10, 20}}, outputs: []*Item{{"+5 Dexterity Vest", 10 - 1, 20 - 1}}},
		{name: "#2",inputs: []*Item{{"Aged Brie", 2, 0}}, outputs: []*Item{{"Aged Brie", 2 - 1, 1}}},
		{name: "#3",inputs: []*Item{{"Elixir of the Mongoose", 5, 7}}, outputs: []*Item{{"Elixir of the Mongoose", 5 -1, 6}}},
		{name: "#4",inputs: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}, outputs: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}},
		{name: "#5",inputs: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}, outputs: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}},
		{name: "#6",inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 15, 1}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 15 -1, 2}}},
		{name: "#7",inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 49}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10 -1, 50}}},
		{name: "#8",inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 49}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5 -1, 50}}},
		//{name: "#9",inputs: []*Item{{"Conjured Mana Cake", 3, 6}},  outputs: []*Item{{"Conjured Mana Cake", 3-1, 6-2}}}, // TODO: implement new requirement
	}

	for _, c := range cases {

		UpdateQuality(c.inputs)

		if !reflect.DeepEqual(c.inputs, c.outputs) {
			t.Errorf("%s not match : \n%#v", c.name, pretty.Diff(c.inputs, c.outputs))
		}
	}

}
