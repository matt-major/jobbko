package main

import (
	"fmt"
	"sync"

	"github.com/matt-major/jobbko/awsc"
)

type Processor struct {
	id                int
	allocatedGroupIds []int
	maxConcurrency    int
}

func NewProcessor(processorId int, allocatedGroupIds []int, maxConcurrency int) *Processor {
	p := &Processor{
		id:                processorId,
		allocatedGroupIds: allocatedGroupIds,
		maxConcurrency:    maxConcurrency,
	}

	return p
}

func (p *Processor) Start() {
	fmt.Println("Starting Processor", p.id, ", allocatedGroupIds:", p.allocatedGroupIds)

	for i := range p.allocatedGroupIds {
		go p.scanGroup(p.allocatedGroupIds[i])
	}
}

func (p *Processor) scanGroup(groupId int) {
	items := awsc.GetProcessableEvents(groupId, 250)

	var wg sync.WaitGroup                            // WaitGroup to track event processing state
	limiter := make(chan struct{}, p.maxConcurrency) // Limiter for number of goroutines

	for i := range items {
		wg.Add(1)
		limiter <- struct{}{} // Add to limiter, blocks if too many concurrent goroutines
		go p.processEvent(items[i], &wg, limiter)
	}

	wg.Wait() // Wait for all events to be processed

	p.scanGroup(groupId) // Trigger next "tick"
}

func (p *Processor) processEvent(event awsc.ScheduledEventItem, wg *sync.WaitGroup, limiter chan struct{}) {
	defer wg.Done()

	lock := p.lockEvent(event)
	if !lock {
		<-limiter
		return
	}

	// TODO -> If successful, try to dispatch the event to SQS

	p.deleteEvent(event.Id) // If dispatched, delete from DynamoDB

	<-limiter
}

func (p *Processor) lockEvent(event awsc.ScheduledEventItem) bool {
	return awsc.LockEvent(event)
}

func (p *Processor) deleteEvent(id string) {
	// TODO Delete event in DynamoDB
	// TODO Log error regarding deletion
}
