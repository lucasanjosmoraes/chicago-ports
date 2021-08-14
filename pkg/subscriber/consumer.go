package subscriber

import (
	"context"

	"github.com/lucasanjosmoraes/chicago-ports/pkg/stoppage"
)

// Consumer defines everything that an adapter needs to provide to every queue/topic consumer.
type Consumer interface {
	Consume(ctx context.Context, handler HandleFunc) error
	stoppage.Stopper
}

// HandleFunc is a callback func called by adapters to pass the message
// and a acknowledge func.
//
// Ack func SHOULD be called when message process is done. If adapters
// returns error on acknowledge, you should rollback your transaction.
type HandleFunc func(Message, Ack, Reject)

// Message represents a event entity.
type Message interface {
	Value() []byte
	Headers() map[string]string
	Subject() string
}

// Ack is where the adapters will handle its ack routines.
type Ack func(ctx context.Context) error

// Reject is where the adapters will handle its no-ack/reject/retry routines.
type Reject func(ctx context.Context, err error)
