package main

import (
	"context"
	"github.com/SaCavid/router"
	"github.com/SaCavid/router/models"
	"log"
	"net/http"
)

func init() {

}

func main() {
	lambda.Start(lambdaHandler)
}

type Server struct {
	Router router.Router
}

func lambdaHandler(ctx context.Context, event models.LambdaRequest) {

	s := Server{}

	s.Router = router.NewLambdaRouter(ctx, &event)
	s.Router.W.Set("Content-type", "application/json")

	s.Router.Handler("GET", "/", s.Ping)

	log.Fatalln(s.Router.Run())
}

func (s *Server) Ping(ctx context.Context, r models.LambdaRequest) (models.LambdaResponse, error) {

	s.Router.W.Body = "Pong"
	s.Router.W.StatusCode = http.StatusOK

	return *s.Router.W, nil
}
