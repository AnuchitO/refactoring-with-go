package rental

import "fmt"

type Customer struct {
	_name    string
	_rentals []Rental
}

func NewCustomer(name string) Customer {
	return Customer{
		_name:    name,
		_rentals: []Rental{},
	}
}
func (rcvr Customer) AddRental(arg Rental) Customer {
	rcvr._rentals = append(rcvr._rentals, arg)
	return rcvr
}
func (rcvr Customer) GetName() string {
	return rcvr._name
}
func (rcvr Customer) Statement() string {
	totalAmount := 0.0
	frequentRenterPoints := 0
	result := fmt.Sprintf("%v%v%v", "Rental Record for ", rcvr.GetName(), "\n")
	for _, each := range rcvr._rentals {
		thisAmount := 0.0
		switch each.GetMovie().PriceCode() {
		case REGULAR:
			thisAmount += 2
			if each.GetDaysRented() > 2 {
				thisAmount += float64(each.GetDaysRented()-2) * 1.5
			}
		case NEW_RELEASE:
			thisAmount += float64(each.GetDaysRented()) * 3.0
		case CHILDRENS:
			thisAmount += 1.5
			if each.GetDaysRented() > 3 {
				thisAmount += float64(each.GetDaysRented()-3) * 1.5
			}
		}
		frequentRenterPoints++
		if each.GetMovie().PriceCode() == NEW_RELEASE && each.GetDaysRented() > 1 {
			frequentRenterPoints++
		}
		result += fmt.Sprintf("%v%v%v%.1f%v", "\t", each.GetMovie().Title(), "\t", thisAmount, "\n")
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("%v%.1f%v", "Amount owed is ", totalAmount, "\n")
	result += fmt.Sprintf("%v%v%v", "You earned ", frequentRenterPoints, " frequent renter points")
	return result
}
