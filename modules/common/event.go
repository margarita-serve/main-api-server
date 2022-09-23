package common

// Event Interface for Domain Event
type Event interface {
	Name() string
}

// GeneralError actual event
type GeneralError string

func NewGeneralError(err error) Event {
	return GeneralError(err.Error())
}

func (e GeneralError) Name() string {
	return "event.generalerror"
}

type EventHandler interface {
	Update(event Event)
}

type EventPublisher struct {
	handlers map[string][]EventHandler
}

func (e *EventPublisher) Subscribe(handler EventHandler, events ...Event) {
	for _, event := range events {
		handlers := e.handlers[event.Name()]
		handlers = append(handlers, handler)
		e.handlers[event.Name()] = handlers
	}
}

func (e *EventPublisher) Notify(event Event) {
	for _, handler := range e.handlers[event.Name()] {
		handler.Update(event)
	}
}

func NewEventPublisher() EventPublisher {
	return EventPublisher{make(map[string][]EventHandler)}
}
