package models

import (
	"errors"
	"fmt"
	"github.com/hoohawa/golang_web/routing"
	"math/rand"
	"time"
)

type Room struct {
	Id       string
	Name     string
	Users    []User
	Messages []Message
	Router   *routing.Router
}

func (r *Room) AddUser(user User) {
	found := false
	for _, v := range r.Users {
		if v.Username == user.Username {
			found = true
		}
	}

	if !found {
		println("Adding user", user.Username, "to room", r.Name)
		r.Users = append(r.Users, user)
	}
}

type RoomRepository struct {
	Rooms   []Room
	roomIds map[string]string
}

func (r *RoomRepository) GetRoom(id string) (*Room, error) {
	for idx, room := range r.Rooms {
		if room.Id == id {
			return &r.Rooms[idx], nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Could not find room with id: %s", id))
}

func (r *RoomRepository) Init() {
	println("[RoomRepository] Initializing")
	r.Rooms = []Room{}

	for i := 0; i < 10; i++ {
		r.CreateNewRoom(fmt.Sprintf("Room %d", i))
	}

	// Extra testing room
	r.Rooms = append(r.Rooms, Room{
		"test",
		"TestRoom",
		[]User{User{"SystemUser"}},
		[]Message{
			Message{
				"SystemUser",
				"Welcome to the room",
				time.Now(),
			},
		},
		routing.NewRouter("test"),
	})
}

func (r *RoomRepository) CreateNewRoom(name string) {
	roomId := r.genRoomsId()
	r.Rooms = append(r.Rooms, Room{
		roomId,
		name,
		[]User{},
		[]Message{},
		routing.NewRouter(roomId),
	})
	rand.Seed(time.Now().UnixNano())
}

func (r *RoomRepository) genRoomsId() string {
	randId := genRandomString(10)
	_, idPresent := r.roomIds[randId]
	for idPresent == true {
		randId = genRandomString(10)
	}
	return randId
}

func genRandomString(length int) string {
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456798")
	str := make([]rune, length)
	for i := range str {
		str[i] = runes[rand.Intn(length)]
	}
	return string(str)
}
