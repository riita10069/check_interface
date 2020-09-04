package repository

var repos *AllRepository

type AllRepository interface {
	Create() error
}
