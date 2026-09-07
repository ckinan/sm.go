package main

import "fmt"

func watch(src EventSource, be Backend) {
	for {
		event := src.nextEvent()
		if event != EventNone {
			fmt.Printf("event=%s\n", event)
		}
		action := transition(event)
		if action != ActionNone {
			be.execute(action)
		}
	}
}
