package handlers

import (
	"net/http"

	"github.com/edaywalid/planify-backend/internal/models"
	"github.com/edaywalid/planify-backend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BusinessHandler struct {
	businessService *services.BusinessService
}

func NewBusinessHandler(businessService *services.BusinessService) *BusinessHandler {
	return &BusinessHandler{
		businessService,
	}
}

func (bh *BusinessHandler) GetBusinessById(ctx *gin.Context) {
	businessIDStr := ctx.Param("id")
	businessID, err := uuid.Parse(businessIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid business ID"})
		return
	}

	business, err := bh.businessService.GetBusinessById(businessID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve business"})
		return
	}

	ctx.JSON(http.StatusOK, business)
}

func (bh *BusinessHandler) GetAllBusinesses(ctx *gin.Context) {
	userIDStr := ctx.MustGet("user_id").(string) // Assuming user ID is stored in the context after auth middleware
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	businesses, err := bh.businessService.GetAllBusinesses(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve businesses"})
		return
	}

	ctx.JSON(http.StatusOK, businesses)
}

func (bh *BusinessHandler) AddBusiness(ctx *gin.Context) {
	var business models.Business
	if err := ctx.ShouldBindJSON(&business); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userIDStr := ctx.MustGet("user_id").(string)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	business.UserID = userID
	newBusiness, err := bh.businessService.CreateBusiness(&business)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create business"})
		return
	}

	ctx.JSON(http.StatusCreated, newBusiness)
}

func (bh *BusinessHandler) DeleteBusiness(ctx *gin.Context) {
	businessIDStr := ctx.Param("id")
	businessID, err := uuid.Parse(businessIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid business ID"})
		return
	}

	err = bh.businessService.DeleteBusiness(businessID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete business"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
