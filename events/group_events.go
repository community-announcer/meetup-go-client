package events

import (
	"net/url"

	"github.com/community-announcer/meetup-go-client/models"
)

type Attendance struct {
}

type EventHost struct {
}

type Events struct {
	attendance_count   int
	attendance_sample  []Attendance
	comment_count      int
	created            int
	description        string
	description_images []url.URL
	duration           int
	event_hosts        []EventHost
}

type EventClient struct {
	*models.Client
}

func (c *EventClient) GetGroupEvents(desc bool, fields []string, page int, scroll models.Scroll, status []models.EventStatus) (events *Events, e error) {

	return nil, nil
}
