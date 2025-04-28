package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/w0ikid/world-map-tracker/internal/domain/usecase"
)

type CountryStatusesHandler struct {
	usecase usecase.CountryStatusesUseCaseInterface
	user usecase.UserUseCaseInterface
}

func NewCountryStatusesHandler(usecase usecase.CountryStatusesUseCaseInterface, user usecase.UserUseCaseInterface) *CountryStatusesHandler {
	return &CountryStatusesHandler{
		usecase: usecase,
		user: user,
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

func (h *CountryStatusesHandler) GetVisitedPercentageUsername(c *gin.Context){
	username := c.Param("username")

	user, err := h.user.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	visitedPercantage, err := h.usecase.GetVisitedPercentage(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"visited_percentage": visitedPercantage,
	})
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

func (h *CountryStatusesHandler) GetVisitedCountUsername(c *gin.Context){
	username := c.Param("username")

	user, err := h.user.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	visitedCount, err := h.usecase.GetVisitedCount(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"visited_count": visitedCount,
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

func (h *CountryStatusesHandler) GetWishListCountUsername(c *gin.Context){
	username := c.Param("username")
	user, err := h.user.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wishListCount, err := h.usecase.GetWishListCount(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"wish_list_count": wishListCount,
	})
}


func (h *CountryStatusesHandler) GetWishListCount(c *gin.Context) {
	userID, _ := c.Get("user_id") 
	ctx := c.Request.Context()
	wishListCount, err := h.usecase.GetWishListCount(ctx, userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"wish_list_count": wishListCount,
	})
}

func (h *CountryStatusesHandler) GetTopFiveVisitedCountries(c *gin.Context) {
	ctx := c.Request.Context()

	topCountries, err := h.usecase.GetTopFiveVisitedCountries(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, topCountries)
}



func (h *CountryStatusesHandler) GetTopFiveWishlistCountries(c *gin.Context) {
	ctx := c.Request.Context()

	topCountries, err := h.usecase.GetTopFiveWishlistCountries(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, topCountries)
}


func (h *CountryStatusesHandler) GetUsersWithSimilarList(c *gin.Context) {
	userID, _ := c.Get("user_id") 
	ctx := c.Request.Context()

	users, err := h.usecase.FindUsersWithSimilarList(ctx, userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}