package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type password struct {
	password string `json:"password"`
}

func checkPassword(body *password) error {
	if body.password == "n4th4n43l" {
		return nil
	}
	if body.password != "n4th4n43l" {
		return fmt.Errorf("incorrect secret")
	}
	return fmt.Errorf("System error")

}

func Home(c *fiber.Ctx) error {
	// Render index template

	bearer := c.Cookies("token")
	_, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized

		}
		return nil, nil
	})
	if err != nil {
		return c.Render("login", nil)
	}

	return c.Render("info", nil)
}

func Login(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body := new(password)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	err := checkPassword(body)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	c.Status(fiber.StatusOK)
	return nil
}
