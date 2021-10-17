package main

type ScheduledEvent struct {
	Id      string             `json:"Id"`
	GroupId string             `json:"groupId"`
	State   string             `json:"state"`
	Data    ScheduledEventData `json:"event"`
}

type ScheduledEventData struct {
	Type        string `json:"type"`
	Destination string `json:"destination"`
	CreatedAt   int64  `json:"createdAt"`
	ScheduledAt int    `json:"scheduledAt"`
	Payload     []byte `json:"payload"`
}
