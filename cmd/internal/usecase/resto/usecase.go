package resto

import "github.com/ArifMun/go-rest-api.git/cmd/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
