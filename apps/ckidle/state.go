package main

// transition calculates the next action based on the given event
func transition(event Event) Action {
	switch event{
	case EventIdle5m:
		return ActionLock
	case EventIdle10m:
		return ActionDisplayOff
	case EventIdle15m:
		return ActionSuspend
	case EventIdle20m:
		return ActionHibernate
	case EventUserInput:
		return ActionDisplayOn
	}
	return ActionNone
}
