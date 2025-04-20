package workflow

import (
	"context"
	"strings"

	"github.com/alirezazeynali75/gad-based-workflow/internal/orchestrator"
)

var LowercaseResultKey = "lowercase"

type Lowercase struct {
	name       string
	sharedData *SharedData
}

func NewLowercase(sharedData *SharedData) orchestrator.Step {
	return &Lowercase{
		name:       "lowercase",
		sharedData: sharedData,
	}
}

func (l *Lowercase) Name() string {
	return l.name
}

func (l *Lowercase) Do(ctx context.Context) error {
	if l.sharedData == nil {
		return ErrInvalidInput
	}
	trimed, ok := (*l.sharedData)[TrimedResultKey].(string)
	if !ok {
		return ErrInvalidInput
	}
	(*l.sharedData)[LowercaseResultKey] = strings.ToLower(trimed)
	return nil
}
