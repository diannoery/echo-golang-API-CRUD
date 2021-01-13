package models

import (
	"database/sql"
	"echo/databases"
	"echo/helpers"
	"fmt"
)

type UserLogin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj UserLogin
	var pwd string

	con := databases.CreateCon()
	sqlDb := "select * from user where username = ? ;"

	err := con.QueryRow(sqlDb, username).Scan(&obj.ID, &obj.Username, &pwd)

	if err == sql.ErrNoRows {
		fmt.Println("username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helpers.CheckPassword(password, pwd)
	if !match {
		fmt.Println("password tidak saama")
		return false, err
	}

	return true, err
}
