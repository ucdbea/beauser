package models

type User struct {
	FirstName string
	LastName  string
	UserName  string
	Year      string
	Email     string
	Major     string
	Schedule  *Schedule
}
