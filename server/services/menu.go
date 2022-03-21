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
		MenuRepository: repository,
		DB:             db,
	}
}

func (m *MenuServices) CreateNewMenu(request *params.MenuCreate) (bool, error) {

	menu := request.ParseToModel()

	err := m.MenuRepository.Save(menu)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (m *MenuServices) GetAllMenu() (*[]params.MenuSingleView, error) {
	menus, err := m.MenuRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return makeMenuListView(menus), nil
}

func makeMenuListView(models *[]models.Menu) *[]params.MenuSingleView {
	var menuListView []params.MenuSingleView
	for _, model := range *models {
		menuListView = append(menuListView, *makeMenuSingleView(&model))
	}

	return &menuListView
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

func (m *MenuServices) GetMenuByID(id string) (*params.MenuSingleView, error) {
	menu, err := m.MenuRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return makeMenuSingleView(menu), nil
}

func (m *MenuServices) UpdateMenuByID(request *params.MenuUpdate) (bool, error) {
	model := request.ParseToModel()
	err := m.MenuRepository.UpdateByID(model)
	if err != nil {
		return false, err
	}

	return true, nil
}
