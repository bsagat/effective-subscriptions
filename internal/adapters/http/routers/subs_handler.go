package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SubsHandler handles subscription CRUDL routes.
type SubsHandler struct{}

func NewSubsHandler() *SubsHandler {
	return &SubsHandler{}
}

// RegisterSubsRoutes registers all subs http operations
func (h *SubsHandler) RegisterSubsRoutes(r *gin.RouterGroup) {
	r.POST("/", h.CreateSubsHandler)
	r.GET("/:user_id/:service_name", h.GetSubsHandler)
	r.GET("/:user_id", h.ListSubsHandler)
	r.PUT("/", h.UpdateSubsHandler)
	r.DELETE("/:user_id/:service_name", h.DeleteSubsHandler)
	r.DELETE("/:user_id", h.DeleteSubsListHandler)
}

// CreateSubsHandler creates a new subscription.
func (h *SubsHandler) CreateSubsHandler(ctx *gin.Context) {
	// TODO: Parse JSON body, validate, call service to create subscription
	ctx.JSON(http.StatusCreated, gin.H{"message": "Subscription created"})
}

// GetSubsHandler returns user subscription by specific service
func (h *SubsHandler) GetSubsHandler(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	serviceName := ctx.Param("service_name")

	// TODO: Validate userID, fetch subscription from DB/service by userID and serviceName
	ctx.JSON(http.StatusOK, gin.H{
		"user_id":      userID,
		"service_name": serviceName,
	})
}

// ListSubsHandler returns user subscriptions list
func (h *SubsHandler) ListSubsHandler(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	// TODO: Validate userID, fetch all subscriptions for userID from DB/service
	ctx.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"subs":    []string{},
	})
}

// UpdateSubsHandler updates subscription.
func (h *SubsHandler) UpdateSubsHandler(ctx *gin.Context) {
	// TODO: Parse JSON body, validate, call service to update subscription
	ctx.JSON(http.StatusOK, gin.H{"message": "Subscription updated"})
}

// DeleteSubsHandler deletes user subscription by specific service
func (h *SubsHandler) DeleteSubsHandler(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	serviceName := ctx.Param("service_name")

	// TODO: Validate, call service to delete specific subscription
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "Subscription deleted",
		"user_id":      userID,
		"service_name": serviceName,
	})
}

// DeleteSubsListHandler deletes all user subscriptions
func (h *SubsHandler) DeleteSubsListHandler(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	// TODO: Validate, call service to delete all subscriptions for userID
	ctx.JSON(http.StatusOK, gin.H{
		"message": "All subscriptions deleted",
		"user_id": userID,
	})
}
