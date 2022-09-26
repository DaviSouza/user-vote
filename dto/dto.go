package dto

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	DateBirth time.Time `json:"dateBirth"`
	Address   Address   `json:"address"`
}

type Address struct {
	Street  string `json:"street"`
	Number  string `json:"number"`
	County  string `json:"county"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}
