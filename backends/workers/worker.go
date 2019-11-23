package workers

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/relistan/go-director"
)

type Worker struct {
	ID             uuid.UUID
	WorkChannel    chan string
	DefaultContext context.Context
	Looper         *director.FreeLooper
}

func NewWorker(id uuid.UUID, workChannel chan string, defaultContext context.Context) Worker {
	// Create, and return the worker.
	worker := Worker{
		ID:             id,
		WorkChannel:    workChannel,
		DefaultContext: defaultContext,
		Looper:         director.NewFreeLooper(director.FOREVER, make(chan error)),
	}

	return worker
}

func (w *Worker) Run() {
	w.Looper.Loop(func() error {
		select {
		case work := <-w.WorkChannel:
			// Receive a work request.
			fmt.Printf("Here is my work %+v", work)
			fmt.Printf("worker: %s", w.ID.String())

		case <-w.DefaultContext.Done():
			// We have been asked to stop.
			fmt.Printf("Stopping worker %d", w.ID)
			w.Looper.Quit()
		}
		return nil
	})
}
