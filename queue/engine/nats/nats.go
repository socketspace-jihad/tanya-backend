package nats

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/socketspace-jihad/tanya-backend/queue/engine"
)

type NATS struct {
	Host     string
	Username string
	Password string
	*nats.Conn
}

func NewNATS() engine.QueueEngine {
	return &NATS{}
}

func (n *NATS) Publish(data any, topic string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return n.Conn.Publish(topic, jsonData)
}

func (n *NATS) Consume(topic string) (chan any, chan error) {
	c := make(chan any)
	errorChan := make(chan error)
	go func(c chan any) {
		_, err := n.Conn.Subscribe(topic, func(msg *nats.Msg) {
			c <- msg.Data
		})
		if err != nil {
			errorChan <- err
			close(c)
			close(errorChan)
			return
		}
	}(c)
	return c, errorChan
}

func (n *NATS) Connect(creds *engine.EngineAuthData) (err error) {
	n.Conn, err = nats.Connect(creds.Host)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	engine.RegisterQueueEngine("nats", NewNATS)
}
