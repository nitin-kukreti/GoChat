package domain

import "time"

type Message struct {
	ID         int
	GroupID    int
	FromUserID int
	Body       string
	CreatedAt  time.Time
}
