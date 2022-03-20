package services

import (
	"database/sql"
	"nbcamp_project/server/models"
	params "nbcamp_project/server/params/menu"
	repositories "nbcamp_project/server/repositories/menu"
)

type MenuServices struct {
	MenuRepository repositories.MenuRepository
	DB             *sql.DB
}

func NewMenuService(db *sql.DB) *MenuServices {
	repository := repositories.NewMenuRepository(db)
	return &MenuServices{
		DB:             db,
		MenuRepository: repository,
	}
}

func (m *MenuServices) GetAllMenu() (*[]params.MenuSingleView, error) {
	menus, err := m.MenuRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return makeMenuListView(menus), nil
}

func makeMenuSingleView(models *models.Menu) *params.MenuSingleView {
	return &params.MenuSingleView{
		ID:          models.ID,
		Name:        models.Name,
		Category:    models.Category,
		Description: models.Description,
		CreatedAt:   models.CreatedAt,
		UpdatedAt:   models.UpdatedAt,
	}
}

func makeMenuListView(models *[]models.Menu) *[]params.MenuSingleView {
	var menuListView []params.MenuSingleView
	for _, model := range *models {
		menuListView = append(menuListView, *makeMenuSingleView(&model))
	}

	return &menuListView
}

func (m *MenuServices) CreateNewMenu(request *params.MenuCreate) (bool, error) {

	menu := request.ParseToModel()

	err := m.MenuRepository.CreateMenu(menu)
	if err != nil {
		return false, nil
	}

	return true, nil
}
