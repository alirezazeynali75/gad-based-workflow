package api

import (
	"net/http"

	flow "github.com/Azure/go-workflow"
	"github.com/alirezazeynali75/gad-based-workflow/internal/orchestrator"
	"github.com/alirezazeynali75/gad-based-workflow/internal/workflow"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) Trim(c *gin.Context) {
    // Parse input from the request
    var input struct {
        Text string `json:"text" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        h.logger.Error("invalid input", "error", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    // Prepare shared data
    sharedData := workflow.SharedData{
        workflow.InputKey: input.Text,
    }

    // Create a new Trim workflow step
    trimStep := workflow.NewTrim(&sharedData)

		o := orchestrator.NewOrchestrator(h.logger, flow.Step(trimStep))
		o.Build()

    // Execute the Trim workflow step
    if err := o.Run(c.Request.Context()); err != nil {
        h.logger.Error("failed to execute trim workflow", "error", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process input"})
        return
    }

    // Retrieve the result from shared data
    trimmedText, ok := sharedData[workflow.TrimedResultKey].(string)
    if !ok {
        h.logger.Error("failed to retrieve trimmed result")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process input"})
        return
    }

    // Return the result
    c.JSON(http.StatusOK, gin.H{"trimmed_text": trimmedText})
}