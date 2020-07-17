package controller

import (
	"github.com/hsmtkk/clean_arch_study_2/pkg/domain"
	"github.com/hsmtkk/clean_arch_study_2/pkg/usecase"
)

type Controller interface {
	CreateUser(u domain.User) error
	GetUsers() ([]domain.User, error)
}

type controllerImpl struct {
	useCase usecase.UseCase
}

func New(useCase usecase.UseCase) Controller {
	return &controllerImpl{useCase: useCase}
}

func (ctrl *controllerImpl) CreateUser(u domain.User) error {
	return ctrl.useCase.Create(u)
}

func (ctrl *controllerImpl) GetUsers() ([]domain.User, error) {
	return ctrl.useCase.Get([]domain.User{})
}
