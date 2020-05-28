package project

import (
	"github.com/go-playground/validator/v10"
)

type Project struct {
	Domain string `validate:"hostname"`
}

func (obj *Project) setDomain(s string) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		return err
	}
	obj.Domain = s
	return nil
}
