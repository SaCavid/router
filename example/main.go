package main

import (
	"context"
	"github.com/SaCavid/router/models"
)

func init() {

}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event models.LambdaRequest) {

}
