package models

import "errors"

type CreateUserRequest struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
	Email    string `json:"email"`
}

func (cur *CreateUserRequest) Validate() error {
	if cur.Username == "" || cur.Passwd == "" || cur.Email == "" {
		return errors.New("missing required user fields")
	}
	return nil
}

type LoginUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"passwd"`
}

func (lur *LoginUserRequest) Validate() error {
	if lur.Login == "" || lur.Password == "" {
		return errors.New("bad request")
	}
	return nil
}
