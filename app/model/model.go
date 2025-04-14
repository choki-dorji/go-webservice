package model

import (
	"fmt"
	"myapp/app/datastore/postgres"
)

type Student struct {
	StdId     int64  `json:"stdid"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
}

const queryInsertUser = "INSERT INTO student(stdid, firstname, lastname, email) VALUES($1, $2, $3, $4);"
const querygetUser = "SELECT stdid, firstname, lastname, email FROM student where stdid=$1"
const queryUpdateUser = "UPDATE student SET stdid=$1, firstname=$2,lastname=$3, email=$4 WHERE stdid=$5 RETURNING stdid;"

func (s *Student) Create() error {
	_, err := postgres.Db.Exec(queryInsertUser, s.StdId, s.FirstName, s.LastName, s.Email)
	fmt.Println("err", err)
	return err
}

func (s *Student) Read() error {
	return postgres.Db.QueryRow(querygetUser, s.StdId).Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
}
func (s *Student) Update(oldId int64) error {
	err := postgres.Db.QueryRow(queryUpdateUser, s.StdId, s.FirstName, s.LastName, s.Email, oldId).Scan(&s.StdId)
	return err
}
