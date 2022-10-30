package rate

import (
	"context"
	"net/http"
	"reflect"
	"time"

	sv "github.com/core-go/core"
)

type RateHandler interface {
	Load(w http.ResponseWriter, r *http.Request)
	Rate(w http.ResponseWriter, r *http.Request)
}

func NewRateHandler(
	service RateService,
	status sv.StatusConfig,
	logError func(context.Context, string, ...map[string]interface{}),
	validate func(ctx context.Context, model interface{}) ([]sv.ErrorMessage, error),
	action *sv.ActionConfig,
) RateHandler {
	modelType := reflect.TypeOf(Rate{})
	params := sv.CreateParams(modelType, &status, logError, validate, action)
	return &rateHandler{service: service, Params: params}
}

type rateHandler struct {
	service RateService
	*sv.Params
}

func (h *rateHandler) Load(w http.ResponseWriter, r *http.Request) {
	author := sv.GetRequiredParam(w, r)
	id := sv.GetRequiredParam(w, r, 1)
	if len(id) > 0 {
		res, err := h.service.Load(r.Context(), id, author)
		sv.RespondModel(w, r, res, err, h.Error, nil)
	}
}

func (h *rateHandler) Rate(w http.ResponseWriter, r *http.Request) {
	var rate Rate
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
