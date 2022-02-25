package rental

type Record struct {
	renter       string
	rentals      []Rental
	totalCharges float64
	points       int
}
