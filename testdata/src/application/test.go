package application

import "github.com/riita10069/check_interface/testdata/src/application/repository"

type Test struct {
	repos *repository.AllRepository
}

func NewTest(repos *repository.AllRepository) *Test {
	return &Test{
		repos: repos,
	}
}

func (a *Test) Create() error {
	return nil
}
