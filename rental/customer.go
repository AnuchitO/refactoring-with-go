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
	result := fmt.Sprintf("%v%v%v", "Rental Record for ", c.Name(), "\n")
	for _, each := range c.rentals {
		thisAmount := 0.0
		switch each.Movie().PriceCode() {
		case REGULAR:
			thisAmount += 2
			if each.DaysRented() > 2 {
				thisAmount += float64(each.DaysRented()-2) * 1.5
			}
		case NEW_RELEASE:
			thisAmount += float64(each.DaysRented()) * 3.0
		case CHILDRENS:
			thisAmount += 1.5
			if each.DaysRented() > 3 {
				thisAmount += float64(each.DaysRented()-3) * 1.5
			}
		}
		frequentRenterPoints++
		if each.Movie().PriceCode() == NEW_RELEASE && each.DaysRented() > 1 {
			frequentRenterPoints++
		}
		result += fmt.Sprintf("%v%v%v%.1f%v", "\t", each.Movie().Title(), "\t", thisAmount, "\n")
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("%v%.1f%v", "Amount owed is ", totalAmount, "\n")
	result += fmt.Sprintf("%v%v%v", "You earned ", frequentRenterPoints, " frequent renter points")
	return result
}
