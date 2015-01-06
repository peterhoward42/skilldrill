package model

// The person type models a person in terms of the user name part of their email
// address and the UID for that person.
type person struct {
	email string
}

// Compulsory constructor.
func newPerson(email string) *person {
	return &person{
		email: email,
	}
}
