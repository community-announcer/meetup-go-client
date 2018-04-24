package models

type Scroll string

const (
	RecentPast   Scroll = "recent_past"
	NextUpcoming Scroll = "next_upcoming"
	FutureOrPast Scroll = "future_or_past"
)

type EventStatus string

const (
	Cancelled EventStatus = "cancelled"
	Draft     EventStatus = "draft"
	Past      EventStatus = "past"
	Proposed  EventStatus = "proposed"
	Suggested EventStatus = "suggested"
	Upcoming  EventStatus = "upcoming"
)
