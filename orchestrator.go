package main

type ProcessorOrchestrator struct {
	numProcessors int
	numGroups     int
}

func (po *ProcessorOrchestrator) StartProcessors() {
	processors := make([]int, po.numProcessors)
	for i := range processors {
		go func(id int) {
			groupIds := po.getGroupIdsForProcessor(id)
			p := NewProcessor(id, groupIds)
			p.Start()
		}(i)
	}
}

func (po *ProcessorOrchestrator) getGroupIdsForProcessor(processorId int) []int {
	bucket := po.numGroups / po.numProcessors

	startGroup := processorId * bucket

	var endGroup int
	if processorId == po.numProcessors-1 {
		endGroup = po.numGroups - 1
	} else {
		endGroup = ((processorId + 1) * bucket) - 1
	}

	groups := make([]int, endGroup-startGroup+1)
	for i := range groups {
		groups[i] = startGroup + i
	}

	return groups
}
