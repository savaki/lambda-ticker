package internal

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestUnmarshalEvent(t *testing.T) {
	text := `{"version":"0","id":"bc4fcaac-7cb8-4da9-9abd-3c5b1aaa7da0","detail-type":"Scheduled Event","source":"aws.events","account":"554068800329","time":"2016-04-04T04:03:44Z","region":"us-east-1","resources":["arn:aws:events:us-east-1:554068800329:rule/hourly"],"detail":{}}`

	event := Event{}
	err := json.Unmarshal([]byte(text), &event)
	if err != nil {
		t.Errorf("expected nil err; got %v", err)
		return
	}

	triggered, err := event.TriggeredAt()
	if err != nil {
		t.Errorf("expected nil err; got %v", err)
		return
	}

	if v := triggered.Format(time.UnixDate); v != "Mon Apr  4 04:03:44 UTC 2016" {
		t.Errorf("expected Mon Apr  4 04:03:44 UTC 2016; got %v", v)
	}
}

func TestResourcesContain(t *testing.T) {
	text := `{"version":"0","id":"bc4fcaac-7cb8-4da9-9abd-3c5b1aaa7da0","detail-type":"Scheduled Event","source":"aws.events","account":"554068800329","time":"2016-04-04T04:03:44Z","region":"us-east-1","resources":["arn:aws:events:us-east-1:554068800329:rule/hourly"],"detail":{}}`

	event := Event{}
	json.NewDecoder(strings.NewReader(text)).Decode(&event)

	if v := event.ResourcesContain("hourly"); v != true {
		t.Errorf("expected true; got %v", v)
		return
	}
	if v := event.ResourcesContain("blah"); v != false {
		t.Errorf("expected false; got %v", v)
		return
	}
}
