package api

import (
	"context"
	"exval/pkg"
	"exval/pkg/controllers"
	"github.com/gorilla/mux"
	"github.com/swaggo/swag"
	"net/http"
	"strconv"
)

type ExvalApi struct {
	options              *pkg.ExvalOptions
	expressionController *controllers.ExpressionController
	server               *http.Server
	router               *mux.Router
	swagger              *swag.Spec
}

func NewExvalApi(options *pkg.ExvalOptions, expressionController *controllers.ExpressionController, swagger *swag.Spec) *ExvalApi {
	return &ExvalApi{
		options:              options,
		expressionController: expressionController,
		swagger:              swagger,
	}
}

func (e *ExvalApi) Start(ctx context.Context) error {
	e.router = mux.NewRouter()

	e.setupRoutes()

	e.server = &http.Server{
		Handler: e.router,
		Addr:    e.options.Address + ":" + strconv.Itoa(e.options.Port),
	}

	return e.server.ListenAndServe()
}

func (e *ExvalApi) Shutdown(ctx context.Context) error {
	return e.server.Shutdown(ctx)
}
