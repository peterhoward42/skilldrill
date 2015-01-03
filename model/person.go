package model

// The person type models a person in terms of the user name part of their email
// address and the UID for that person.
type person struct {
	uid   int64
	email string
}

// Compulsory constructor.
func newPerson(uid int64, email string) *person {
	return &person{
		uid:   uid,
		email: email,
	}
}
