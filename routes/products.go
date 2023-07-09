package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/prranavv/fiber_1/database"
	"github.com/prranavv/fiber_1/models"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productmodel models.Product) Product {
	return Product{
		ID:           productmodel.ID,
		Name:         productmodel.Name,
		SerialNumber: productmodel.SerialNumber,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.Database.Db.Find(&products)
	responseProducts := []Product{}
	for _, product := range products {
		responseproduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseproduct)
	}
	return c.Status(200).JSON(responseProducts)
}

func findproduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id=?", id)
	if product.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product
	if err != nil {
		return c.Status(400).JSON("Please ensure that id is an integer")
	}
	if err := findproduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseproduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseproduct)
}

func Updateproduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product
	if err != nil {
		return c.Status(400).JSON("Please ensure that id is an integer")
	}
	if err := findproduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}
	var updatedata UpdateProduct
	if err := c.BodyParser(&updatedata); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	product.Name = updatedata.Name
	product.SerialNumber = updatedata.SerialNumber
	database.Database.Db.Save(&product)
	responseproduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseproduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product
	if err != nil {
		return c.Status(400).JSON("Please ensure that id is an integer")
	}
	if err := findproduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).SendString("Successfully deleted")
}
