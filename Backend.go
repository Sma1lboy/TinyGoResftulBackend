package main

type User struct {
	ID   int
	Name string
	Male bool
}

var users = []User{
	{ID: 1, Name: "Blue Train", Male: true},
	{ID: 2, Name: "Jeru", Male: false},
	{ID: 3, Name: "Sarah Vaughan and Clifford Brown", Male: false},
}
var id = 4

func getAllUsers() []User {
	return users
}
func addUser(user User) {
	users = append(users, user)
}
func removeUser(id int) User {
	var user User
	for i, e := range users {
		if e.ID == id {
			user = e
			users = append(users[:i], users[i+1:]...)
		}

	}
	return user
}
