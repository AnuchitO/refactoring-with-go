package rental

type Rental struct {
	_movie      *Movie
	_daysRented int
}

func NewRental(movie *Movie, daysRented int) (rcvr *Rental) {
	rcvr = &Rental{}
	rcvr._movie = movie
	rcvr._daysRented = daysRented
	return
}
func (rcvr *Rental) GetDaysRented() int {
	return rcvr._daysRented
}
func (rcvr *Rental) GetMovie() *Movie {
	return rcvr._movie
}
