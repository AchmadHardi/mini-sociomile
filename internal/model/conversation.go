package model

import "time"

type Conversation struct {
	ID        int64
	TicketID  int64
	SenderID  int64
	Message   string
	CreatedAt time.Time
}
