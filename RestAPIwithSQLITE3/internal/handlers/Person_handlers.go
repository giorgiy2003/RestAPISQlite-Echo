package handler

import (
	"fmt"
	"log"
	logic "myapp/internal/logic"
	model "myapp/internal/model"
	"net/http"

	"github.com/labstack/echo"
)

func Add(c echo.Context) error {
	var p model.Person
	p.Email = c.FormValue("email")
	p.Phone = c.FormValue("phone")
	p.FirstName = c.FormValue("firstName")
	p.LastName = c.FormValue("lastName")
	err := logic.Create(p)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Println("Создана запись:", p)
	return c.JSON(http.StatusCreated, p)
}

func GetAll(c echo.Context) error {
	persons, err := logic.GetAll()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, persons)
}

func GetById(c echo.Context) error {
	id := c.Param("id")
	persons, err := logic.GetById(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Println(persons)
	return c.JSON(http.StatusOK, persons)
}

func Delete(c echo.Context) error {
	id := c.Param("id")
	err := logic.Delete(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Printf("Запись с id = %s  успешно удалена", id)
	return c.JSON(http.StatusOK, fmt.Sprintf("Запись с id = %s  успешно удалена", id))
}

func Update(c echo.Context) error {
	var newPerson model.Person
	id := c.Param("id")
	newPerson.Email = c.FormValue("email")
	newPerson.Phone = c.FormValue("phone")
	newPerson.FirstName = c.FormValue("firstName")
	newPerson.LastName = c.FormValue("lastName")
	err := logic.Update(newPerson, id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Printf("Запись с id = %s  успешно обновлена", id)
	return c.JSON(http.StatusOK, fmt.Sprintf("Запись с id = %s  успешно обновлена", id))
}
