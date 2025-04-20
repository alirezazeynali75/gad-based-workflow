package workflow

import (
	"context"
	"strings"

	"github.com/alirezazeynali75/gad-based-workflow/internal/orchestrator"
)

var TrimedResultKey = "trimed"


type Trim struct {
	name string
	sharedData *SharedData
}

func NewTrim(sharedData *SharedData) orchestrator.Step {
	return &Trim{
		name:       "trim",
		sharedData: sharedData,
	}
}
func (t *Trim) Name() string {
	return t.name
}
func (t *Trim) Do(ctx context.Context) error {
	if t.sharedData == nil {
		return ErrInvalidInput
	}
	input, ok := (*t.sharedData)[InputKey].(string)
	if !ok {
		return ErrInvalidInput
	}
	trimed := strings.TrimSpace(input)
	(*t.sharedData)[TrimedResultKey] = trimed
	return nil
}