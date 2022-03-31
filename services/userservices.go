package services

import (
	"context"
	"github.com/ArifulProtik/arifulprotik-api/ent"
	"github.com/ArifulProtik/arifulprotik-api/ent/user"
	"github.com/ArifulProtik/arifulprotik-api/utility"
	"github.com/google/uuid"
)

func CreateUser(client *ent.Client, input *utility.UserInput) (*ent.User, error) {

	newpass, _ := utility.HashBeforeSave(input.Password)
	user, err := client.Debug().User.Create().SetName(input.Name).SetUsername(input.Username).SetEmail(input.Email).SetPassword(string(newpass)).Save(context.Background())
	if err != nil {
		return &ent.User{}, err
	}
	return user, nil
}
func GetUserByid(client *ent.Client, ID uuid.UUID) (*ent.User, error) {
	user, err := client.User.Query().Where(user.IDEQ(ID)).First(context.Background())
	if err != nil {
		return &ent.User{}, err
	}
	return user, nil
}
func GetUserByEmail(client *ent.Client, email string) (*ent.User, error) {
	user, err := client.User.Query().Where(user.EmailEQ(email)).First(context.Background())
	if err != nil {
		return &ent.User{}, err
	}
	return user, nil
}
