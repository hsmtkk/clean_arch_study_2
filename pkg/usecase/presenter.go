package usecase

import "github.com/hsmtkk/clean_arch_study_2/pkg/domain"

type Presenter interface {
	Response(us []domain.User) []domain.User
}
