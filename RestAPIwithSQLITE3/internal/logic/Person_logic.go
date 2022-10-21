package logic

import (
	"errors"
	"fmt"
	model "myapp/internal/model"
	repository "myapp/internal/repository"
	"strconv"
)

func Create(p model.Person) error {
	if _, err := repository.Connection.Exec(`INSERT INTO "person" ("person_email", "person_phone", "person_firstName", "person_lastName") VALUES ($1, $2,$3,$4)`, p.Email, p.Phone, p.FirstName, p.LastName); err != nil {
		return err
	}
	return nil
}

func ReadOne(id string) ([]model.Person, error) {
	person_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("Error: неверно введён параметр id")
	}
	row, err := repository.Connection.Query(`SELECT * FROM "person" WHERE "person_id" = $1`, person_id)
	if err != nil {
		return nil, err
	}
	var personInfo = []model.Person{}
	for row.Next() {
		var p model.Person
		err := row.Scan(&p.Id, &p.Email, &p.Phone, &p.FirstName, &p.LastName)
		if err != nil {
			return nil, err
		}
		personInfo = append(personInfo, p)
	}
	return personInfo, nil
}

func Read() ([]model.Person, error) {
	row, err := repository.Connection.Query(`SELECT * FROM "person" ORDER BY "person_id"`)
	if err != nil {
		return nil, err
	}
	var personInfo = []model.Person{}
	for row.Next() {
		var p model.Person
		err := row.Scan(&p.Id, &p.Email, &p.Phone, &p.FirstName, &p.LastName)
		if err != nil {
			return nil, err
		}
		personInfo = append(personInfo, p)
	}
	return personInfo, nil
}

func Update(p model.Person, id string) error {
	if err := dataExist(id); err != nil {
		return err
	}
	if _, err := repository.Connection.Exec(`UPDATE "person" SET "person_email" = $1,"person_phone" = $2,"person_firstName" = $3,"person_lastName" = $4  WHERE "person_id" = $5`, p.Email, p.Phone, p.FirstName, p.LastName, id); err != nil {
		return err
	}
	return nil
}

func Delete(id string) error {
	if err := dataExist(id); err != nil {
		return err
	}
	if _, err := repository.Connection.Exec(`DELETE FROM "person" WHERE "person_id" = $1`, id); err != nil {
		return err
	}
	return nil
}

func dataExist(id string) error {
	persons, err := ReadOne(id)
	if err != nil {
		return err
	}
	if len(persons) == 0 {
		return fmt.Errorf("Error: записи с id = %s не существует", id)
	}
	return nil
}