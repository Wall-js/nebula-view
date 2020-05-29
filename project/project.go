package project

import (
	"github.com/Wall-js/nebula-core/backend"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Project struct {
	Domain  string `validate:"hostname"`
	Root    string
	backend backend.Backend
}

func NewProject() *Project {
	return &Project{
		Root:    "",
		backend: backend.NewGinBackend(),
	}
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

func (obj *Project) Router() http.Handler {
	return obj.Router()
}

func (obj *Project) Run() error {
	return obj.backend.Run(obj.Root)
}
