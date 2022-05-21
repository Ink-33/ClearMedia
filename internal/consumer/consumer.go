package consumer

import (
	"github.com/Ink-33/ClearMedia/internal/order"
)

// Consumer is the interface that wraps the Consume method.
type Consumer interface {
	Consume(order.Order) error
}

// Func is a function that implements the Consumer interface.
type Func func(order.Order) error

// Consume calls f(order).
func (f Func) Consume(order order.Order) error {
	return f(order)
}

// NewConsumer returns a ConsumerFunc that calls f(order).
func NewConsumer(f Func) Consumer {
	return f
}
