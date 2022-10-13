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

type Pay struct {
	KeySender string `json:"sender"`
	Recipient string `json:"recipient"`
	ValuePay  int64  `json:"value"`
	IdUser    int    `json:"id_user"`
	IdGame    int    `json:"id_game"`
}

type Payment struct {
	KeySender string
	Recipient string
	ValuePay  int64
}

type Order struct {
	IdUser int    `json:"id_user"`
	IdGame int    `json:"id_game"`
	Erro   string `json:"erro"`
}
