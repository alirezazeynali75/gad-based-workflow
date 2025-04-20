package workflow

import (
	"context"

	"github.com/alirezazeynali75/gad-based-workflow/internal/orchestrator"
)

var ReverseResultKey = "reversed"

type Reverse struct {
	name       string
	sharedData *SharedData
}


func NewReverse(sharedData *SharedData) orchestrator.Step {
	return &Reverse{
		name:       "reverse",
		sharedData: sharedData,
	}
}

func (l *Reverse) Name() string {
	return l.name
}


func (l *Reverse) Do(ctx context.Context) error {
	if l.sharedData == nil {
		return ErrInvalidInput
	}
	input, ok := (*l.sharedData)[InputKey].(string)
	if !ok {
		return ErrInvalidInput
	}
	reversed := ""
	for _, r := range input {
		reversed = string(r) + reversed
	}
	(*l.sharedData)[ReverseResultKey] = reversed
	return nil
}