package main

type ClientProfile struct {
	Email string
	Id    string
	Name  string
	Token string
}

var database = map[string]ClientProfile{
	"user1": {
		Email: "user1@mail.com",
		Id:    "user1",
		Name:  "User One",
		Token: "123",
	},
	"user2": {
		Email: "user2@mail.com",
		Id:    "user2",
		Name:  "User Two",
		Token: "456",
	},
}
