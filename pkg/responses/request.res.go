package responses

import "time"

type RequestRes struct {
	Id         int
	SenderID   int
	ReceiverID int
	Status     int
	CreatedAt  time.Time
}