package table

import (
	"github.com/kr/pretty"
	"reflect"
	"testing"
)

type Cases struct {
	in  []*Item
	out []*Item
}

func TestTableExample(t *testing.T) {
	t.Run("test table", func(t *testing.T) {
		cases := []Cases{
			{in: []*Item{{"Dexterity", 10}}, out: []*Item{{"Dexterity", 9}}},
			{in: []*Item{{"Aged Brie", 2}}, out: []*Item{{"Aged Brie", 1}}},
			{in: []*Item{{"Elixir", 5}}, out: []*Item{{"Elixir", 4}}},
			{in: []*Item{{"Sulfuras", 0}}, out: []*Item{{"Sulfuras", 0}}},
			{in: []*Item{{"Backstage passes", 15}}, out: []*Item{{"Backstage passes", 14}}},
		}

		assert(t, cases)
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
