package routes

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prranavv/fiber_1/database"
	"github.com/prranavv/fiber_1/models"
)

type Order struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{
		ID:        order.ID,
		User:      user,
		Product:   product,
		CreatedAt: order.CreatedAt,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	if err := finduser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var product models.Product
	if err := findproduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&order)
	responseuser := CreateResponseUser(user)
	responseproduct := CreateResponseProduct(product)
	responseorder := CreateResponseOrder(order, responseuser, responseproduct)
	return c.Status(200).JSON(responseorder)
}
func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	responseorders := []Order{}

	for _, order := range orders {
		var user models.User
		var product models.Product
		database.Database.Db.Find(&user, "id=?", order.UserRefer)
		database.Database.Db.Find(&product, "id=?", order.ProductRefer)
		responseorder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))
		responseorders = append(responseorders, responseorder)
	}
	return c.Status(200).JSON(responseorders)
}

func findorder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id=?", id)
	if order.ID == 0 {
		return errors.New("Order does not exist")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var order models.Order
	if err != nil {
		return c.Status(400).JSON("Please ensure id is an integer")
	}
	if err := findorder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	var product models.Product

	database.Database.Db.First(&user, order.UserRefer)
	database.Database.Db.First(&product, order.ProductRefer)
	responseuser := CreateResponseUser(user)
	responseproduct := CreateResponseProduct(product)
	responseorder := CreateResponseOrder(order, responseuser, responseproduct)
	return c.Status(200).JSON(responseorder)
}
