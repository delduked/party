package routes

import (
	"fmt"

	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Password struct {
	Password string `json:"Password"`
}

var (
	secret = []byte(`asdf908hj90fdsah908dsafh009q340937109f14f09hsd980fasdf980ahsdf0(SD)F(&*HSDF)(&709SD)F*&$@)(@&#$#F@)H&`)
)

func getPassword() (string, bool) {
	importPassword, ok := os.LookupEnv("PASSWORD")
	if !ok {
		return "n4th4n43l", false
	}
	return importPassword, true

}

func checkPassword(body *Password) error {

	environmentPassword, _ := getPassword()
	if body.Password == environmentPassword {
		return nil
	}
	if body.Password != environmentPassword {
		return fmt.Errorf("incorrect secret")
	}
	return fmt.Errorf("System error")

}

func Home(c *fiber.Ctx) error {
	// Render index template
	bearer := c.Cookies("token")
	token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("2")
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil {
		return c.Render("login", nil)
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Render("login", nil)
	}

	if err != nil {
		fmt.Println(err)
		return c.Render("login", nil)
	}
	return c.Render("info", nil)
}

func Login(c *fiber.Ctx) error {
	c.Accepts("application/json")

	body := new(Password)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	err := checkPassword(body)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(secret)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	// Generate new cookie for the clients request
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = t

	// Set the JWT inside the newly created cookie
	c.Cookie(cookie)

	c.Status(fiber.StatusOK)
	return nil
}
