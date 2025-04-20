package workflow

import (
	"context"

	"github.com/alirezazeynali75/gad-based-workflow/internal/orchestrator"
)


var counter int

type Increase struct {
	name string
}


func NewIncrease() orchestrator.Step {
	return &Increase{
		name: "increase",
	}
}


func (i *Increase) Name() string {
	return i.name
}

func (i *Increase) Do(ctx context.Context) error {
	counter = counter + 1
	return nil
}
