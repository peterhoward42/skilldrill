package model

type Person struct {
	email string
}

func NewPerson(email string) *Person {
	return &Person{email: email}
}
