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
	var newPerson model.Person
	newPerson.Email = c.FormValue("email")
	newPerson.Phone = c.FormValue("phone")
	newPerson.FirstName = c.FormValue("firstName")
	newPerson.LastName = c.FormValue("lastName")
	err := logic.Create(newPerson)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprint(err))
	}
	log.Println("Добавлена запись", newPerson)
	return c.JSON(http.StatusCreated, newPerson)
}

func GetAll(c echo.Context) error {
	persons, err := logic.Read()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprint(err))
	}
	return c.JSON(http.StatusOK, persons)
}

func GetById(c echo.Context) error {
	id := c.Param("id")
	persons, err := logic.ReadOne(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprint(err))
	}
	log.Println(persons)
	return c.JSON(http.StatusOK, persons)
}

func Delete(c echo.Context) error {
	id := c.Param("id")
	err := logic.Delete(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprint(err))
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
		return c.JSON(http.StatusBadRequest, fmt.Sprint(err))
	}
	log.Printf("Запись с id = %s  успешно обновлена", id)
	return c.JSON(http.StatusOK, fmt.Sprintf("Запись с id = %s  успешно обновлена", id))
}