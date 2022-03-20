package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"nbcamp_project/server/helper"
	params "nbcamp_project/server/params/menu"
	"nbcamp_project/server/services"
	"net/http"
)

type MenuController interface {
	FindAll(rw http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
}

type menuController struct {
	db *sql.DB
}

func NewMenuController(db *sql.DB) MenuController {
	return &menuController{db}
}

//func (*menuController) FindALl(rw http.ResponseWriter, r *http.Request) {
//
//}

func (m *menuController) FindAll(w http.ResponseWriter, r *http.Request) {
	menus, err := services.NewMenuService(m.db).GetAllMenu()

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			helper.HandleNotFound(w, errors.New("No Data"))
		} else {
			helper.HandleInternalServerError(w, err)
		}
		return
	}

	if len(*menus) == 0 {
		helper.HandleNotFound(w, errors.New("No Data"))
		return
	}

	helper.HandleSuccess(w, menus)
}

func (m *menuController) Add(w http.ResponseWriter, r *http.Request) {
	var request params.MenuCreate

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helper.HandleBadRequest(w, err)
		return
	}

	_, err = services.NewMenuService(m.db).CreateNewMenu(&request)

	if err != nil {
		helper.HandleInternalServerError(w, errors.New("Internal Server Error"))
		return
	} else {
		helper.HandleCreateSuccess(w, "Create new menu success!")
		return
	}
}
