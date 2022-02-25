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
		charge := r.Movie().getCharge(r.daysRented)
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
		charge := r.Movie().getCharge(r.daysRented)
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
		result += r.Movie().getCharge(r.daysRented)
	}
	return result
}

func (r Rental) getFrequentRenterPoints() int {
	if r.Movie().PriceCode() == NEW_RELEASE && r.DaysRented() > 1 {
		return 2
	}
	return 1
}

func (m Movie) getCharge(daysRented int) (result float64) {
	price := NewPrice(m)
	return price.getCharge(daysRented)
}

type PriceState interface {
	getCharge(daysRented int) float64
	Next() PriceState
}

func NewPrice(m Movie) PriceState {
	switch m.PriceCode() {
	case REGULAR:
		return RegularPrice{Movie: m}
	case CHILDRENS:
		return ChildrensPrice{Movie: m}
	case NEW_RELEASE:
		return NewReleasePrice{Movie: m}
	default:
		return RegularPrice{Movie: m}
	}
}

type RegularPrice struct {
	Movie Movie
}

func (p RegularPrice) getCharge(daysRented int) (result float64) {
	result += 2
	if daysRented > 2 {
		result += float64(daysRented-2) * 1.5
	}
	return
}

func (p RegularPrice) Next() PriceState {
	return RegularPrice{Movie: p.Movie}
}

type ChildrensPrice struct {
	Movie Movie
}

func (p ChildrensPrice) getCharge(daysRented int) (result float64) {
	result += 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * 1.5
	}
	return
}

func (p ChildrensPrice) Next() PriceState {
	return RegularPrice{Movie: p.Movie}
}

type NewReleasePrice struct {
	Movie Movie
}

func (p NewReleasePrice) getCharge(daysRented int) (result float64) {
	return float64(daysRented) * 3.0
}

func (p NewReleasePrice) Next() PriceState {
	return RegularPrice{Movie: p.Movie}
}
