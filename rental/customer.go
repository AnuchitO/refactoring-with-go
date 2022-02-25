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
	name := c.Name()
	totalCharges := c.totalCharge()
	points := c.totalFrequentRenterPoints()

	result := fmt.Sprintf("Rental Record for %s\n", name)
	for _, r := range c.rentals {
		title := r.Movie().Title()
		charge := r.getCharge()
		result += fmt.Sprintf("\t%s\t%.1f\n", title, charge)
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", totalCharges)
	result += fmt.Sprintf("You earned %v frequent renter points", points)
	return result
}

func renderPlainText(c Customer) string {
	return c.Statement()
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
