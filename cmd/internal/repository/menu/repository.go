package menu

import (
	"github.com/ArifMun/go-rest-api.git/cmd/internal/model"
)

type Repository interface {
	GetMenu(MenuType string) ([]model.MenuItem, error)
}
