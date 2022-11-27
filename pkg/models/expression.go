package models

import (
	"github.com/google/uuid"
	"time"
)

type ExpressionModel struct {
	Id         uuid.UUID `json:"id" gorm:"type:uuid"`
	Expression string    `json:"expression"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewExpressionModel(expression string) *ExpressionModel {
	return &ExpressionModel{
		Id:         uuid.New(),
		Expression: expression,
		CreatedAt:  time.Now(),
	}
}
