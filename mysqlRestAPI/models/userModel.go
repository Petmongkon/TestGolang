package models

import (
	"database/sql"
	_ "database/sql"
	"fmt"

	"github.com/myproject/mysqlRestAPI/db"
)

type User struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Users struct {
	Users []User `json:"user"`
}

var con *sql.DB

func getAllUsers() Users {
	con := db.CreateCon()
	sqlStatement := "SELECT * FROM users"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := Users{}

	for rows.Next() {
		user := User{}

		err2 := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Firstname, &user.Lastname)
		if err2 != nil {
			fmt.Print(err2)
		}
		result.Users = append(result.Users, user)
	}
	return result
}
