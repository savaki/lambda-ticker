package internal

import (
	"strings"
	"time"
)

type Event struct {
	Id        string   `json:"id"`
	Time      string   `json:"time"`
	Resources []string `json:"resources"`
}

func (e Event) TriggeredAt() (time.Time, error) {
	return time.Parse(time.RFC3339, e.Time)
}

func (e Event) ResourcesContain(s string) bool {
	if e.Resources == nil {
		return false
	}

	for _, r := range e.Resources {
		if strings.Contains(r, s) {
			return true
		}
	}

	return false
}
