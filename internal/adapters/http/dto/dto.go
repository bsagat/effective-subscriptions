package dto

import (
	"fmt"
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

	timeLayout := time.DateOnly
	subs.StartDate, err = time.Parse(timeLayout, subsReq.StartDate)
	if err != nil {
		return domain.Subscription{}, err
	}

	if len(subsReq.EndDate) != 0 {
		subs.EndDate, err = time.Parse(timeLayout, subsReq.EndDate)
		if err != nil {
			return domain.Subscription{}, err
		}
	}

	return subs, nil
}

type SummaryQueries struct {
	Start       time.Time
	End         time.Time
	ServiceName string
	UserID      string
}

func GetSummaryQuery(ctx *gin.Context) (SummaryQueries, error) {
	var (
		sumQuery SummaryQueries
		err      error
	)
	errFormat := "missing required query value %s"
	startStr, ok := ctx.GetQuery("start")
	if !ok {
		return SummaryQueries{}, fmt.Errorf(errFormat, "start")
	}

	endStr, ok := ctx.GetQuery("end")
	if !ok {
		return SummaryQueries{}, fmt.Errorf(errFormat, "end")
	}

	timeLayout := time.DateOnly
	sumQuery.Start, err = time.Parse(timeLayout, startStr)
	if err != nil {
		return SummaryQueries{}, err
	}

	sumQuery.End, err = time.Parse(timeLayout, endStr)
	if err != nil {
		return SummaryQueries{}, err
	}

	sumQuery.UserID, _ = ctx.GetQuery("user_ID")
	sumQuery.ServiceName, _ = ctx.GetQuery("service_name")

	return sumQuery, nil
}
