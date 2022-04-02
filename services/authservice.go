package services

import (
	"context"
	"github.com/ArifulProtik/arifulprotik-api/ent"
	"github.com/ArifulProtik/arifulprotik-api/ent/auth"
	"github.com/google/uuid"
)

func CreateAuth(client *ent.Client, id uuid.UUID, rf_id uuid.UUID, token string, ex_time int64) error {
	_, err := client.Auth.Create().SetRfUUID(rf_id).SetUserid(id).SetRfToken(token).SetExpireAt(ex_time).Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func GetAuthByRfId(client *ent.Client, rfid uuid.UUID) (*ent.Auth, error) {
	auth, err := client.Auth.Query().Where(auth.RfUUIDEQ(rfid)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return auth, nil
}
func GetAuthByUserID(client *ent.Client, userid uuid.UUID) (*ent.Auth, error) {
	auth, err := client.Auth.Query().Where(auth.UseridEQ(userid)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return auth, nil
}
func DeleteAuth(client *ent.Client, userid uuid.UUID) error {
	_, err := client.Auth.Delete().Where(auth.UseridEQ(userid)).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
