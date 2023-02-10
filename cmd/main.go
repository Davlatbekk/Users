package main

import (
	"app/controller"
	"app/models"
	"fmt"
)

func main() {

	controller.CreateUser(
		models.User{
			Id:      1,
			Name:    "Shohruh",
			Surname: "Safarov",
		},
	)

	controller.CreateUser(
		models.User{
			Id:      2,
			Name:    "Abduqodir",
			Surname: "Musayev",
		},
	)

	fmt.Println(controller.GetByIdUser(2))
	fmt.Println(controller.GetByIdUser(1))
	// controller.DeleteUser(models.User{Id: 1,Name: "q",Surname: "QWER"})
	// fmt.Println(controller.GetByIdUser(1))

	// for _, user := range controller.GetListUser() {
	// 	fmt.Println(user)
	// }
}
