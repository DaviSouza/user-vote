package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	DateBirth time.Time          `json:"dateBirth,omitempty" bson:"dateBirth,omitempty"`
	Address   Address            `json:"address,omitempty" bson:"address,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

type Address struct {
	Street  string `json:"street,omitempty" bson:"street"`
	Number  string `json:"number,omitempty" bson:"number"`
	County  string `json:"county,omitempty" bson:"county"`
	City    string `json:"city,omitempty" bson:"city"`
	State   string `json:"state,omitempty" bson:"state"`
	ZipCode string `json:"zipCode,omitempty" bson:"zipCode"`
}

func NewUser() *User {
	u := &User{}
	u.ID = primitive.NewObjectID()
	u.CreatedAt = time.Now()
	return u
}
