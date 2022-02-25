package rental

import "testing"

type CustomerTest struct {
}

func NewCustomerTest() (rcvr *CustomerTest) {
	rcvr = &CustomerTest{}
	return
}
func TestCustomer(t *testing.T) {
	customer := NewCustomer("AnuchitO")
	customer.rentals = append(customer.rentals, NewRental(NewMovie("Kingsman", REGULAR), 2))
	customer.rentals = append(customer.rentals, NewRental(NewMovie("Iron Man", REGULAR), 3))
	customer.rentals = append(customer.rentals, NewRental(NewMovie("The Avengers", NEW_RELEASE), 1))
	customer.rentals = append(customer.rentals, NewRental(NewMovie("Shang-chi", NEW_RELEASE), 2))
	customer.rentals = append(customer.rentals, NewRental(NewMovie("Ant-Man", CHILDRENS), 3))
	customer.rentals = append(customer.rentals, NewRental(NewMovie("The Batman", CHILDRENS), 4))

	want := `Rental Record for AnuchitO
	Kingsman	2.0
	Iron Man	3.5
	The Avengers	3.0
	Shang-chi	6.0
	Ant-Man	1.5
	The Batman	3.0
Amount owed is 19.0
You earned 7 frequent renter points`

	if want != customer.Statement() {
		t.Errorf("Expect \n%v\n, got \n%v", want, customer.Statement())
	}
}
