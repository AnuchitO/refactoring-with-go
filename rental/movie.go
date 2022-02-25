package rental

type PriceCode int

const (
	Childrens PriceCode = iota
	NewRelease
	Regular
)

type Movie struct {
	title     string
	priceCode PriceCode
}

func NewMovie(title string, priceCode PriceCode) Movie {
	return Movie{
		title:     title,
		priceCode: priceCode,
	}
}

func (m Movie) PriceCode() PriceCode {
	return m.priceCode
}

func (m Movie) Title() string {
	return m.title
}
