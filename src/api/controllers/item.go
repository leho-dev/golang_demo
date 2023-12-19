package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lehodev/golang-server/src/api/models"
	"github.com/lehodev/golang-server/src/configs/db/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAll(c *fiber.Ctx) error {
	todos := mongodb.DB.Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := todos.Find(ctx, bson.M{}, options.Find())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Todos not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting todos",
			"error":   err.Error(),
		})
	}

	var todoList []models.Item
	if err := cursor.All(ctx, &todoList); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting todos",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data":  todoList,
		"total": len(todoList),
	})
}

func GetById(c *fiber.Ctx) error {
	todos := mongodb.DB.Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(c.Params("id"))
	options := options.FindOne()
	filter := bson.D{{"_id", id}}

	var item models.Item
	if err := todos.FindOne(ctx, filter, options).Decode(&item); err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Todo not found",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting todo",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": item,
	})
}

func Create(c *fiber.Ctx) error {
	todos := mongodb.DB.Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing body",
			"error":   err.Error(),
		})
	}

	result, err := todos.InsertOne(ctx, item)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating todo",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"data": result,
	})
}
