package workflow

import (
	"context"
	"strings"

	"github.com/alirezazeynali75/gad-based-workflow/internal/orchestrator"
)

var UppercaseKey = "uppercase"


type Uppercase struct {
	name string
	sharedData *SharedData
}
func NewUppercase(sharedData *SharedData) orchestrator.Step {
	return &Uppercase{
		name:       "uppercase",
		sharedData: sharedData,
	}
}
func (u *Uppercase) Name() string {
	return u.name
}
func (u *Uppercase) Do(ctx context.Context) error {
	if u.sharedData == nil {
		return ErrInvalidInput
	}
	trimed, ok := (*u.sharedData)[TrimedResultKey].(string)
	if !ok {
		return ErrInvalidInput
	}
	(*u.sharedData)[UppercaseKey] = strings.ToUpper(trimed)
	return nil
}