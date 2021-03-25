package main

import (
	"github.com/kr/pretty"
	"reflect"
	"testing"
)

type Cases struct {
	in  []*Item
	out []*Item
}

func TestUpdateQuality(t *testing.T) {
	t.Run("fixture test cases", func(t *testing.T) {
		cases := []Cases{
			{in: []*Item{{"+5 Dexterity Vest", 10, 20}}, out: []*Item{{"+5 Dexterity Vest", 10 - 1, 20 - 1}}},
			{in: []*Item{{"Aged Brie", 2, 0}}, out: []*Item{{"Aged Brie", 2 - 1, 1}}},
			{in: []*Item{{"Elixir of the Mongoose", 5, 7}}, out: []*Item{{"Elixir of the Mongoose", 5 - 1, 6}}},
			{in: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}, out: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}},
			{in: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}, out: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 15, 1}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 15 - 1, 2}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 49}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10 - 1, 50}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 49}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5 - 1, 50}}},
			//{in: []*Item{{"Conjured Mana Cake", 3, 6}},  out: []*Item{{"Conjured Mana Cake", 3-1, 6-2}}}, // TODO: implement new requirement
		}

		assert(t, cases)
	})

	t.Run("normal items", func(t *testing.T) {
		cases := []Cases{
			{in: []*Item{{"+5 Dexterity Vest", 5, 15}}, out: []*Item{{"+5 Dexterity Vest", 4, 14}}},
			{in: []*Item{{"+5 Dexterity Vest", 4, 14}}, out: []*Item{{"+5 Dexterity Vest", 3, 13}}},
			{in: []*Item{{"+5 Dexterity Vest", 3, 13}}, out: []*Item{{"+5 Dexterity Vest", 2, 12}}},
			{in: []*Item{{"+5 Dexterity Vest", 2, 12}}, out: []*Item{{"+5 Dexterity Vest", 1, 11}}},
			{in: []*Item{{"+5 Dexterity Vest", 1, 11}}, out: []*Item{{"+5 Dexterity Vest", 0, 10}}},
			{in: []*Item{{"+5 Dexterity Vest", 0, 10}}, out: []*Item{{"+5 Dexterity Vest", -1, 8}}},
			{in: []*Item{{"+5 Dexterity Vest", -1, 8}}, out: []*Item{{"+5 Dexterity Vest", -2, 6}}},
			{in: []*Item{{"+5 Dexterity Vest", -2, 6}}, out: []*Item{{"+5 Dexterity Vest", -3, 4}}},
			{in: []*Item{{"+5 Dexterity Vest", -3, 4}}, out: []*Item{{"+5 Dexterity Vest", -4, 2}}},
			{in: []*Item{{"+5 Dexterity Vest", -4, 2}}, out: []*Item{{"+5 Dexterity Vest", -5, 0}}},
			{in: []*Item{{"+5 Dexterity Vest", -5, 0}}, out: []*Item{{"+5 Dexterity Vest", -6, 0}}},

			{in: []*Item{{"Elixir of the Mongoose", 1, 3}}, out: []*Item{{"Elixir of the Mongoose", 0, 2}}},
			{in: []*Item{{"Elixir of the Mongoose", 0, 2}}, out: []*Item{{"Elixir of the Mongoose", -1, 0}}},
			{in: []*Item{{"Elixir of the Mongoose", -1, 0}}, out: []*Item{{"Elixir of the Mongoose", -2, 0}}},
			{in: []*Item{{"Elixir of the Mongoose", -2, 0}}, out: []*Item{{"Elixir of the Mongoose", -3, 0}}},
		}

		assert(t, cases)
	})

	t.Run("Aged Brie items", func(t *testing.T) {

		cases := []Cases{
			{in: []*Item{{"Aged Brie", 2, 0}}, out: []*Item{{"Aged Brie", 1, 1}}},
			{in: []*Item{{"Aged Brie", 1, 1}}, out: []*Item{{"Aged Brie", 0, 2}}},
			{in: []*Item{{"Aged Brie", 0, 2}}, out: []*Item{{"Aged Brie", -1, 4}}},
			{in: []*Item{{"Aged Brie", -1, 4}}, out: []*Item{{"Aged Brie", -2, 6}}},
			{in: []*Item{{"Aged Brie", -2, 6}}, out: []*Item{{"Aged Brie", -3, 8}}},
			{in: []*Item{{"Aged Brie", -3, 8}}, out: []*Item{{"Aged Brie", -4, 10}}},
		}

		assert(t, cases)
	})

	t.Run("Sulfuras, Hand of Ragnaros items", func(t *testing.T) {

		cases := []Cases{
			{in: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}, out: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}},
			{in: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}, out: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}},
		}

		assert(t, cases)
	})

	t.Run("Backstage passes to a TAFKAL80ETC concert items", func(t *testing.T) {

		cases := []Cases{
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 12, 23}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 11, 24}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 11, 24}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 25}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 25}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 9, 27}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 9, 27}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 8, 29}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 8, 29}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 7, 31}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 7, 31}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 6, 33}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 6, 33}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 35}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 35}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 4, 38}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 4, 38}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 3, 41}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 3, 41}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 2, 44}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 2, 44}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 1, 47}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 1, 47}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 0, 50}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 0, 50}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -1, 0}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -1, 0}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -2, 0}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -2, 0}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -3, 0}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -3, 0}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", -4, 0}}},
		}

		assert(t, cases)
	})

	t.Run("Conjured items", func(t *testing.T) {
		t.Skip("TODO: implement new requirements")
	})


}

func assert(t *testing.T, cases []Cases) {
	for i, c := range cases {

		UpdateQuality(c.in)

		if !reflect.DeepEqual(c.in, c.out) {
			t.Errorf("test %s nubmer #%d not match : \n%#v", t.Name(), i, pretty.Diff(c.in, c.out))
		}
	}
}
