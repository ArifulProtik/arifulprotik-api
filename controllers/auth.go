package controllers

import (
	"context"
	"github.com/ArifulProtik/arifulprotik-api/auth"
	"github.com/ArifulProtik/arifulprotik-api/db"
	"github.com/ArifulProtik/arifulprotik-api/services"
	"github.com/ArifulProtik/arifulprotik-api/utility"
	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {
	c.Accepts("json", "text")
	userinput := &utility.UserInput{}

	if err := c.BodyParser(userinput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}
	validator := utility.NewValidator()
	err := validator.Struct(userinput)
	if err != nil {
		response := utility.ToErrResponse(err)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	user, err := services.CreateUser(db.Ent, userinput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Username or Email Already Exist",
		})
	}
	token, err := auth.GenAccessToken(context.Background(), user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}
func Signin(c *fiber.Ctx) error {
	c.Accepts("json", "text")
	signinInput := &utility.UserSigninInput{}
	if err := c.BodyParser(signinInput); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Error": "Input Valid JSON"})
	}
	validator := utility.NewValidator()
	err2 := validator.Struct(signinInput)
	if err2 != nil {
		response := utility.ToErrResponse(err2)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	user, err := services.GetUserByEmail(db.Ent, signinInput.Email)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"Error": "invalid email or password"})
	}
	err = utility.VerifyPass(user.Password, signinInput.Password)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"Error": "invalid email or password"})
	}
	token, err := auth.GenAccessToken(context.Background(), user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Internal Server Error"})
	}
	return c.JSON(fiber.Map{"token": token, "user": user})
}
