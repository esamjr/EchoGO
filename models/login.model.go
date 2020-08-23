package models

import (
	"database/sql"
	"echo/db"
	"echo/helpers"
	"fmt"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func CreateUser(username, password string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sql := "INSERT INTO users(username, password) VALUES(?, ?)"

	stmt, err := con.Prepare(sql)

	if err != nil {
		return res, err
	}

	hash, _ := helpers.HashPassword(password)
	result, err := stmt.Exec(username, hash)
	if err != nil {
		return res, err
	}

	resp, err := result.LastInsertId()
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"response": resp,
	}
	res.Data = map[string]string{
		"username": username,
		"password": hash,
	}
	return res, nil
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
