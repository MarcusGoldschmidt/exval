package persistence

import (
	"context"
	"exval/pkg/models"
	"github.com/google/uuid"
)

type ExpressionRepository interface {
	GetById(ctx context.Context, id uuid.UUID) (*models.ExpressionModel, error)
	List(ctx context.Context) ([]*models.ExpressionModel, error)
	Create(context.Context, *models.ExpressionModel) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
