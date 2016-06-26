package models

import (
	"errors"
	"fmt"
)

type User struct {
	Username string
}

type UserRepository struct {
	users map[string]string
}

func (r *UserRepository) Init() {
	r.users = map[string]string{}
}

func (r *UserRepository) GetUser(username string) (User, error) {
	_, ok := r.users[username]
	if ok {
		println("[UserRepository][GetUser] Found user ", username)
		return User{username}, nil
	}
	return User{}, errors.New(fmt.Sprintf("User %s doesn't exist", username))
}

func (r *UserRepository) AddUser(username string) error {
	_, ok := r.users[username]
	if !ok {
		println("[UserRepository][AddUser] Adding user ", username)
		r.users[username] = username
		return nil
	} else {
		return errors.New(fmt.Sprintf("User %s already exists", username))
	}
}

func (r *UserRepository) Print() {
	print("[UserRep]: ")
	for k, v := range r.users {
		print("[", k, ",", v, "],")
	}
	println()
}
