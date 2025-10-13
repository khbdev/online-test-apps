package client

type User struct {
	Username string
	Password string
}


var Users = []User{
	{Username: "admin", Password: "123456"},
	{Username: "aziz", Password: "password"},
}


func FindUser(username, password string) bool {
	for _, u := range Users {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}
