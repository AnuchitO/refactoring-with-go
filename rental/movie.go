package rental

const CHILDRENS = 2
const NEW_RELEASE = 1
const REGULAR = 0

type Movie struct {
	title     string
	priceCode int
}

func NewMovie(title string, priceCode int) Movie {
	return Movie{
		title:     title,
		priceCode: priceCode,
	}
}

func (m Movie) PriceCode() int {
	return m.priceCode
}

func (m Movie) Title() string {
	return m.title
}
