package utils

import "github.com/gofiber/fiber/v2"


type Respone struct {
	Status string `json:"status"`
	ResponeCode int64 `json:"respone_code"`
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

func Success(c *fiber.Ctx ,message string, data interface{}) error{
	return c.Status(fiber.StatusOK).JSON(Respone{
		Status: "Success",
		ResponeCode: fiber.StatusOK,
		Message: message,
		Data: data,
	})
}

func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Respone{
		Status: "Created",
		ResponeCode: fiber.StatusCreated,
		Message: message,
		Data: data,
	})
}

func BadRequest(c *fiber.Ctx, message string, err string) error {
	return c.Status(fiber.StatusBadRequest).JSON(Respone{
		Status: "Error Bad Request",
		ResponeCode: fiber.StatusCreated,
		Message: message,
		Error: err,
	})
}

func NotFound(c *fiber.Ctx, message string,err string) error {
	return c.Status(fiber.StatusNotFound).JSON(Respone{
		Status: "Not Found",
		ResponeCode: fiber.StatusNotFound,
		Message: message,
		Error: err,
	})
}