package domain

type User struct {
	Id uint64
	FirstName string
	LastName string
	Email string
}

var morning = "Good Morning"
var Morning = "Hey, " + morning
