package router

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"html/template"
	"net/http"
	"regexp"
	"strings"

	"github.com/SaCavid/router/models"
	"github.com/rs/zerolog/log"
)

type Router struct {
	R        *models.LambdaRequest
	W        *models.LambdaResponse
	Ctx      context.Context
	RouteMap map[string]map[string]Urls
}

type Urls struct {
	Path  string
	Regex string
	Vars  map[string]string
	F     func(ctx context.Context, event models.LambdaRequest) (models.LambdaResponse, error)
}

func NewLambdaRouter(ctx context.Context, event *models.LambdaRequest) (router Router) {

	router.R = event
	router.W = &models.LambdaResponse{
		Headers: make(map[string]string),
	}
	router.Ctx = ctx
	router.RouteMap = make(map[string]map[string]Urls)

	return
}

func (r *Router) AllowedMethods(methods ...string) {

	for _, v := range methods {
		newMethod := make(map[string]Urls)
		r.RouteMap[v] = newMethod
	}
}

func (r Router) Handler(method, path string, f func(ctx context.Context, event models.LambdaRequest) (models.LambdaResponse, error)) {
	url := Urls{
		Path: path,
		Vars: make(map[string]string),
		F:    f,
	}

	pathArr := strings.Split(path, "/")
	pathArr = pathArr[1:]
	regexString := "^"
	for _, val := range pathArr {
		matched, _ := regexp.MatchString("{([0-9A-Za-z_-]+)}", val)
		if matched {
			regexString += "/([0-9A-Za-z_-]+)"
			val = strings.Replace(val, "{", "", -1)
			val = strings.Replace(val, "}", "", -1)
			url.Vars[val] = ""
		} else {
			regexString += "/" + val
		}
	}
	regexString += "$"
	url.Regex = regexString

	r.RouteMap[method][path] = url
}

func (r Router) Run() (models.LambdaResponse, error) {
	// decode base64 body to string
	data, err := base64.StdEncoding.DecodeString(r.R.Body)
	if err != nil {
		log.Fatal().Msg("error: " + err.Error())
	}

	r.R.Body = string(data)

	for _, url := range r.RouteMap[r.R.RequestContext.HTTP.Method] {
		matched, _ := regexp.MatchString(url.Regex, r.R.RequestContext.HTTP.Path)
		if matched {
			pathArr := strings.Split(url.Path, "/")
			pathArr = pathArr[1:]
			realArr := strings.Split(r.R.RequestContext.HTTP.Path, "/")
			realArr = realArr[1:]
			for key, val := range pathArr {
				matched, _ := regexp.MatchString("{([0-9A-Za-z_-]+)}", val)
				if matched {
					val = strings.Replace(val, "{", "", -1)
					val = strings.Replace(val, "}", "", -1)
					if _, ok := url.Vars[val]; ok {
						url.Vars[val] = realArr[key]
					} else {
						log.Fatal().Msg("Key not exists. Router Error")
					}
				}
			}

			return url.F(r.Ctx, *r.R)
		}
	}

	r.W.StatusCode = http.StatusNotFound
	return r.Response(nil)
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

func (r Router) Middleware() Router {

	return r
}

func (r Router) Response(err error) (models.LambdaResponse, error) {
	return *r.W, err
}

func (r Router) BindJson(d any) error {

	err := json.Unmarshal([]byte(r.R.Body), &d)
	if err != nil {
		return err
	}

	return nil
}
