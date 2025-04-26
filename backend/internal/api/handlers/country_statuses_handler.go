package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/w0ikid/world-map-tracker/internal/domain/usecase"
)

type CountryStatusesHandler struct {
	usecase usecase.CountryStatusesUseCaseInterface
}

func NewCountryStatusesHandler(usecase usecase.CountryStatusesUseCaseInterface) *CountryStatusesHandler {
	return &CountryStatusesHandler{
		usecase: usecase,
	}
}

func (h *CountryStatusesHandler) CreateCountryStatus(c *gin.Context) {
	var request struct {
		CountryISO string `json:"country_iso" binding:"required"`
		Status     string `json:"status" binding:"required"`
	}
	userID, _ := c.Get("user_id") 
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	countryStatus, err := h.usecase.CreateCountryStatus(ctx, &usecase.CountryStatusInput{
		UserID:     userID.(int),
		CountryISO: request.CountryISO,
		Status:     request.Status,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user_id":     countryStatus.UserID,
		"country_iso": countryStatus.CountryISO,
		"status":      countryStatus.Status,
	})
}

func (h *CountryStatusesHandler) GetCountryStatuses(c *gin.Context) {
	userID, _ := c.Get("user_id") 
	ctx := c.Request.Context()

	countryStatuses, err := h.usecase.GetCountryStatuses(ctx, userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, countryStatuses)
}

func (h *CountryStatusesHandler) UpdateCountryStatus(c *gin.Context) {
	var request struct {
		CountryISO string `json:"country_iso" binding:"required"`
		Status     string `json:"status" binding:"required"`
	}
	userID, _ := c.Get("user_id") 
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	countryStatus, err := h.usecase.UpdateCountryStatus(ctx, &usecase.CountryStatusInput{
		UserID:     userID.(int),
		CountryISO: request.CountryISO,
		Status:     request.Status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":     countryStatus.UserID,
		"country_iso": countryStatus.CountryISO,
		"status":      countryStatus.Status,
	})
}

func (h *CountryStatusesHandler) DeleteCountryStatus(c *gin.Context) {
	var request struct {
		CountryISO string `json:"country_iso" binding:"required"`
	}

	userID, _ := c.Get("user_id")

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	err := h.usecase.DeleteCountryStatus(ctx, userID.(int), request.CountryISO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *CountryStatusesHandler) GetVisitedPercentage(c *gin.Context) {
	userID, _ := c.Get("user_id") 
	ctx := c.Request.Context()

	visitedCount, err := h.usecase.GetVisitedPercentage(ctx, userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"visited_percentage": visitedCount,
	})
}

func (h *CountryStatusesHandler) GetVisitedCount(c *gin.Context) {
	userID, _ := c.Get("user_id") 
	ctx := c.Request.Context()

	visitedCount, err := h.usecase.GetVisitedCount(ctx, userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"visited_count": visitedCount,
	})
}