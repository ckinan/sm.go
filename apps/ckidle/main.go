package main

import "fmt"

type Event string
type Action string

const (
	EventIdle5m    Event = "idle5m"
	EventIdle10m   Event = "idle10m"
	EventIdle15m   Event = "idle15m"
	EventIdle20m   Event = "idle20m"
	EventUserInput Event = "userInput"
	EventNone      Event = "none"
)

const (
	ActionLock       Action = "lock"
	ActionDisplayOn  Action = "displayOn"
	ActionDisplayOff Action = "displayOff"
	ActionSuspend    Action = "suspend"
	ActionHibernate  Action = "hibernate"
	ActionNone       Action = "none"
)

func main() {
	fmt.Println("starting ckidle...")
	src := NewSyntheticEventSource()
	be := NewSyntheticBackend()
	watch(src, be)
}
