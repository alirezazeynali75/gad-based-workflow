package orchestrator

import (
	"context"
	"log/slog"

	"github.com/Azure/go-workflow"
)

type Step interface {
	Do(context.Context) error
	Name() string
}

type Orchestrator struct {
	logger *slog.Logger
	steps  []flow.AddStep[Step]
	w      *flow.Workflow
}

func NewOrchestrator(logger *slog.Logger, steps ...flow.AddStep[Step]) *Orchestrator {
	return &Orchestrator{
		logger: logger.With("orchestrator"),
		steps:  steps,
	}
}


func (o *Orchestrator) Build() {
	logger := o.logger.With("build")
	logger.Debug("building orchestrator")
	w := new(flow.Workflow)
	for _, step := range o.steps {
		w.Add(step)
	}
	o.w = w
}

func (o *Orchestrator) Run(ctx context.Context) error {
	logger := o.logger.With("run")
	logger.Debug("running orchestrator")
	if o.w == nil {
		return ErrWorkflowNotBuilt
	}
	if err := o.w.Do(ctx); err != nil {
		logger.Error("running orchestrator error", "err", err)
		return err
	}
	return nil
}
