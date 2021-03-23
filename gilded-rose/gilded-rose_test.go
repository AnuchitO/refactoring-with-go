package main

import (
	"github.com/kr/pretty"
	"reflect"
	"testing"
)

type Cases struct {
	inputs  []*Item
	outputs []*Item
}

func TestUpdateQuality(t *testing.T) {
	t.Run("fixture test cases", func(t *testing.T) {
		cases := []Cases{
			{inputs: []*Item{{"+5 Dexterity Vest", 10, 20}}, outputs: []*Item{{"+5 Dexterity Vest", 10 - 1, 20 - 1}}},
			{inputs: []*Item{{"Aged Brie", 2, 0}}, outputs: []*Item{{"Aged Brie", 2 - 1, 1}}},
			{inputs: []*Item{{"Elixir of the Mongoose", 5, 7}}, outputs: []*Item{{"Elixir of the Mongoose", 5 - 1, 6}}},
			{inputs: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}, outputs: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}},
			{inputs: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}, outputs: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 15, 1}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 15 - 1, 2}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 49}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10 - 1, 50}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 49}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5 - 1, 50}}},
			//{inputs: []*Item{{"Conjured Mana Cake", 3, 6}},  outputs: []*Item{{"Conjured Mana Cake", 3-1, 6-2}}}, // TODO: implement new requirement
		}

		assert(t, cases)
	})

	t.Run("normal items", func(t *testing.T) {
		cases := []Cases{
			{inputs: []*Item{{"+5 Dexterity Vest", 5, 15}}, outputs: []*Item{{"+5 Dexterity Vest", 4, 14}}},
			{inputs: []*Item{{"+5 Dexterity Vest", 4, 14}}, outputs: []*Item{{"+5 Dexterity Vest", 3, 13}}},
			{inputs: []*Item{{"+5 Dexterity Vest", 3, 13}}, outputs: []*Item{{"+5 Dexterity Vest", 2, 12}}},
			{inputs: []*Item{{"+5 Dexterity Vest", 2, 12}}, outputs: []*Item{{"+5 Dexterity Vest", 1, 11}}},
			{inputs: []*Item{{"+5 Dexterity Vest", 1, 11}}, outputs: []*Item{{"+5 Dexterity Vest", 0, 10}}},
			{inputs: []*Item{{"+5 Dexterity Vest", 0, 10}}, outputs: []*Item{{"+5 Dexterity Vest", -1, 8}}},
			{inputs: []*Item{{"+5 Dexterity Vest", -1, 8}}, outputs: []*Item{{"+5 Dexterity Vest", -2, 6}}},
			{inputs: []*Item{{"+5 Dexterity Vest", -2, 6}}, outputs: []*Item{{"+5 Dexterity Vest", -3, 4}}},
			{inputs: []*Item{{"+5 Dexterity Vest", -3, 4}}, outputs: []*Item{{"+5 Dexterity Vest", -4, 2}}},
			{inputs: []*Item{{"+5 Dexterity Vest", -4, 2}}, outputs: []*Item{{"+5 Dexterity Vest", -5, 0}}},
			{inputs: []*Item{{"+5 Dexterity Vest", -5, 0}}, outputs: []*Item{{"+5 Dexterity Vest", -6, 0}}},

			{inputs: []*Item{{"Elixir of the Mongoose", 1, 3}}, outputs: []*Item{{"Elixir of the Mongoose", 0, 2}}},
			{inputs: []*Item{{"Elixir of the Mongoose", 0, 2}}, outputs: []*Item{{"Elixir of the Mongoose", -1, 0}}},
			{inputs: []*Item{{"Elixir of the Mongoose", -1, 0}}, outputs: []*Item{{"Elixir of the Mongoose", -2, 0}}},
			{inputs: []*Item{{"Elixir of the Mongoose", -2, 0}}, outputs: []*Item{{"Elixir of the Mongoose", -3, 0}}},
		}

		assert(t, cases)
	})

	t.Run("Aged Brie items", func(t *testing.T) {

		cases := []Cases{
			{inputs: []*Item{{"Aged Brie", 2, 0}}, outputs: []*Item{{"Aged Brie", 1, 1}}},
			{inputs: []*Item{{"Aged Brie", 1, 1}}, outputs: []*Item{{"Aged Brie", 0, 2}}},
			{inputs: []*Item{{"Aged Brie", 0, 2}}, outputs: []*Item{{"Aged Brie", -1, 4}}},
			{inputs: []*Item{{"Aged Brie", -1, 4}}, outputs: []*Item{{"Aged Brie", -2, 6}}},
			{inputs: []*Item{{"Aged Brie", -2, 6}}, outputs: []*Item{{"Aged Brie", -3, 8}}},
			{inputs: []*Item{{"Aged Brie", -3, 8}}, outputs: []*Item{{"Aged Brie", -4, 10}}},
		}

		assert(t, cases)
	})

	t.Run("Sulfuras, Hand of Ragnaros items", func(t *testing.T) {

		cases := []Cases{
			{inputs: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}, outputs: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}},
			{inputs: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}, outputs: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}},
		}

		assert(t, cases)
	})

	t.Run("Backstage passes to a TAFKAL80ETC concert items", func(t *testing.T) {

		cases := []Cases{
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 12, 23}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 11, 24}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 11, 24}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 25}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 25}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 9, 27}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 9, 27}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 8, 29}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 8, 29}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 7, 31}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 7, 31}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 6, 33}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 6, 33}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 35}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 35}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 4, 38}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 4, 38}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 3, 41}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 3, 41}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 2, 44}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 2, 44}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 1, 47}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 1, 47}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 0, 50}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 0, 50}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -1, 0}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -1, 0}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -2, 0}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -2, 0}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -3, 0}}},
			{inputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -3, 0}}, outputs: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -4, 0}}},
		}

		assert(t, cases)
	})

	t.Run("Conjured items", func(t *testing.T) {
		t.Skip("TODO: implement new requirements")
	})


}

func assert(t *testing.T, cases []Cases) {
	for i, c := range cases {

		UpdateQuality(c.inputs)

		if !reflect.DeepEqual(c.inputs, c.outputs) {
			t.Errorf("test %s nubmer #%d not match : \n%#v", t.Name(), i, pretty.Diff(c.inputs, c.outputs))
		}
	}
}
