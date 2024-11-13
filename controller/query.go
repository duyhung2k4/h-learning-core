package controller

import (
	"app/constant"
	"app/dto/request"
	"app/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

type queryController[T any] struct {
	query service.QueryService[T]
}

type QueryController[T any] interface {
	Query(w http.ResponseWriter, r *http.Request)
}

func (c *queryController[T]) Query(w http.ResponseWriter, r *http.Request) {
	var payload request.QueryReq[T]
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	var result interface{}
	var errHandle error

	switch payload.Method {
	case constant.GET:
		result, errHandle = c.query.First(payload)
	case constant.GET_ALL:
		result, errHandle = c.query.Find(payload)
	case constant.CREATE:
		result, errHandle = c.query.Create(payload.Data)
	case constant.MULTI_CREATE:
		result, errHandle = c.query.MultiCreate(payload.Datas)
	case constant.UPDATE:
		result, errHandle = c.query.Update(payload)
	case constant.DELETE:
		result = nil
		errHandle = c.query.Delete(payload)
	}

	if errHandle != nil {
		InternalServerError(w, r, errHandle)
		return
	}

	res := Response{
		Data:    result,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func NewQueryController[T any]() QueryController[T] {
	return &queryController[T]{
		query: service.NewQueryService[T](),
	}
}
