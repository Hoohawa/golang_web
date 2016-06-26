package models

import "time"

type Message struct {
	Username string
	Text     string
	Time     time.Time
}
