package main

type User struct {
	ID   int
	Name string
	Male bool
}

var users = []User{
	{ID: 1, Name: "John Doe", Male: true},
	{ID: 2, Name: "Jane Smith", Male: false},
	{ID: 3, Name: "Bob Johnson", Male: true},
	{ID: 4, Name: "Emily Brown", Male: false},
	{ID: 5, Name: "Michael Wilson", Male: true},
	{ID: 6, Name: "Samantha Davis", Male: false},
	{ID: 7, Name: "William Lee", Male: true},
	{ID: 8, Name: "Olivia Jones", Male: false},
	{ID: 9, Name: "Daniel Garcia", Male: true},
	{ID: 10, Name: "Sophia Martinez", Male: false},
}

var id = users[len(users)-1].ID + 1

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
