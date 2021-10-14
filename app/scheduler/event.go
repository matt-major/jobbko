package scheduler

type ScheduledEvent struct {
	ScheduleId string             `json:"scheduleId"`
	ShardId    string             `json:"shardId"`
	State      string             `json:"state"`
	Event      ScheduledEventData `json:"event"`
}

type ScheduledEventData struct {
	Type        string `json:"type"`
	Destination string `json:"destination"`
	CreatedAt   int64  `json:"createdAt"`
	ScheduledAt int    `json:"scheduledAt"`
	Payload     []byte `json:"payload"`
}
