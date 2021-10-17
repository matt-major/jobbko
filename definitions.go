package main

type ScheduledEventData struct {
	Type        string `json:"type"`
	Destination string `json:"destination"`
	CreatedAt   int64  `json:"createdAt"`
	ScheduledAt int    `json:"scheduledAt"`
	Payload     []byte `json:"payload"`
}
