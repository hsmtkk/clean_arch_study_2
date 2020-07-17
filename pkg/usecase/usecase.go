package usecase

import (
	"github.com/hsmtkk/clean_arch_study_2/pkg/domain"
)

type UseCase interface {
	Create(u domain.User) error
	Get(us []domain.User) ([]domain.User, error)
}

type useCaseImpl struct {
	repository Repository
	presenter  Presenter
}

func New(repository Repository, presenter Presenter) UseCase {
	return &useCaseImpl{repository: repository, presenter: presenter}
}

func (useCase *useCaseImpl) Create(u domain.User) error {
	return useCase.repository.Store(u)
}

func (useCase *useCaseImpl) Get(us []domain.User) ([]domain.User, error) {
	us2, err := useCase.repository.FindAll(us)
	if err != nil {
		return nil, err
	}
	return useCase.presenter.Response(us2), nil
}
