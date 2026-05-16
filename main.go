package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sraisl/faker-api/faker"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

type RandomNameOutput struct {
	Body struct {
		Name string `json:"name" example:"John Doe" doc:"Randomly generated name"`
	}
}

type healthStatus struct {
	Body struct {
		Status string `json:"status" example:"OK" doc:"Health status"`
	}
}

func newRouter() http.Handler {
	// Create a new router & API.
	router := chi.NewMux()
	api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

	// Register GET /greeting/{name} handler.
	huma.Get(api, "/greeting/{name}", func(ctx context.Context, input *struct {
		Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	}) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})

	huma.Get(api, "/random-name", func(ctx context.Context, input *struct{}) (*RandomNameOutput, error) {
		resp := &RandomNameOutput{}
		resp.Body.Name = faker.FakeName()
		return resp, nil
	})

	huma.Get(api, "/health", func(ctx context.Context, input *struct{}) (*healthStatus, error) {
		resp := &healthStatus{}
		resp.Body.Status = "OK"
		return resp, nil
	})

	return router
}

func main() {
	router := newRouter()

	// Start the server!
	http.ListenAndServe("0.0.0.0:8888", router)
}
