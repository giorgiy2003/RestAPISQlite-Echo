package logic

import (
	model "myapp/internal/model"
	repository "myapp/internal/repository"
	"strconv"
)

func GetAll() ([]model.Person, error) {
	rows, err := repository.Connection.Query(`SELECT * FROM "person" ORDER BY "person_id"`)
	if err != nil{
	   return nil, err
   }
   defer rows.Close()
   persons := []model.Person{}
   for rows.Next(){
   var p model.Person
	   err := rows.Scan(&p.Id,&p.Email,&p.Phone,&p.FirstName, &p.LastName)
	   if err != nil{
		   return nil,err
	   }
	   persons = append(persons, p)
   }
   return persons, nil
}

func GetById(id string) ([]model.Person, error) {
	person_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	rows, err := repository.Connection.Query(`SELECT * FROM "person" WHERE "person_id" = $1`, person_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	persons := []model.Person{}
	for rows.Next(){
	var p model.Person
		err := rows.Scan(&p.Id,&p.Email,&p.Phone,&p.FirstName, &p.LastName)
		if err != nil{
			return nil,err
		}
		persons = append(persons, p)
	}
	return persons, nil
}

func Create(p model.Person) error {
	if _, err := repository.Connection.Exec(`INSERT INTO "person" ("person_email", "person_phone", "person_firstName", "person_lastName") VALUES ($1, $2,$3,$4)`, p.Email, p.Phone, p.FirstName, p.LastName); err != nil {
		return err
	}
	return nil
}

func Update(p model.Person, id string) error {
	person_id,err := strconv.Atoi(id)
	if err != nil{
		return err
	}
	if _, err := repository.Connection.Exec(`UPDATE "person" SET "person_email" = $1,"person_phone" = $2,"person_firstName" = $3,"person_lastName" = $4  WHERE "person_id" = $5`, p.Email, p.Phone, p.FirstName, p.LastName, person_id); err != nil {
		return err
	}
	return nil
}

func Delete(id string) error {
	person_id,err := strconv.Atoi(id)
	if err != nil{
		return err
	}
	if _, err := repository.Connection.Exec(`DELETE FROM "person" WHERE "person_id" = $1`, person_id); err != nil {
		return err
	}
	return nil
}


