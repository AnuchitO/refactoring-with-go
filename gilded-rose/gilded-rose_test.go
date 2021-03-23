package main

import (
	"github.com/kr/pretty"
	"reflect"
	"testing"
)

func TestFixture(t *testing.T) {
	cases := []struct {
		inputs  []*Item
		outputs []*Item
	}{
		{inputs: []*Item{{"+5 Dexterity Vest", 10, 20}}, outputs: []*Item{{"+5 Dexterity Vest", 10 - 1, 20 - 1}}},
		{inputs: []*Item{{"Aged Brie", 2, 0}}, outputs: []*Item{{"Aged Brie", 2 - 1, 1}}},
		{inputs: []*Item{{"Elixir of the Mongoose", 5, 7}}, outputs: []*Item{{"Elixir of the Mongoose", 5 -1, 6}}},
		{inputs: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}, outputs: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}},
		{inputs: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}, outputs: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}},
		{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 15, 1}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 15 -1, 2}}},
		{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 49}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10 -1, 50}}},
		{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 49}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5 -1, 50}}},
		//{inputs: []*Item{{"Conjured Mana Cake", 3, 6}},  outputs: []*Item{{"Conjured Mana Cake", 3-1, 6-2}}}, // TODO: implement new requirement
	}

	for i, c := range cases {

		UpdateQuality(c.inputs)

		if !reflect.DeepEqual(c.inputs, c.outputs) {
			t.Errorf("test nubmer #%d not match : \n%#v", i, pretty.Diff(c.inputs, c.outputs))
		}
	}

}

func TestNormalItems(t *testing.T) {
	cases := []struct {
		inputs  []*Item
		outputs []*Item
	}{
		{inputs: []*Item{{"+5 Dexterity Vest", 5 ,15}}, outputs: []*Item{{"+5 Dexterity Vest", 4 ,14}}},
		{inputs: []*Item{{"+5 Dexterity Vest", 4 ,14}}, outputs: []*Item{{"+5 Dexterity Vest", 3 ,13}}},
		{inputs: []*Item{{"+5 Dexterity Vest", 3 ,13}}, outputs: []*Item{{"+5 Dexterity Vest", 2 ,12}}},
		{inputs: []*Item{{"+5 Dexterity Vest", 2 ,12}}, outputs: []*Item{{"+5 Dexterity Vest", 1 ,11}}},
		{inputs: []*Item{{"+5 Dexterity Vest", 1 ,11}}, outputs: []*Item{{"+5 Dexterity Vest", 0 ,10}}},
		{inputs: []*Item{{"+5 Dexterity Vest", 0 ,10}}, outputs: []*Item{{"+5 Dexterity Vest", -1, 8}}},
		{inputs: []*Item{{"+5 Dexterity Vest", -1, 8}}, outputs: []*Item{{"+5 Dexterity Vest", -2, 6}}},
		{inputs: []*Item{{"+5 Dexterity Vest", -2, 6}}, outputs: []*Item{{"+5 Dexterity Vest", -3, 4}}},
		{inputs: []*Item{{"+5 Dexterity Vest", -3, 4}}, outputs: []*Item{{"+5 Dexterity Vest", -4, 2}}},
		{inputs: []*Item{{"+5 Dexterity Vest", -4, 2}}, outputs: []*Item{{"+5 Dexterity Vest", -5, 0}}},
		{inputs: []*Item{{"+5 Dexterity Vest", -5, 0}}, outputs: []*Item{{"+5 Dexterity Vest", -6, 0}}},
	}

	for i, c := range cases {

		UpdateQuality(c.inputs)

		if !reflect.DeepEqual(c.inputs, c.outputs) {
			t.Errorf("test nubmer #%d not match : \n%#v", i, pretty.Diff(c.inputs, c.outputs))
		}
	}
}
