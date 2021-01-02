package models

type Schedule struct {
	quarter string
	classes *Class
	user    *User
}
