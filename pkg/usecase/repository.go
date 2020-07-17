package usecase

import "github.com/hsmtkk/clean_arch_study_2/pkg/domain"

type Repository interface {
	Store(u domain.User) error
	FindAll(us []domain.User) ([]domain.User, error)
}
