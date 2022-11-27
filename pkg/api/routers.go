package api

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/http-swagger"
	"net/http"
)

func (e *ExvalApi) setupRoutes() {
	e.router.Use(mux.CORSMethodMiddleware(e.router))
	e.router.Use(recoveryMiddleware)

	e.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
	)).Methods(http.MethodGet)

	e.router.HandleFunc("/evaluate/{id}", e.expressionController.ExpressionEvaluate).Methods(http.MethodGet)
	e.router.HandleFunc("/expressions", e.expressionController.ExpressionList).Methods(http.MethodGet)
	e.router.HandleFunc("/expressions", e.expressionController.ExpressionCreate).Methods(http.MethodPost)
	e.router.HandleFunc("/expressions/{id}", e.expressionController.ExpressionDeleteById).Methods(http.MethodDelete)
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				log.Error(err)
				w.WriteHeader(http.StatusInternalServerError)
			}

		}()

		next.ServeHTTP(w, r)

	})
}
