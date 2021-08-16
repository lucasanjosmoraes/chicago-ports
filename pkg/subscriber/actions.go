package subscriber

import "context"

// Actions allows devs to trig actions during acnkownledge or reject routines.
type Actions interface {
	BeforeAck(ctx context.Context)
	AfterAck(ctx context.Context)
	BeforeRetry(ctx context.Context)
	AfterRetry(ctx context.Context)
	BeforeSendToDL(ctx context.Context)
	AfterSendToDL(ctx context.Context)
}

// DoNothing implements Actions with empty methods to allow devs to use it to avoid check if actions
// are nil before use them.
var DoNothing = doNothing{}

type doNothing struct{}

func (d doNothing) BeforeAck(_ context.Context) {}

func (d doNothing) AfterAck(_ context.Context) {}

func (d doNothing) BeforeRetry(_ context.Context) {}

func (d doNothing) AfterRetry(_ context.Context) {}

func (d doNothing) BeforeSendToDL(_ context.Context) {}

func (d doNothing) AfterSendToDL(_ context.Context) {}
