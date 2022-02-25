package rental

type Rental struct {
	_movie      Movie
	_daysRented int
}

func NewRental(movie Movie, daysRented int) Rental {
	return Rental{
		_movie:      movie,
		_daysRented: daysRented,
	}
}
func (rcvr Rental) GetDaysRented() int {
	return rcvr._daysRented
}
func (rcvr Rental) GetMovie() Movie {
	return rcvr._movie
}
