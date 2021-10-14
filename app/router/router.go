package router

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	aws "github.com/matt-major/jobbko/app/aws"
	"github.com/matt-major/jobbko/app/scheduler"
)

func New() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/status", statusHandler).Methods("GET")
	r.HandleFunc("/schedule", scheduleEventHandler).Methods("POST")

	return r
}

func statusHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "UP"})
}

func scheduleEventHandler(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	eventType := params["type"][0]
	destination := params["destination"][0]
	scheduleAt, _ := strconv.Atoi(params["scheduleAt"][0])
	shardId := strconv.Itoa(getShardId())

	reqBody, _ := ioutil.ReadAll(req.Body)

	newEvent := scheduler.ScheduledEvent{
		ScheduleId: uuid.New().String(),
		ShardId:    shardId,
		State:      "SCHEDULED",
		Event: scheduler.ScheduledEventData{
			Type:        eventType,
			Destination: destination,
			CreatedAt:   time.Now().Unix(),
			ScheduledAt: scheduleAt,
			Payload:     reqBody,
		},
	}

	aws.InsertEvent(newEvent)

	w.WriteHeader(http.StatusCreated)
}

func getShardId() int {
	min := 0
	max := 5

	return rand.Intn(max-min) + min
}
