package routing

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

type Router struct {
	roomId          string
	userConnections map[string]*websocket.Conn
	inputChannel    chan string
}

func NewRouter(roomId string) *Router {
	fmt.Printf("[Router] Initializing router for room: %s\n", roomId)

	r := new(Router)
	r.userConnections = map[string]*websocket.Conn{}
	r.inputChannel = make(chan string)

	go func() {
		for {
			msg := <-r.inputChannel
			fmt.Printf("[Router] Received message from user [%s]\n", msg)
			go r.Broadcast(msg)
		}
	}()
	return r
}

func (r *Router) Broadcast(msg string) {
	fmt.Printf("[Router] Broadcasting message from user %s\n", msg)
	for _, conn := range r.userConnections {
		println("----SENDING")
		websocket.Message.Send(conn, msg)
		println("----SENT")
	}
}

func (r *Router) AddUser(username string, c *gin.Context) {
	fmt.Printf("[Router] Adding user %s to router\n", username)

	handler := websocket.Handler(func(conn *websocket.Conn) {
		// Add user connection to pool of connections
		r.userConnections[username] = conn

		//Start listening for message from user and push the to router input channel
		for {
			var msg string
			err := websocket.Message.Receive(conn, &msg)

			if err != nil {
				fmt.Printf("Error while receiving message from user %s\n", username)
			}
			r.inputChannel <- msg
		}
	})
	handler.ServeHTTP(c.Writer, c.Request)
}
