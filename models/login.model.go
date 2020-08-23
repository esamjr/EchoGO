package models

import (
	"database/sql"
	"echo/db"
	"echo/helpers"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.ID, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("username Not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query Error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash And Password Doesn't match")
		return false, err
	}

	return true, nil

}
