package controllers

import (
	"errors"
	"exval/pkg/expression"
	"exval/pkg/models"
	"exval/pkg/persistence"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ExpressionCreateRequest struct {
	Expression string `json:"expression" validate:"required"`
}

type ExpressionController struct {
	repository persistence.ExpressionRepository
}

func NewExpressionController(repository persistence.ExpressionRepository) *ExpressionController {
	return &ExpressionController{repository: repository}
}

// ExpressionList returns a list of all expressions
//
// @Summary      List all expressions
// @Produce      json
// @Tags         expression
// @Success      200  {object}  []models.ExpressionModel
// @Failure      500
// @Router       /expressions [get]
func (ec *ExpressionController) ExpressionList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	expressions, err := ec.repository.List(ctx)

	if err != nil {
		respondError(ctx, w, err)
		return
	}

	respondJson(ctx, w, http.StatusOK, expressions)
}

// ExpressionCreate creates a new expression
//
// @Summary      Create a new expression
// @Produce      json
// @Tags         expression
// @Accept       json
// @Param        expression     body    ExpressionCreateRequest     true        "Expression to create"
// @Success      200  {object}  models.ExpressionModel
// @Failure      500
// @Router       /expressions [post]
func (ec *ExpressionController) ExpressionCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	request, err := validateFromBody[ExpressionCreateRequest](r)
	if err != nil {
		respondError(ctx, w, err)
		return
	}

	_, err = expression.NewParser(request.Expression)
	if err != nil {
		respondError(ctx, w, err)
		return
	}

	model := models.NewExpressionModel(request.Expression)

	err = ec.repository.Create(ctx, model)
	if err != nil {
		respondError(ctx, w, err)
		return
	}

	respondJson(ctx, w, http.StatusCreated, model)
}

// ExpressionDeleteById deletes an expression by id
//
// @Summary      Delete an expression by id
// @Tags         expression
// @Param        id    path    string     true        "Expression id"
// @Success      200
// @Failure      500
// @Router       /expressions/{id} [delete]
func (ec *ExpressionController) ExpressionDeleteById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		respondError(ctx, w, errors.New("id must be an uuid4"))
		return
	}

	err = ec.repository.DeleteById(ctx, id)
	if err != nil {
		respondError(ctx, w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type EvaluateResponse struct {
	Result bool `json:"result"`
}

// ExpressionEvaluate evaluates an expression by id
//
// @Summary      Evaluate an expression by id
// @Produce      json
// @Tags         expression
// @Param        id    		path    string     				true        "Expression id"
// @Success      200  {object}  EvaluateResponse
// @Router       /evaluate/{id} [get]
func (ec *ExpressionController) ExpressionEvaluate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		respondError(ctx, w, errors.New("id must be an uuid4"))
		return
	}

	model, err := ec.repository.GetById(ctx, id)
	if err != nil {
		respondError(ctx, w, err)
		return
	}

	parser, err := expression.NewParser(model.Expression)
	if err != nil {
		respondError(ctx, w, err)
		return
	}

	symbolTable := map[string]int{}

	for key, value := range r.URL.Query() {
		atoi, err := strconv.Atoi(value[0])
		if err != nil {
			continue
		}
		symbolTable[key] = atoi
	}

	eval, err := parser.Eval(symbolTable)
	if err != nil {
		respondError(ctx, w, err)
		return
	}

	respondJson(ctx, w, http.StatusOK, EvaluateResponse{
		Result: eval,
	})
}
