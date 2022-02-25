package rental

import "fmt"

type Customer struct {
	name    string
	rentals []Rental
}

func NewCustomer(name string) Customer {
	return Customer{
		name:    name,
		rentals: []Rental{},
	}
}
func (c Customer) AddRental(arg Rental) Customer {
	c.rentals = append(c.rentals, arg)
	return c
}
func (c Customer) Name() string {
	return c.name
}
func (c Customer) Statement() string {
	frequentRenterPoints := 0
	result := fmt.Sprintf("Rental Record for %v\n", c.Name())
	frequentRenterPoints = totalFrequentRenterPoints(c)
	for _, r := range c.rentals {
		result += fmt.Sprintf("\t%v\t%.1f\n", r.Movie().Title(), r.getCharge())
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", totalCharge(c))
	result += fmt.Sprintf("You earned %v frequent renter points", frequentRenterPoints)
	return result
}

func totalFrequentRenterPoints(c Customer) (frequentRenterPoints int) {
	for _, r := range c.rentals {
		frequentRenterPoints += getFrequentRenterPoints(r)
	}
	return frequentRenterPoints
}

func totalCharge(c Customer) (result float64) {
	for _, r := range c.rentals {
		result += r.getCharge()
	}
	return result
}

func getFrequentRenterPoints(r Rental) int {
	if r.Movie().PriceCode() == NEW_RELEASE && r.DaysRented() > 1 {
		return 2
	}
	return 1
}

func (r Rental) getCharge() (result float64) {
	switch r.Movie().PriceCode() {
	case REGULAR:
		result += 2
		if r.DaysRented() > 2 {
			result += float64(r.DaysRented()-2) * 1.5
		}
	case NEW_RELEASE:
		result += float64(r.DaysRented()) * 3.0
	case CHILDRENS:
		result += 1.5
		if r.DaysRented() > 3 {
			result += float64(r.DaysRented()-3) * 1.5
		}
	}
	return result
}
