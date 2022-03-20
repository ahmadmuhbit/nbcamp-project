package params

import (
	"github.com/google/uuid"
	"nbcamp_project/server/models"
	"time"
)

type MenuCreate struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

func (m *MenuCreate) ParseToModel() *models.Menu {
	menu := models.NewMenu()
	menu.Category = m.Category
	menu.Name = m.Name
	menu.Description = m.Description
	return menu
}

type MenuUpdate struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
}

func (m *MenuUpdate) ParseToModel() *models.Menu {
	return &models.Menu{
		ID:          m.ID,
		Name:        m.Name,
		Category:    m.Category,
		Description: m.Description,
		UpdatedAt:   time.Now(),
	}
}
