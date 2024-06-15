package handlers

import (
    "context"
    "time"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"

    "crm/database"
    "crm/models" 
)



func GetCustomers(c *fiber.Ctx) error {
	var customers []models.Customer 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.Collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var customer models.Customer 
		if err := cursor.Decode(&customer); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		customers = append(customers, customer)
	}

	if err := cursor.Err(); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(customers)
}

func GetCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	var customer models.Customer 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = database.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&customer)
	if err != nil {
		return c.Status(404).SendString("Customer not found")
	}

	return c.JSON(customer)
}

func CreateCustomer(c *fiber.Ctx) error {
	customer := new(models.Customer) 
	if err := c.BodyParser(customer); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	customer.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := database.Collection.InsertOne(ctx, customer)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(customer)
}

func UpdateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	customer := new(models.Customer) 
	if err := c.BodyParser(customer); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": customer,
	}
	_, err = database.Collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(customer)
}

func DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = database.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("Customer successfully deleted")
}
