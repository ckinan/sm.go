# ckidle

experimental swayidle replacement

## architecture

### diagram

```
main (the entrypoint)
    |
    v
watcher (the event loop)
    |
    |-> statemachine (states and transitions)
    |-> events (wait and return signals)
    |-> backend (run instructions based on events)
```

### components

main.go: the entry point, it injects the signal and backend instances to the
watcher

watcher.go: the event loop that waits for a signal (an event) and attempts to 
act on it. this is the orchestrator, the one who glues everything

statemachine.go: holds the current state of the program, and knows about the
transition of the states based on the events

events.go: responsible for receive notifications for specific events that
indicate idle time and user activity
- e.g. idle for 5m, 10m, etc..
- e.g. user input
- e.g. otherwise just wait

backend.go: the resource that gets the instructions from the event loop
- display off
- lock screen
- suspend
- hibernate
- wake up

### external components
events.go and backend.go need to establish communication with the Operating
System and external tools.

- events.go: talks to Linux to be notified about user input and idle duration
- backend.go: talks to Linux to display off, suspend, hibernate, wake up. also
talks to swaylock to lock the screen.

they should meet interface contracts so that synthetic backends and events can
be used for unit testing

