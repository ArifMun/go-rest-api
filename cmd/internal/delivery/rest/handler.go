package rest

import "github.com/ArifMun/go-rest-api.git/cmd/internal/usecase/resto"

type handler struct {
	restoUsecase resto.Usecase
}

func NewHandler(restoUscase resto.Usecase) *handler {
	return &handler{
		restoUsecase: restoUscase,
	}
}
