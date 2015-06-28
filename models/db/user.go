package db

import "fmt"

type User struct {
	ID                 int
	Login, Email, Role string
}

func (u *User) Authorize(username string, password string) bool {
	query := Connection.Where("login = ? AND password = ?", username, password).First(u)
	if query.RowsAffected > 0 {
		fmt.Println(query)
		fmt.Println(u)
		return true
	}
	return false
}

func (u *User) GetByID(id int) bool {
	query := Connection.Where("id = ?", id).First(u)
	return query.RowsAffected > 0
}
