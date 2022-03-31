package utility

type UserInput struct {
	Name             string `json:"name" validate:"required"`
	Username         string `json:"username" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required,min=8" `
	Confirm_password string `json:"confirm_password" validate:"required,eqfield=Password"`
}
type UserSigninInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
