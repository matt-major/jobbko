package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/matt-major/jobbko/src/awsc"
	"github.com/matt-major/jobbko/src/context"
)

func NewRouter(context *context.ApplicationContext) http.Handler {
	handlers := Handlers{
		context: context,
	}

	r := mux.NewRouter()
	r.HandleFunc("/status", handlers.StatusHandler).Methods("GET")
	r.HandleFunc("/schedule", handlers.ScheduleEventHandler).Methods("POST")

	return r
}

type Handlers struct {
	context *context.ApplicationContext
}

func (h *Handlers) StatusHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "UP"})
}

func (h *Handlers) ScheduleEventHandler(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	eventType := params["type"][0]
	destination := params["destination"][0]
	scheduleAt, _ := strconv.Atoi(params["scheduleAt"][0])
	groupId := strconv.Itoa(getGroupId())

	reqBody, _ := ioutil.ReadAll(req.Body)

	newEvent := awsc.ScheduledEventItem{
		Id:      getEventId(scheduleAt),
		GroupId: groupId,
		State:   "PENDING",
		Data: ScheduledEventData{
			Type:        eventType,
			Destination: destination,
			CreatedAt:   time.Now().Unix(),
			ScheduledAt: scheduleAt,
			Payload:     reqBody,
		},
	}

	h.context.AwsClient.InsertEvent(newEvent)

	w.WriteHeader(http.StatusCreated)
}

func getEventId(scheduledAt int) string {
	return fmt.Sprintf("%s#%s", strconv.Itoa(scheduledAt), uuid.New().String())
}

func getGroupId() int {
	min := 0
	max := 10

	return rand.Intn(max-min) + min
}
