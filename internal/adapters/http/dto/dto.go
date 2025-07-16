package dto

import (
	"submanager/internal/domain"
	"time"

	"github.com/gin-gonic/gin"
)

type subsReq struct {
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserID      string `json:"user_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

func GetSubsJSON(ctx *gin.Context) (domain.Subscription, error) {
	var (
		subsReq subsReq
		err     error
	)
	if err := ctx.BindJSON(&subsReq); err != nil {
		return domain.Subscription{}, err
	}

	subs := domain.Subscription{
		ServiceName: subsReq.ServiceName,
		Price:       subsReq.Price,
		UserID:      subsReq.UserID,
	}

	subs.StartDate, err = time.Parse(time.DateOnly, subsReq.StartDate)
	if err != nil {
		return domain.Subscription{}, err
	}

	if len(subsReq.EndDate) != 0 {
		subs.EndDate, err = time.Parse(time.DateOnly, subsReq.EndDate)
		if err != nil {
			return domain.Subscription{}, err
		}
	}

	return subs, nil
}
