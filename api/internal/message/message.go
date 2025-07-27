package message

import "time"

type Message struct {
	Id        int
	Text      string
	CreatedAt time.Time
}
