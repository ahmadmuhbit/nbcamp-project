package repositories

import "nbcamp_project/server/models"

type MenuRepository interface {
	FindAll() (*[]models.Menu, error)
	CreateMenu(menu *models.Menu) error
}
