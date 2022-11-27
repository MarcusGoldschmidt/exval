package controllers

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var validate = validator.New()

func respondJson(ctx context.Context, w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("%s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(response)
	if err != nil {
		log.Errorf("%s", err.Error())
		return
	}
}

func toJsonReadCloser(payload interface{}) io.ReadCloser {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("%s", err.Error())
		return nil
	}
	return io.NopCloser(strings.NewReader(string(response)))
}

func validateFromBody[T any](r *http.Request) (*T, error) {
	var response T

	all, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(all, &response)
	if err != nil {
		return nil, err
	}

	err = validate.StructCtx(r.Context(), response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func respondError(ctx context.Context, w http.ResponseWriter, err error) {
	log.Errorf("%s", err.Error())
	respondJson(ctx, w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
}
