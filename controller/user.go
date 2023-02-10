package controller

import (
	"app/models"
	"fmt"
)

var Users []models.User

func CreateUser(data models.User) {
	Users = append(Users, data)
}

func GetListUser() []models.User {
	return Users
}

func GetByIdUser(id int) models.User {
	for _, val := range Users {
		if val.Id == id {
			return val
		}
	}
	return models.User{}
}

func UpdateUser(data models.User) {
	for in, val := range Users {
		if val.Id == data.Id {
			Users[in].Name = data.Name
			Users[in].Surname = data.Surname
			return
		}
	}
	fmt.Println("Bunday User yo'q")
}

func DeleteUser(data models.User)  {
	for in, val := range Users {
		if val.Id == data.Id {
			Users = append(Users[:in], Users[in+1:]...)
			return 
		}
	}
	fmt.Println("Bunday User yo'q")
}

// getbyid
// update
// delete
