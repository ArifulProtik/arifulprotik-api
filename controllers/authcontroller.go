package controllers

import (
	"context"
	"errors"
	"fmt"
	authservice "github.com/ArifulProtik/arifulprotik-api/auth"
	"github.com/ArifulProtik/arifulprotik-api/config"
	"github.com/ArifulProtik/arifulprotik-api/db"
	"github.com/ArifulProtik/arifulprotik-api/services"
	"github.com/ArifulProtik/arifulprotik-api/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	"time"
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
			"error": "Username or Email Already Exist",
		})
	}
	rf_uuid := uuid.New()
	token, err := authservice.GenAccessToken(context.Background(), user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	exp := time.Now().Add(time.Hour * 24 * 30).Unix()
	refreshToken, err := authservice.GenRefreshToken(context.Background(), user.ID, rf_uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}
	err = services.CreateAuth(db.Ent, user.ID, rf_uuid, refreshToken, exp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"access_token":  token,
		"refresh_token": refreshToken,
		"user":          user,
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
	token, err := authservice.GenAccessToken(context.Background(), user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Internal Server Error"})
	}
	rf_uuid := uuid.New()
	refreshToken, err := authservice.GenRefreshToken(context.Background(), user.ID, rf_uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}
	exp := time.Now().Add(time.Hour * 24 * 30).Unix()
	auth, err := services.GetAuthByUserID(db.Ent, user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	if auth != nil {
		err := services.DeleteAuth(db.Ent, user.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		}
		err = services.CreateAuth(db.Ent, user.ID, rf_uuid, refreshToken, exp)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		}
		return c.JSON(fiber.Map{"access_token": token, "refresh_token": refreshToken, "user": user})

	}
	err = services.CreateAuth(db.Ent, user.ID, rf_uuid, refreshToken, exp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	return c.JSON(fiber.Map{"access_token": token, "refresh_token": refreshToken, "user": user})
}

func Refresh(c *fiber.Ctx) error {
	type Refresh struct {
		Token string `json:"token"`
	}
	var token = &Refresh{}
	if err := c.BodyParser(token); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Error": "Input Valid JSON"})
	}
	if token.Token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "provide a valid token"})
	}
	maptoken, err := jwt.Parse(token.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("InValid Token")
		}
		return []byte(config.C.Jwt_Key), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "provide a valid token"})
	}
	if !maptoken.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token Has Expired Please Login Again"})
	}
	data, _ := maptoken.Claims.(jwt.MapClaims)
	rfuuid, err := uuid.Parse(fmt.Sprint(data["RF_UUID"]))
	if err != nil {
		fmt.Println(err)
		return err
	}
	auth, err := services.GetAuthByRfId(db.Ent, rfuuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": " Internal Server Error"})

	}
	if auth == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "login Again"})
	}
	access_token, _ := authservice.GenAccessToken(context.Background(), auth.Userid)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"access_token": access_token})
}
