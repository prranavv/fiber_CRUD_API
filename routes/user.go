package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/prranavv/fiber_1/database"
	"github.com/prranavv/fiber_1/models"
)

type User struct {
	//this is not the model User, see this as a serialzer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(usermodel models.User) User {
	return User{
		ID:        usermodel.ID,
		FirstName: usermodel.FirstName,
		LastName:  usermodel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)
	responseusers := []User{}
	for _, user := range users {
		responduser := CreateResponseUser(user)
		responseusers = append(responseusers, responduser)
	}
	return c.Status(200).JSON(responseusers)
}

func finduser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id=?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func Getuser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON("Please ensure that id is an integer")
	}
	if err := finduser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseuser := CreateResponseUser(user)
	return c.Status(200).JSON(responseuser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON("Please ensure that id is an integer")
	}
	if err := finduser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	var updatedata UpdateUser
	if err := c.BodyParser(&updatedata); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	user.FirstName = updatedata.FirstName
	user.LastName = updatedata.LastName
	database.Database.Db.Save(&user)
	responseuser := CreateResponseUser(user)
	return c.Status(200).JSON(responseuser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON("Please ensure that id is an integer")
	}
	if err := finduser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).SendString("Successfully deleted")
}
