package publisher

import (
	"context"

	"github.com/lucasanjosmoraes/chicago-ports/pkg/stoppage"
)

// EventHeader contains meta information about an event.
type EventHeader struct {
	Key   string
	Value []byte
}

// Event is the main entity produced by a producer.
type Event struct {
	Key     []byte
	Value   []byte
	Headers []EventHeader
}

// Producer defines everything that an adapter needs to provide to every queue/topic producer.
type Producer interface {
	WriteEvent(ctx context.Context, topic string, event Event) error
	WriteEvents(ctx context.Context, topic string, events []Event) error
	stoppage.Stopper
}
