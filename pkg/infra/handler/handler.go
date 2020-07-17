package handler

import (
	"github.com/hsmtkk/clean_arch_study_2/pkg/interface/controller"
	"github.com/labstack/echo"
)

type Handler interface {
	CreateUser(ctx echo.Context) error
	GetUsers(ctx echo.Context) error
}

type handlerImpl struct {
	controller controller.Controller
}

func New(controller controller.Controller) Handler {
	return &handlerImpl{controller: controller}
}

func (handler *handlerImpl) CreateUser(ctx echo.Context) error {
	return nil
}

func (handler *handlerImpl) GetUsers(ctx echo.Context) error {
	return nil
}
