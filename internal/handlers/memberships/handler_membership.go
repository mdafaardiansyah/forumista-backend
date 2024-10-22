package memberships

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mdafaardiansyah/forumista-backend/internal/model/memberships"
)

type membeshipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
}

type Handler struct {
	*gin.Engine

	membershipSvc membeshipService
}

func NewHandler(api *gin.Engine, membershipSvc membeshipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/memberships")
	route.GET("/ping", h.Ping)
	route.POST("/sign-up", h.SignUp)
}
