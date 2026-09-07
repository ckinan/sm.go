package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var syntheticIdle int

const EVENT_IDLE_5_MINS = "idle5m"
const EVENT_IDLE_10_MINS = "idle10m"
const EVENT_IDLE_15_MINS = "idle15m"
const EVENT_IDLE_20_MINS = "idle20m"
const EVENT_USER_INPUT = "userinput"
const EVENT_BLOCKED = "blocked"

func main() {
	// event loop that watches:
	// 1) when the machine is idle for X minutes
	// 2) when there is an input from user: keyboard or mouse
	var state = "active"
	for {
		var action string
		time.Sleep(10 * time.Millisecond)
		event := next()

		fmt.Printf("event=%s\n", event)

		// no events to process right now
		if event == "" {
			continue
		}
		state, event = transition(state, event)
		if action != "" && event != EVENT_USER_INPUT {
			execute(action)
			continue
		}
		if event == EVENT_USER_INPUT {
			// when there is activity from the user, the idle time
			// gets reset, so that the state machine goes back again
			syntheticIdle = 0
		}
	}
}

func next() string {
	idleEvent := getIdleEvent()

	fmt.Printf("idleEvent=%s\n", idleEvent)
	if idleEvent != "" {
		return idleEvent
	}
	userInputEvent := getUserInputEvent()
	if userInputEvent != "" {
		return userInputEvent
	}
	return ""
}

func getUserInputEvent() string {
	randomNumber := rand.IntN((20*60)-1+1) + 1
	if randomNumber > syntheticIdle {
		return EVENT_USER_INPUT
	}
	return ""
}

// execute runs the action determined by the state machine
func execute(action string) {
	fmt.Printf("executing action=%s\n", action)
}

func getIdleEvent() string {
	idleSecs := getIdle()
	switch {
	case idleSecs > 20*60:
		return EVENT_IDLE_20_MINS
	case idleSecs > 15*60:
		return EVENT_IDLE_15_MINS
	case idleSecs > 10*60:
		return EVENT_IDLE_10_MINS
	case idleSecs > 5*60:
		return EVENT_IDLE_5_MINS
	}
	return ""
}

// getIdle returns the number of seconds the machine hasn't received any input
// from the user.
// TODO return the idle time from the OS, right now it's returning a synthetic
// value.
func getIdle() int {
	syntheticIdle++
	return syntheticIdle
}

// Transitions calculates the next state of the machine and the action to take
// given the current state of the machine and an event.
// It knows when a session should be locked, if the display should be off or on.
func transition(currentState string, event string) (string, string) {
	switch {
	case currentState == "active" && event == "idle5m":
		return "locked", "lock"
	case currentState == "locked" && event == "idle10m":
		return "displayoff", "displayoff"
	case currentState == "displayoff" && event == "idle15m":
		return "displayoff", "suspend"
	case currentState == "displayoff" && event == "idle20m":
		return "displayoff", "hibernate"
	case currentState == "displayoff" && event == "userinput":
		return "locked", "displayon"
	}
	return currentState, ""
}
