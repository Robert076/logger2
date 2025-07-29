package message

import "time"

type Message struct {
	Id        int
	Message   string
	CreatedAt time.Time
}
