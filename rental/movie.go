package rental

const CHILDRENS = 2
const NEW_RELEASE = 1
const REGULAR = 0

type Movie struct {
	_title     string
	_priceCode int
}

func NewMovie(title string, priceCode int) (rcvr *Movie) {
	rcvr = &Movie{}
	rcvr._title = title
	rcvr._priceCode = priceCode
	return
}
func (rcvr *Movie) GetPriceCode() int {
	return rcvr._priceCode
}
func (rcvr *Movie) GetTitle() string {
	return rcvr._title
}
func (rcvr *Movie) SetPriceCode(arg int) {
	rcvr._priceCode = arg
}
