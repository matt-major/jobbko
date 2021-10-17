package main

import (
	"fmt"
)

type Processor struct {
	id                int
	allocatedGroupIds []int
}

func NewProcessor(processorId int, allocatedGroupIds []int) *Processor {
	p := &Processor{
		id:                processorId,
		allocatedGroupIds: allocatedGroupIds,
	}

	return p
}

func (p *Processor) Start() {
	fmt.Println("Starting Processor", p.id, ", allocatedGroupIds:", p.allocatedGroupIds)
	p.process()
}

func (p *Processor) process() {
	// TODO Scan DynamoDB for group, returning events to process
	// TODO Loop to create goroutines for each event, create waitgroup for these
	// TODO -> In each goroutine, try lock the event, stop if not successful
	// TODO -> If successful, try to dispatch the event to SQS
	// TODO -> When done, delete item from DynamoDB
	// TODO Wait for all goroutines to complete, restart processor when done
}
