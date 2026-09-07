package main

import "time"

type EventSource interface {
	nextEvent() Event
}

type SyntheticEventSource struct {
	idleSecs int
}

func NewSyntheticEventSource() *SyntheticEventSource {
	return &SyntheticEventSource{
		idleSecs: 0,
	}
}

func (n *SyntheticEventSource) nextEvent() Event {
	// for testing, let's just return idle 5, 10, 15 and 20 events
	// in fractions of seconds
	// also emit the userinput event at specific points during the 20 mins
	// time window
	time.Sleep(10 * time.Millisecond)
	n.idleSecs += 1
	switch n.idleSecs {
	case 20 * 60: // 20 mins
		return EventIdle20m
	case 15 * 60: // 15 mins
		return EventIdle15m
	case 10 * 60: // 10 mins
		return EventIdle10m
	case 5 * 60: // 5 mins
		return EventIdle5m
	}
	return EventNone
}
