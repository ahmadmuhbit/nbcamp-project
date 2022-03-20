package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"nbcamp_project/server/helper"
	params "nbcamp_project/server/params/menu"
	"nbcamp_project/server/services"
	"net/http"
)

type MenuController interface {
	Add(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
}

type menuController struct {
	db *sql.DB
}

func NewMenuController(db *sql.DB) MenuController {
	return &menuController{db}
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
		helper.HandleInternalServerError(w, errors.New("INTERNAL SERVER ERROR"))
		return
	} else {
		helper.HandleCreateSuccess(w, "Create new menu success!")
		return
	}
}

func (m *menuController) FindAll(w http.ResponseWriter, r *http.Request) {
	menus, err := services.NewMenuService(m.db).GetAllMenu()

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			helper.HandleNotFound(w, errors.New("NO DATA"))
		} else {
			helper.HandleInternalServerError(w, err)
		}
		return
	}

	if len(*menus) == 0 {
		helper.HandleNotFound(w, errors.New("NO DATA"))
		return
	}

	helper.HandleSuccess(w, menus)
}

func (m *menuController) FindByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	_, err := uuid.Parse(id)
	if err != nil {
		helper.HandleBadRequest(w, err)
		return
	}

	menu, err := services.NewMenuService(m.db).GetMenuByID(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			helper.HandleNotFound(w, errors.New("NO DATA"))
		} else {
			helper.HandleInternalServerError(w, err)
		}
		return
	}

	helper.HandleSuccess(w, menu)
}
