package engine

import (
	"fmt"
)

type QueueEngine interface {
	Connect(*EngineAuthData) error
	Publish(any, string) error
	Consume(string) (chan any, chan error)
}

type QueueEngineFactory func() QueueEngine

var queueEngineMap map[string]QueueEngineFactory = make(map[string]QueueEngineFactory)

func RegisterQueueEngine(name string, engine QueueEngineFactory) {
	queueEngineMap[name] = engine
}

func GetQueueEngine(name string) (QueueEngineFactory, error) {
	if _, ok := queueEngineMap[name]; !ok {
		return nil, fmt.Errorf("queue engine named: %v not found", name)
	}
	return queueEngineMap[name], nil
}
