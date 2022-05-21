package factory

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/Ink-33/ClearMedia/internal/consumer"
	"github.com/Ink-33/ClearMedia/internal/order"
	"github.com/pkg/errors"
)

const (
	// DoAudioStripe is the action number for audioStriper.
	DoAudioStripe = iota
)

// MaxWorkers is the maximum number of workers that can be used.
var MaxWorkers = runtime.NumCPU()

// Do calls the function f(order) for each order in the slice of orders.
func Do(files []string, output string, action int) (err error) {
	var worker consumer.Consumer
	var createOrder func(string, string) any
	switch action {
	case DoAudioStripe:
		worker = consumer.NewConsumer(audioStriper)
		createOrder = func(name, output string) any {
			return &order.StripeOrder{Name: name, Output: output}
		}
	default:
		return errors.Errorf("unknown action number : %d", action)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(files))

	bucket := make(chan struct{}, MaxWorkers)
	for i := 0; i < MaxWorkers; i++ {
		bucket <- struct{}{}
	}

	for _, file := range files {
		<-bucket
		go func(file string) {
			defer func() {
				wg.Done()
				bucket <- struct{}{}
			}()
			fmt.Println("Processing file :", file)
			err := worker.Consume(createOrder(file, output).(order.Order))
			if err != nil {
				fmt.Printf("Error processing file %s : %s\n", file, err.Error())
			}
		}(file)
	}

	wg.Wait()
	return nil
}
