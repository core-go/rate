package search

import (
	"context"
	"net/http"
	"reflect"

	"github.com/core-go/search"
)

type RateSearchHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
}

func NewRateSearchHandler(
	find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error),
	logError func(context.Context, string, ...map[string]interface{}),
) RateSearchHandler {
	searchModelType := reflect.TypeOf(RateFilter{})
	modelType := reflect.TypeOf(Rate{})
	var writeLog func(context.Context, string, string, bool, string) error
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	return &searchRateHandler{SearchHandler: searchHandler}
}

func NewRateCriteriaSearchHandler(
	find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error),
	logError func(context.Context, string, ...map[string]interface{}),
) RateSearchHandler {
	searchModelType := reflect.TypeOf(RateFilter{})
	modelType := reflect.TypeOf(RateCriteria{})
	var writeLog func(context.Context, string, string, bool, string) error
	searchHandler := search.NewSearchHandler(find, modelType, searchModelType, logError, writeLog)
	return &searchRateHandler{SearchHandler: searchHandler}
}

type searchRateHandler struct {
	*search.SearchHandler
}
