package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/w0ikid/world-map-tracker/internal/domain/services/llm" // поправь путь под свой проект
)

type LLMHandler struct {
	groqClient *llm.GroqClient
}

func NewLLMHandler(groqClient *llm.GroqClient) *LLMHandler {
	return &LLMHandler{groqClient: groqClient}
}

func (h *LLMHandler) Ask(c *gin.Context) {
	prompt := c.Query("prompt")
	if prompt == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prompt is required"})
		return
	}

	answer, err := h.groqClient.Chat("llama3-70b-8192", prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": answer})
}
