package repositories

import (
	"database/sql"
	"nbcamp_project/server/models"
)

type menuRepo struct {
	DB *sql.DB
}

func NewMenuRepository(db *sql.DB) MenuRepository {
	return &menuRepo{
		DB: db,
	}
}

func (m *menuRepo) Save(menu *models.Menu) error {
	query := `
		INSERT INTO menus (
			id, name, category,
			description, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
		`

	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		menu.ID, menu.Name, menu.Category,
		menu.Description, menu.CreatedAt, menu.UpdatedAt,
	)

	return err
}

func (m *menuRepo) FindAll() (*[]models.Menu, error) {
	query := `
		SELECT
			id, COALESCE(name, ''), COALESCE(category, ''), COALESCE(description, ''), created_at, updated_at
		FROM
			menus
	`

	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var menus []models.Menu

	for rows.Next() {
		var menu models.Menu
		err := rows.Scan(
			&menu.ID, &menu.Name, &menu.Category,
			&menu.Description, &menu.CreatedAt, &menu.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		menus = append(menus, menu)
	}

	return &menus, nil
}

func (m *menuRepo) FindByID(id string) (*models.Menu, error) {
	query := `
		SELECT
			id, COALESCE(name, ''), COALESCE(category, ''), COALESCE(description, ''), created_at, updated_at
		FROM
			menus
		WHERE
			id = $1
		`

	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var menu models.Menu
	err = row.Scan(
		&menu.ID, &menu.Name, &menu.Category,
		&menu.Description, &menu.CreatedAt, &menu.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &menu, nil
}
