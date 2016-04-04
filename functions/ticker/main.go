package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/apex/go-apex"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/savaki/lambda-ticker/internal"
)

type Tick struct {
	Event string                 `json:"event"`
	Data  map[string]interface{} `json:"data,omitempty"`
}

func main() {
	cfg := &aws.Config{Region: aws.String("us-east-1")}
	client := sns.New(session.New(cfg))

	apex.HandleFunc(func(raw json.RawMessage, ctx *apex.Context) (interface{}, error) {
		// 1. Parse the inbound event instance
		//
		event := internal.Event{}
		err := json.Unmarshal(raw, &event)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshal raw message - %v", err)
		}

		// 2. Extract the time the event was triggered
		triggeredAt, err := event.TriggeredAt()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse triggered at - %v", err)
			return nil, nil
		}

		// 3. Create the tickAt
		minute := triggeredAt.Minute()
		minute = minute - (minute % 5)
		tickAt := time.Date(
			triggeredAt.Year(),
			triggeredAt.Month(),
			triggeredAt.Day(),
			triggeredAt.Hour(),
			minute,
			0,
			0,
			triggeredAt.Location(),
		)
		timestamp := tickAt.Format(time.RFC3339)
		data, err := json.Marshal(Tick{
			Event: "tick:5m",
			Data: map[string]interface{}{
				"time": timestamp,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("Unable to marshal tick - %v", err)
		}

		// 4. Find the topic to post to
		out, err := client.CreateTopic(&sns.CreateTopicInput{
			Name: aws.String("ticker-5m"),
		})
		if err != nil {
			return nil, fmt.Errorf("Unable to upsert topic  - %v", err)
		}

		// 5. Publish the event
		_, err = client.Publish(&sns.PublishInput{
			Message:  aws.String(string(data)),
			TopicArn: out.TopicArn,
		})

		fmt.Fprintf(os.Stderr, "Successfully sent event, %v, to arn, %v", timestamp, *out.TopicArn)

		return nil, err
	})
}
