package criteria

import (
	"context"
	"net/http"
	"reflect"
	"time"

	sv "github.com/core-go/core"
)

type RateCriteriaHandler interface {
	Rate(w http.ResponseWriter, r *http.Request)
}

func NewRateCriteriaHandler(
	service RateCriteriaService,
	status sv.StatusConfig,
	logError func(context.Context, string, ...map[string]interface{}),
	validate func(ctx context.Context, model interface{}) ([]sv.ErrorMessage, error),
	action *sv.ActionConfig,
) RateCriteriaHandler {
	modelType := reflect.TypeOf(RateCriteria{})
	params := sv.CreateParams(modelType, &status, logError, validate, action)
	return &rateCriteriaHandler{service: service, Params: params}
}

type rateCriteriaHandler struct {
	service RateCriteriaService
	*sv.Params
}

func (h *rateCriteriaHandler) Rate(w http.ResponseWriter, r *http.Request) {
	var rate RateCriteria
	var t = time.Now()
	rate.Time = &t
	er1 := sv.Decode(w, r, &rate)
	rate.Author = sv.GetRequiredParam(w, r)
	rate.Id = sv.GetRequiredParam(w, r, 1)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &rate)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Create) {
			result, er3 := h.service.Rate(r.Context(), &rate)
			sv.AfterCreated(w, r, &rate, result, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}

}
