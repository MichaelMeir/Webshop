package models

type User struct {
	Username string
	Password string
	Email string

	Session_ip string
	Session_id string

	Admin bool

	Address string
	ZipCode string
	Province string
	Country string
}
