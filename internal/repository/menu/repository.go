package menu

import (
	"github.com/dduuddeekk/go-restaurant-app/internal/model"
)

type Repository interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	GetMenu(orderCode string) ([]model.MenuItem, error)
}
