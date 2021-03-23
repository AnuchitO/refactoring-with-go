package main

type Item struct {
	name            string
	sellIn, quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		if item.name == "Sulfuras, Hand of Ragnaros" {
			continue
		}

		if item.name != "Aged Brie" && item.name != "Backstage passes to a TAFKAL80ETC concert" {
			updateQuality(item, item.quality-1)
		} else {
			updateQuality(item, item.quality+1)
			if item.name == "Backstage passes to a TAFKAL80ETC concert" {
				if item.sellIn < 11 {
					updateQuality(item, item.quality+1)
				}
				if item.sellIn < 6 {
					updateQuality(item, item.quality+1)
				}
			}
		}

		item.sellIn = item.sellIn - 1

		if item.sellIn < 0 {
			if item.name != "Aged Brie" {
				if item.name != "Backstage passes to a TAFKAL80ETC concert" {
					updateQuality(item, item.quality-1)
				} else {
					updateQuality(item, item.quality-item.quality)
				}
			} else {
				updateQuality(item, item.quality+1)
			}
		}
	}

}

func updateQuality(item *Item, quality int) {
	if quality > 50 {
		item.quality = 50
		return
	}

	if quality < 0 {
		item.quality = 0
		return
	}

	item.quality = quality
}
