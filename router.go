package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/matt-major/jobbko/awsc"
)

func NewRouter() http.Handler {
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
	groupId := strconv.Itoa(getGroupId())

	reqBody, _ := ioutil.ReadAll(req.Body)

	newEvent := awsc.ScheduledEventItem{
		Id:      uuid.New().String(),
		GroupId: groupId,
		State:   "SCHEDULED",
		Data: ScheduledEventData{
			Type:        eventType,
			Destination: destination,
			CreatedAt:   time.Now().Unix(),
			ScheduledAt: scheduleAt,
			Payload:     reqBody,
		},
	}

	awsc.InsertEvent(newEvent)

	w.WriteHeader(http.StatusCreated)
}

func getGroupId() int {
	min := 0
	max := 5

	return rand.Intn(max-min) + min
}
