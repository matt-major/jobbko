package main

import (
	"fmt"
	"sync"
	"time"
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
	// TODO Scan DynamoDB for group, returning events to process

	var wg sync.WaitGroup                            // WaitGroup to track event processing state
	limiter := make(chan struct{}, p.maxConcurrency) // Limiter for number of goroutines

	// TODO Loop to create goroutines for each event, create waitgroup for these
	for i := 0; i < 200; i++ {
		wg.Add(1)
		limiter <- struct{}{}
		go func(n int, groupId int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println("Processor", p.id, "Group", groupId, "Loop", n)
			time.Sleep(1 * time.Second)
			<-limiter
		}(i, groupId, &wg)
	}

	// TODO -> In each goroutine, try lock the event, stop if not successful
	// TODO -> If successful, try to dispatch the event to SQS
	// TODO -> When done, delete item from DynamoDB

	wg.Wait() // Wait for all events to be processed
	fmt.Println("Done")

	// p.scanGroup(groupId) // Trigger next "tick"
}
