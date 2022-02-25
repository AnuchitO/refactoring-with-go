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
	rc := Record{
		renter:       c.Name(),
		rentals:      c.rentals,
		totalCharges: c.totalCharge(),
		points:       c.totalFrequentRenterPoints(),
	}

	return renderPlainText(rc)
}

func renderPlainText(rc Record) string {
	result := fmt.Sprintf("Rental Record for %s\n", rc.renter)
	for _, r := range rc.rentals {
		title := r.Movie().Title()
		charge := r.getCharge() // Pre-mature Rental's method. It should belong to Movie.?
		result += fmt.Sprintf("\t%s\t%.1f\n", title, charge)
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", rc.totalCharges)
	result += fmt.Sprintf("You earned %v frequent renter points", rc.points)
	return result
}

func (c Customer) totalFrequentRenterPoints() (result int) {
	for _, r := range c.rentals {
		result += r.getFrequentRenterPoints()
	}
	return result
}

func (c Customer) totalCharge() (result float64) {
	for _, r := range c.rentals {
		result += r.getCharge()
	}
	return result
}

func (r Rental) getFrequentRenterPoints() int {
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
