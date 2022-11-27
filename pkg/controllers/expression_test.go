package controllers

import (
	"context"
	"exval/pkg/models"
	postgresPersistence "exval/pkg/persistence/postgres"
	"exval/pkg/test"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	testHttp "github.com/stretchr/testify/http"
	"net/http"
	"testing"
)

func TestExpressionList(t *testing.T) {
	container := test.SetupPostgresContainer(t)
	defer container.Close()

	db := container.GormDb()
	err := postgresPersistence.AutoMigration(db)
	assert.NoError(t, err)

	controller := NewExpressionController(postgresPersistence.NewExpressionRepositoryImpl(db))

	response := &testHttp.TestResponseWriter{}

	controller.ExpressionList(response, &http.Request{})

	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestExpressionCreate(t *testing.T) {
	container := test.SetupPostgresContainer(t)
	defer container.Close()

	db := container.GormDb()
	err := postgresPersistence.AutoMigration(db)
	assert.NoError(t, err)

	controller := NewExpressionController(postgresPersistence.NewExpressionRepositoryImpl(db))

	response := &testHttp.TestResponseWriter{}

	controller.ExpressionCreate(response, &http.Request{
		Body: toJsonReadCloser(ExpressionCreateRequest{
			Expression: "A AND B",
		}),
	})

	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func TestExpressionDelete(t *testing.T) {
	// Setup
	container := test.SetupPostgresContainer(t)
	defer container.Close()

	ctx := context.Background()
	db := container.GormDb()
	err := postgresPersistence.AutoMigration(db)
	assert.NoError(t, err)

	controller := NewExpressionController(postgresPersistence.NewExpressionRepositoryImpl(db))

	err = controller.repository.Create(ctx, models.NewExpressionModel("A AND B"))
	assert.NoError(t, err)

	response := &testHttp.TestResponseWriter{}
	assert.NoError(t, err)

	r := mux.NewRouter()
	r.HandleFunc("/expressions/{id}", controller.ExpressionDeleteById)

	req, err := http.NewRequest("DELETE", "/expressions/"+uuid.New().String(), nil)
	assert.NoError(t, err)

	// Act

	r.ServeHTTP(response, req)

	// Assert

	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestExpressionEvaluate(t *testing.T) {
	// Setup
	container := test.SetupPostgresContainer(t)
	defer container.Close()

	ctx := context.Background()
	db := container.GormDb()
	err := postgresPersistence.AutoMigration(db)
	assert.NoError(t, err)

	controller := NewExpressionController(postgresPersistence.NewExpressionRepositoryImpl(db))

	model := models.NewExpressionModel("A AND B")

	err = controller.repository.Create(ctx, model)
	assert.NoError(t, err)

	response := &testHttp.TestResponseWriter{}
	assert.NoError(t, err)

	r := mux.NewRouter()
	r.HandleFunc("/evaluate/{id}", controller.ExpressionEvaluate)

	req, err := http.NewRequest("GET", "/evaluate/"+model.Id.String()+"?A=1&B=1", nil)
	assert.NoError(t, err)

	// Act

	r.ServeHTTP(response, req)

	// Assert

	assert.Equal(t, http.StatusOK, response.StatusCode)
}
