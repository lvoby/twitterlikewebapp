package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"

) // use empty import because without them I have err with import in user.go; why?


type Tweet struct {
	Id       int    `json:"id"`
	Message  string `json:"message"`
	PostTime time.Time `json:"post_time"`
	UserId   int `json:"user_id"`
}

// Validate User fields
func (t *Tweet) Validate() error {
	return validation.ValidateStruct(
		t,
		validation.Field(&t.Message, validation.Required, validation.Length(1, 280)),
	)
}