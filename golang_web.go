package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoohawa/golang_web/models"
	"golang.org/x/net/websocket"
	"net/http"
)

var userRepo models.UserRepository
var roomRepo models.RoomRepository

func main() {
	// init repositories
	userRepo = SetupUserRepository()
	roomRepo = SetupRoomRepository()
	//

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "./css")

	r.GET("/", func(c *gin.Context) {
		userRepo.Print()

		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Index"})
	})

	r.POST("/login", func(c *gin.Context) {

		username := c.PostForm("username")

		fmt.Printf("Logged in user: %s\n", username)
		// TODO: Check if username is taken

		// Add auth cookie and redirect to lobby
		userRepo.AddUser(username)
		c.SetCookie("username", username, 0, "/", "localhost", false, false)
		c.Redirect(http.StatusFound, "/lobby")
	})

	r.GET("/lobby", func(c *gin.Context) {
		userRepo.Print()

		username, _ := c.GetCookie("username")

		c.HTML(http.StatusOK, "lobby.tmpl", gin.H{
			"user":      username,
			"roomNames": roomRepo.Rooms,
		})
	})

	r.GET("/room/:roomId", func(c *gin.Context) {
		userRepo.Print()

		//Get Room
		roomId := c.Param("roomId")
		room, _ := roomRepo.GetRoom(roomId)
		room2, _ := roomRepo.GetRoom(roomId)
		println("RoomId", roomId, "RoomName", room.Name, "UsrCnt", len(room.Users))

		// Get user
		username, _ := c.GetCookie("username")
		user, _ := userRepo.GetUser(username)
		println("Username", username, "User", user.Username)

		//Add user to room
		println("Room", &room, "Room2", &room2)
		room.AddUser(user)

		c.HTML(http.StatusOK, "room.tmpl", gin.H{
			"Room": room,
		})
	})

	r.GET("/room-ws/:roomId/", func(c *gin.Context) {
		//Get Room
		roomId := c.Param("roomId")
		room, _ := roomRepo.GetRoom(roomId)
		println("RoomId", roomId, "RoomName", room.Name, "UsrCnt", len(room.Users))

		// Get user
		username, _ := c.GetCookie("username")
		user, _ := userRepo.GetUser(username)
		println("Username", username, "User", user.Username)

		//Setup router
		room.Router.AddUser(username, c)

	})

	r.GET("/conversation", func(c *gin.Context) {
		c.HTML(http.StatusOK, "conversation.tmpl", gin.H{"title": "Conversation"})
	})

	r.GET("/ws", func(c *gin.Context) {
		for k, v := range c.Request.Header {
			fmt.Printf("%s: %s\n", k, v)
		}

		handler := websocket.Handler(func(conn *websocket.Conn) {
			println("Handling connection")

			for {
				println("Listening for a message")
				var message string
				err := websocket.Message.Receive(conn, &message)
				if err != nil {
					print("There was an error")
					break
				}

				print("Echoing back the message: " + message)
				websocket.Message.Send(conn, message)
			}
		})

		handler.ServeHTTP(c.Writer, c.Request)
	})

	r.Run()
}

func SetupRoomRepository() models.RoomRepository {
	roomRepo := models.RoomRepository{}
	roomRepo.Init()
	return roomRepo
}

func SetupUserRepository() models.UserRepository {
	userRepo := models.UserRepository{}
	userRepo.Init()
	return userRepo
}
