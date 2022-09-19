package main

import (
	"context"
	"log"
	"net/http"

	"github.com/SaCavid/router"
	"github.com/SaCavid/router/models"
	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	log.Println("SS Serverless Server")
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

	return s.Router.Response(nil)
}
