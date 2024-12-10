package service

type Service interface {
}

func New() *service {
	return &service{}
}

type service struct{}
