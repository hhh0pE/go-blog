package models

func UserAuthorize(username string, password string) (*User, bool) {
	u := User{}
	query := Connection.Where("login = ? AND password = ?", username, password).First(&u)
	if query.RowsAffected > 0 {

		return &u, true
	}
	return nil, false
}

func GetUserByID(id int) (*User, bool) {
	u := User{}
	query := Connection.Where("id = ?", id).First(&u)
	if query.RowsAffected > 0 {
		return &u, true
	}
	return nil, false
}