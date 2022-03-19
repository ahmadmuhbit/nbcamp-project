package services

import (
	"database/sql"
	"nbcamp_project/server/helper"
	"nbcamp_project/server/models"
	params "nbcamp_project/server/params/employee"
	repositories "nbcamp_project/server/repositories/employee"
	"time"
)

type EmployeeServices struct {
	EmployeeRepository repositories.EmployeeRepository
	DB                 *sql.DB
}

func NewEmployeeService(db *sql.DB) *EmployeeServices {
	repositories := repositories.NewEmployeeRepository(db)
	return &EmployeeServices{
		DB:                 db,
		EmployeeRepository: repositories,
	}
}

func (e *EmployeeServices) CreateNewEmployee(request *params.EmployeeCreate) bool {
	defer helper.HandleError()
	emp := request.ParseToModel()
	err := e.EmployeeRepository.Save(emp)
	if err != nil {
		helper.HandlePanicIfError(err)
		return false
	}
	return true
}

func (e *EmployeeServices) GetAllEmployees() *[]params.EmployeeSingleView {
	defer helper.HandleError()
	employees, err := e.EmployeeRepository.FindAll()
	helper.HandlePanicIfError(err)
	return makeEmployeeListView(employees)
}

func makeEmployeeListView(models *[]models.Employee) *[]params.EmployeeSingleView {
	var params []params.EmployeeSingleView
	for _, model := range *models {
		params = append(params, *makeEmployeeSingleView(&model))
	}
	return &params
}

func makeEmployeeSingleView(models *models.Employee) *params.EmployeeSingleView {
	return &params.EmployeeSingleView{
		ID:        models.ID,
		NIP:       models.NIP,
		Name:      models.Name,
		Address:   models.Address,
		CreatedAt: models.CreatedAt.Format(time.RFC3339),
		UpdatedAt: models.UpdatedAt.Format(time.RFC3339),
	}
}

func (e *EmployeeServices) GetEmployeeByID(id string) *params.EmployeeSingleView {
	defer helper.HandleError()
	employee, err := e.EmployeeRepository.FindByID(id)
	helper.HandlePanicIfError(err)

	return makeEmployeeSingleView(employee)
}

func (e *EmployeeServices) UpdateByID(request *params.EmployeeUpdate) bool {
	defer helper.HandleError()

	model := request.ParseToModel()

	err := e.EmployeeRepository.UpdateByID(model)
	helper.HandlePanicIfError(err)

	return true
}

func (e *EmployeeServices) DeleteEmployeeByID(id string) bool {
	defer helper.HandleError()
	err := e.EmployeeRepository.DeleteByID(id)
	if err != nil {
		helper.HandlePanicIfError(err)
		return false
	}
	return true
}
