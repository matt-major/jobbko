package processor

import "github.com/matt-major/jobbko/pkg/context"

type ProcessorOrchestrator struct {
	NumProcessors  int
	NumGroups      int
	MaxConcurrency int
}

func (po *ProcessorOrchestrator) StartProcessors(context *context.ApplicationContext) {
	context.Logger.Info("Starting event processors...")

	processors := make([]int, po.NumProcessors)
	for i := range processors {
		go func(id int) {
			groupIds := po.getGroupIdsForProcessor(id)
			p := NewProcessor(id, groupIds, po.MaxConcurrency, context)
			context.Logger.Info("Created event processor ", id)
			p.Start()
		}(i)
	}
}

func (po *ProcessorOrchestrator) getGroupIdsForProcessor(processorId int) []int {
	bucket := po.NumGroups / po.NumProcessors

	startGroup := processorId * bucket

	var endGroup int
	if processorId == po.NumProcessors-1 {
		endGroup = po.NumGroups - 1
	} else {
		endGroup = ((processorId + 1) * bucket) - 1
	}

	groups := make([]int, endGroup-startGroup+1)
	for i := range groups {
		groups[i] = startGroup + i
	}

	return groups
}
