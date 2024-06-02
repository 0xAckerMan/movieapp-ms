package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	model "github.com/0xAckerMan/movieapp-ms/rating/pkg"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Ingester struct {
	consumer *kafka.Consumer
	topic    string
}

func NewIngester(addr string, groupID string, topic string) (*Ingester, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"boostrap.servers":  addr,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	return &Ingester{consumer, topic}, nil
}

// Ingest starts ingestion from Kafka and returns a channel containing rating events
// representing the data consumed from the topic.

func (i *Ingester) Ingest(ctx context.Context) (chan model.RatingEvent, error) {
	if err := i.consumer.SubscribeTopics([]string{i.topic}, nil); err != nil {
		return nil, err
	}

	ch := make(chan model.RatingEvent, 1)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				i.consumer.Close()
				return
			default:
			}

            msg, err := i.consumer.ReadMessage(-1)
            if err != nil{
                fmt.Println("consumer error: " + err.Error())
                continue
            }

            var event model.RatingEvent
            if err := json.Unmarshal(msg.Value, &event); err != nil{
                fmt.Println("Unmarshal error: " + err.Error())
                continue
            }
             ch <- event
		}
	}()
    return ch, nil
}
