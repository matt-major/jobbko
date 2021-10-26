package processor

import (
	"encoding/json"
	"sync"

	"github.com/matt-major/jobbko/pkg/awsc"
	"github.com/matt-major/jobbko/pkg/context"
	"github.com/matt-major/jobbko/pkg/definitions"
	"github.com/sirupsen/logrus"
)

type Processor struct {
	id                int
	allocatedGroupIds []int
	maxConcurrency    int
	logger            *logrus.Logger
	context           *context.ApplicationContext
}

func NewProcessor(processorId int, allocatedGroupIds []int, maxConcurrency int, context *context.ApplicationContext) *Processor {
	p := &Processor{
		id:                processorId,
		allocatedGroupIds: allocatedGroupIds,
		maxConcurrency:    maxConcurrency,
		logger:            context.Logger,
		context:           context,
	}

	return p
}

func (p *Processor) Start() {
	p.context.Logger.Info("Starting Processor ", p.id, ", allocatedGroupIds: ", p.allocatedGroupIds)

	for i := range p.allocatedGroupIds {
		go p.scanGroup(p.allocatedGroupIds[i])
	}
}

func (p *Processor) scanGroup(groupId int) {
	items := p.context.AwsClient.GetProcessableEvents(groupId, 250)

	var wg sync.WaitGroup                            // WaitGroup to track event processing state
	limiter := make(chan struct{}, p.maxConcurrency) // Limiter for number of goroutines

	for i := range items {
		wg.Add(1)
		limiter <- struct{}{} // Add to limiter, blocks if too many concurrent goroutines
		go p.processEvent(items[i], &wg, limiter)
	}

	wg.Wait() // Wait for all events to be processed

	//lint:ignore SA5007 because this is meant to be recursive
	p.scanGroup(groupId) // Trigger next "tick"
}

func (p *Processor) processEvent(event awsc.ScheduledEventItem, wg *sync.WaitGroup, limiter chan struct{}) {
	defer wg.Done()

	lock := p.context.AwsClient.LockEvent(event)
	if !lock {
		<-limiter
		return
	}

	jsonData, _ := json.Marshal(event.Data)
	var eventData definitions.ScheduledEventData
	json.Unmarshal(jsonData, &eventData)

	hasSentMessage := p.context.AwsClient.SendEventToQueue(eventData.Payload, eventData.Destination)
	if hasSentMessage {
		p.context.AwsClient.DeleteEvent(event) // If dispatched, delete from DynamoDB
	}

	p.logger.Infof("Processed Event %s", event.Id)

	<-limiter
}
