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
		charge := r.getCharge(r.Movie(), r.daysRented) // Pre-mature Rental's method. It should belong to Movie.?
		result += fmt.Sprintf("\t%s\t%.1f\n", title, charge)
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", rc.totalCharges)
	result += fmt.Sprintf("You earned %v frequent renter points", rc.points)
	return result
}

func renderHtml(rc Record) string {
	result := fmt.Sprintf("<h1>Rental Record for <em>%s</em></h1>\n", rc.renter)
	result += "<table>\n"
	for _, r := range rc.rentals {
		title := r.Movie().Title()
		charge := r.getCharge(r.Movie(), r.daysRented)
		result += fmt.Sprintf("\t<tr><td>%s</td><td>%.1f</td></tr>\n", title, charge)
	}
	result += "</table>\n"
	result += fmt.Sprintf("<p>Amount owed is <em>%.1f</em></p>\n", rc.totalCharges)
	result += fmt.Sprintf("<p>You earned <em>%v</em> frequent renter points</p>", rc.points)
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
		result += r.getCharge(r.Movie(), r.daysRented)
	}
	return result
}

func (r Rental) getFrequentRenterPoints() int {
	if r.Movie().PriceCode() == NEW_RELEASE && r.DaysRented() > 1 {
		return 2
	}
	return 1
}

func (r Rental) getCharge(m Movie, daysRented int) (result float64) {
	switch m.PriceCode() {
	case REGULAR:
		result += 2
		if daysRented > 2 {
			result += float64(daysRented-2) * 1.5
		}
	case NEW_RELEASE:
		result += float64(daysRented) * 3.0
	case CHILDRENS:
		result += 1.5
		if daysRented > 3 {
			result += float64(daysRented-3) * 1.5
		}
	}
	return result
}
