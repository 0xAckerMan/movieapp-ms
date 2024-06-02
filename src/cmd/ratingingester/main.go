package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	model "github.com/0xAckerMan/movieapp-ms/rating/pkg"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	fmt.Println("Creating a Kafka producer")

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

    const fileName = "ratingsdata.json"

    fmt.Println("reading rating data from: ", fileName)

    ratingEvents, err := ReadRatingEvents(fileName)
    if err != nil{
        panic(err)
    }

    var topic = "ratings"
    if err := ProduceRatingEvents(topic,producer,ratingEvents); err != nil{
        panic(err)
    }

    var timeout = 10 * time.Second
    fmt.Println("waiting " + timeout.String() + " until all events get produced")

    producer.Flush(int(timeout.Milliseconds()))
}

func ReadRatingEvents(filename string) ([]model.RatingEvent, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var ratings []model.RatingEvent

	if err := json.NewDecoder(f).Decode(&ratings); err != nil {
		return nil, err
	}

	return ratings, nil
}

func ProduceRatingEvents(topic string, producer *kafka.Producer, events []model.RatingEvent) error {
	for _, event := range events {
		encodedEvent, err := json.Marshal(event)
		if err != nil {
			return err
		}
         if err := producer.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Value: []byte(encodedEvent),
        }, nil); err != nil{
            return err
        }
	}
    return nil
}
