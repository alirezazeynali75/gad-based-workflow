package api

import (
	"fmt"
	"net/http"

	flow "github.com/Azure/go-workflow"
	"github.com/alirezazeynali75/gad-based-workflow/internal/orchestrator"
	"github.com/alirezazeynali75/gad-based-workflow/internal/workflow"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) All(c *gin.Context) {
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

    // Create a new Uppercase workflow step
    uppercaseStep := workflow.NewUppercase(&sharedData)
		trimStep := workflow.NewTrim(&sharedData)
		increaseStep := workflow.NewIncrease()
		lowercaseStep := workflow.NewLowercase(&sharedData) 
		reverseStep := workflow.NewReverse(&sharedData)

		o := orchestrator.NewOrchestrator(h.logger, flow.Step(trimStep), flow.Step(uppercaseStep).DependsOn(trimStep), flow.Step(lowercaseStep).DependsOn(trimStep), flow.Step(reverseStep), flow.Step(increaseStep))
		o.Build()

    // Execute the Uppercase workflow step
    if err := o.Run(c.Request.Context()); err != nil {
        h.logger.Error("failed to execute uppercase workflow", "error", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process input"})
        return
    }

		fmt.Printf("sharedData: %v\n", sharedData)

    // Retrieve the result from shared data
    uppercasedText, ok := sharedData[workflow.UppercaseKey].(string)
    if !ok {
        h.logger.Error("failed to retrieve uppercased result")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process input"})
        return
    }

		lowercaseText, ok := sharedData[workflow.LowercaseResultKey].(string)
    if !ok {
        h.logger.Error("failed to retrieve lowercase result")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process input"})
        return
    }

		reverse, ok := sharedData[workflow.ReverseResultKey].(string)
    if !ok {
        h.logger.Error("failed to retrieve reverse result")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process input"})
        return
    }

    // Return the result
    c.JSON(http.StatusOK, gin.H{"uppercased_text": uppercasedText, "lowercased_text": lowercaseText, "reversed_text": reverse})
}