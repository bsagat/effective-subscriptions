package routers

import (
	"net/http"
	"submanager/internal/adapters/http/dto"
	"submanager/internal/domain"
	"submanager/internal/pkg/httputils"
	"submanager/internal/pkg/logger"

	"github.com/gin-gonic/gin"
)

// SubsHandler handles subscription CRUDL routes.
type SubsHandler struct {
	serv domain.SubsService
	log  logger.Logger
}

func NewSubsHandler(serv domain.SubsService, log logger.Logger) *SubsHandler {
	return &SubsHandler{
		serv: serv,
		log:  log,
	}
}

// RegisterSubsRoutes registers all subs http operations
func (h *SubsHandler) RegisterSubsRoutes(r *gin.RouterGroup) {
	r.POST("/", h.CreateSubsHandler)
	r.GET("/:user_id/:service_name", h.GetSubsHandler)
	r.GET("/:user_id", h.ListSubsHandler)
	r.GET("/summary", h.SummaryHandler)
	r.PUT("/", h.UpdateSubsHandler)
	r.DELETE("/:user_id/:service_name", h.DeleteSubsHandler)
	r.DELETE("/:user_id", h.DeleteSubsListHandler)
}

// CreateSubsHandler creates a new subscription.
func (h *SubsHandler) CreateSubsHandler(ctx *gin.Context) {
	subs, err := dto.GetSubsJSON(ctx)
	if err != nil {
		h.log.Error("Failed to bind subscription JSON request", "error", err)
		httputils.SendError(ctx, http.StatusBadRequest, domain.ErrInvalidJSON)
		return
	}

	if err := validateSubs(subs); err != nil {
		h.log.Error("Failed to create subscription", "error", err)
		httputils.SendError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := h.serv.CreateSubscription(ctx.Request.Context(), subs); err != nil {
		h.log.Error("Failed to create subscription", "error", err)
		httputils.SendError(ctx, httputils.GetStatus(err), err)
		return
	}

	httputils.SendMessage(ctx, http.StatusCreated, "Subscription created")
}

// GetSubsHandler returns user subscription by specific service
func (h *SubsHandler) GetSubsHandler(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	serviceName := ctx.Param("service_name")

	if err := validateSubsParams(serviceName, userID); err != nil {
		h.log.Error("Failed to get subscription", "error", err)
		httputils.SendError(ctx, http.StatusBadRequest, err)
		return
	}

	subs, err := h.serv.GetSubscription(ctx.Request.Context(), serviceName, userID)
	if err != nil {
		h.log.Error("Failed to get subscription", "error", err)
		httputils.SendError(ctx, httputils.GetStatus(err), err)
		return
	}

	ctx.JSON(http.StatusOK, subs)
}

// ListSubsHandler returns user subscriptions list
func (h *SubsHandler) ListSubsHandler(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	if err := validateSubsParams("notempty", userID); err != nil {
		h.log.Error("Failed to get subscription list", "error", err)
		httputils.SendError(ctx, http.StatusBadRequest, err)
		return
	}

	list, err := h.serv.GetSubscriptionList(ctx, userID)
	if err != nil {
		h.log.Error("Failed to get subscription list", "error", err)
		httputils.SendError(ctx, httputils.GetStatus(err), err)
		return
	}

	ctx.JSON(http.StatusOK, list)
}

// UpdateSubsHandler updates subscription.
func (h *SubsHandler) UpdateSubsHandler(ctx *gin.Context) {
	subs, err := dto.GetSubsJSON(ctx)
	if err != nil {
		h.log.Error("Failed to bind subscription JSON request", "error", err)
		httputils.SendError(ctx, http.StatusBadRequest, domain.ErrInvalidJSON)
		return
	}

	if err := validateSubs(subs); err != nil {
		h.log.Error("Failed to update subscription", "error", err)
		httputils.SendError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := h.serv.UpdateSubscription(ctx, subs); err != nil {
		h.log.Error("Failed to update subscription", "error", err)
		httputils.SendError(ctx, httputils.GetStatus(err), err)
		return
	}

	httputils.SendMessage(ctx, http.StatusOK, "Subscription updated")
}

// DeleteSubsHandler deletes user subscription by specific service
func (h *SubsHandler) DeleteSubsHandler(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	serviceName := ctx.Param("service_name")

	if err := validateSubsParams(serviceName, userID); err != nil {
		h.log.Error("Failed to delete subscription", "error", err)
		httputils.SendError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := h.serv.DeleteSubscription(ctx, serviceName, userID); err != nil {
		h.log.Error("Failed to delete subscription", "error", err)
		httputils.SendError(ctx, httputils.GetStatus(err), err)
		return
	}

	httputils.SendMessage(ctx, http.StatusOK, "Subscription deleted")
}

// DeleteSubsListHandler deletes all user subscriptions
func (h *SubsHandler) DeleteSubsListHandler(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	if err := validateSubsParams("notempty", userID); err != nil {
		h.log.Error("Failed to get subscription list", "error", err)
		httputils.SendError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := h.serv.DeleteSubscriptionList(ctx, userID); err != nil {
		h.log.Error("Failed to delete subscription list", "error", err)
		httputils.SendError(ctx, httputils.GetStatus(err), err)
		return
	}

	httputils.SendMessage(ctx, http.StatusOK, "All user subscriptions deleted")
}

// SummaryHandler retrieves a summary of subscriptions based on the filter.
func (h *SubsHandler) SummaryHandler(ctx *gin.Context) {
	summQuery, err := dto.GetSummaryQuery(ctx)
	if err != nil {
		h.log.Error("Failed to get summary queries", "error", err)
		httputils.SendError(ctx, http.StatusBadRequest, err)
		return
	}

	summResp, err := h.serv.GetSummaryByFilter(ctx, summQuery.Start, summQuery.End, summQuery.ServiceName, summQuery.UserID)
	if err != nil {
		h.log.Error("Failed to get summary by filter", "error", err)
		httputils.SendError(ctx, httputils.GetStatus(err), err)
		return
	}

	ctx.JSON(http.StatusOK, summResp)
}
