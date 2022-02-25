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
	totalAmount := 0.0
	frequentRenterPoints := 0
	result := fmt.Sprintf("Rental Record for %v\n", c.Name())
	for _, each := range c.rentals {
		thisAmount := each.amountFor()
		frequentRenterPoints++
		if each.Movie().PriceCode() == NEW_RELEASE && each.DaysRented() > 1 {
			frequentRenterPoints++
		}
		result += fmt.Sprintf("\t%v\t%.1f\n", each.Movie().Title(), thisAmount)
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", frequentRenterPoints)
	return result
}

func (r Rental) amountFor() (result float64) {
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
