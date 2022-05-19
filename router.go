package router

import (
	"bytes"
	"context"
	"github.com/SaCavid/router/models"
	"html/template"
	"net/http"
)

type Router struct {
	R        *models.LambdaRequest
	W        *models.LambdaResponse
	Ctx      context.Context
	RouteMap map[string]map[string]func(ctx context.Context, event models.LambdaRequest) (models.LambdaResponse, error)
}

func NewLambdaRouter(ctx context.Context, event *models.LambdaRequest) (router Router) {

	router.R = event
	router.W = &models.LambdaResponse{
		Headers: make(map[string]string),
	}
	router.Ctx = ctx
	router.RouteMap = make(map[string]map[string]func(ctx context.Context, event models.LambdaRequest) (models.LambdaResponse, error))

	return
}

func (r Router) AllowedMethods(methods ...string) *Router {

	for _, v := range methods {
		newMethod := make(map[string]func(ctx context.Context, event models.LambdaRequest) (models.LambdaResponse, error))
		r.RouteMap[v] = newMethod
	}

	return &r
}

func (r Router) Handler(method, path string, f func(ctx context.Context, event models.LambdaRequest) (models.LambdaResponse, error)) {
	r.RouteMap[method][path] = f
}

func (r Router) Run() (models.LambdaResponse, error) {

	if route, ok := r.RouteMap[r.R.RequestContext.Http.Method][r.R.RequestContext.Http.Path]; ok {
		return route(r.Ctx, *r.R)
	}

	r.W.StatusCode = http.StatusNotFound
	return *r.W, nil
}

func (r Router) Execute(name, path string, data any) (string, error) {
	t := template.New(name)

	var err error
	t, err = t.ParseFiles(path)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
