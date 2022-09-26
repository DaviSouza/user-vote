package service

import (
	"fmt"
	"log"
	"user-vote/domain"
	"user-vote/dto"
	"user-vote/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllUsers(db *mongo.Database) []*dto.User {
	var usersDto []*dto.User
	users, err := repository.GetAllUsers(db)

	if err != nil {
		fmt.Println(err.Error())
	}

	for index, element := range users {
		dto := &dto.User{}
		dto = convertDomainToDto(*element, dto)
		usersDto[index] = dto
	}

	return usersDto
}

func CreateUser(userDto dto.User, db *mongo.Database) error {
	log.Printf("Create collection user Doma")
	//fmt.Printf("%+v", userDto)
	user := newUser(userDto)
	//fmt.Printf("%+v", user)
	err := repository.CreateUser(user, db)
	if err != nil {
		return err
	}
	return err
}

func newUser(userDto dto.User) *domain.User {
	u := domain.NewUser()
	return convertDtoToDomain(userDto, u)
}

func convertDtoToDomain(userDto dto.User, u *domain.User) *domain.User {
	u.Name = userDto.Name
	u.Address.City = userDto.Address.City
	u.Address.County = userDto.Address.County
	u.Address.Number = userDto.Address.Number
	u.Address.State = userDto.Address.State
	u.Address.Street = userDto.Address.Street
	u.Address.ZipCode = userDto.Address.ZipCode
	return u
}

func convertDomainToDto(user domain.User, u *dto.User) *dto.User {
	u.Name = user.Name
	u.Address.City = user.Address.City
	u.Address.County = user.Address.County
	u.Address.Number = user.Address.Number
	u.Address.State = user.Address.State
	u.Address.Street = user.Address.Street
	u.Address.ZipCode = user.Address.ZipCode
	return u
}

func UpdateUser(userDto dto.User, db *mongo.Database) error {
	user := &domain.User{}
	user = convertDtoToDomain(userDto, user)
	err := repository.UpdateUser(user, db)
	if err != nil {
		return err
	}
	return err

}

func DeleteUser(id string, db *mongo.Database) error {
	err := repository.DeleteUser(id, db)
	if err != nil {
		return err
	}
	return err
}
