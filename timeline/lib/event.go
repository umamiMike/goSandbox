package event

type event struct {
	id      int
	payload object
	next    Event
	prev    Event
}
