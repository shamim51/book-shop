package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrResponse struct {
	Errors []string `json:"errors"`
}

func ToErrResponse(err error) *ErrResponse {
	if fieldErrors, ok := err.(validator.ValidationErrors); ok {

		resp := ErrResponse{
			Errors: make([]string, len(fieldErrors)),
		}

		for i, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				resp.Errors[i] = fmt.Sprintf("%s is a required field", err.Field())

			case "max":
				resp.Errors[i] = fmt.Sprintf("length of %s can't exceed than %s.", err.Field(), err.Param())

			case "url":
				resp.Errors[i] = fmt.Sprintf("%s must be a valid url", err.Field())

			case "alphaspace":
				resp.Errors[i] = fmt.Sprintf("%s can only contain alphabetic and space chracters", err.Field())

			case "datetime":
				resp.Errors[i] = fmt.Sprintf("%s must follow %s format", err.Field(), err.Param())

			default:
				resp.Errors[i] = fmt.Sprintf("something went wrong one %s; %s", err.Field(), err.Tag())

			}
		}
		return &resp
	}
	return nil
}
