package model

/*
The person type models a person in terms of the user name part of their email
address and the UID for that person.  The design intent is that none of Api
fields are exported, but the reason that some are, is solely to facilitate
automated serialization by yaml.Marshal().
*/
type person struct {
	Email string
}

// Compulsory constructor.
func newPerson(email string) *person {
	return &person{
		Email: email,
	}
}
