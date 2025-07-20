package dto

import (
	"errors"
	"fmt"
	"strconv"
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

// GetSubsJSON extracts subscription data from the request context.
// It returns a domain.Subscription object or an error if the data is invalid.
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
	PageNumber  int
	PageSize    int
}

// GetSummaryQuery extracts summary query parameters from the request context.
// It returns a SummaryQueries struct or an error if required parameters are missing or invalid.
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
	// Использование:
	sumQuery.PageNumber, sumQuery.PageSize, err = GetPaginationArgs(ctx)
	if err != nil {
		return sumQuery, err
	}
	return sumQuery, nil
}

func GetPaginationArgs(ctx *gin.Context) (int, int, error) {
	pageNumberStr := ctx.DefaultQuery("page_number", "1") // по умолчанию 1
	pageSizeStr := ctx.DefaultQuery("page_size", "10")    // по умолчанию 10

	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil {
		return 0, 0, errors.New("page_number must be integer")
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return 0, 0, errors.New("page_size must be integer")
	}

	if pageNumber < 1 {
		return 0, 0, errors.New("page_number must be greater than 0")
	}

	if pageSize < 1 || pageSize > 100 {
		return 0, 0, errors.New("page_size must be betqeen 1 and 100")
	}

	return pageNumber, pageSize, nil
}
