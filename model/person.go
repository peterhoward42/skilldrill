package model

type person struct {
	uid   int64
	email string
}

func newPerson(uid int64, email string) *person {
	return &person{
		uid:   uid,
		email: email,
	}
}
