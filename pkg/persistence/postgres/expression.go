package postgres

import (
	"context"
	"exval/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExpressionRepositoryImpl struct {
	db *gorm.DB
}

func (e *ExpressionRepositoryImpl) GetById(ctx context.Context, id uuid.UUID) (*models.ExpressionModel, error) {
	db := e.db.WithContext(ctx)

	var data models.ExpressionModel
	err := db.Model(&models.ExpressionModel{}).Where("id = ?", id).First(&data).Error

	return &data, err
}

func NewExpressionRepositoryImpl(db *gorm.DB) *ExpressionRepositoryImpl {
	return &ExpressionRepositoryImpl{db: db}
}

func (e *ExpressionRepositoryImpl) List(ctx context.Context) ([]*models.ExpressionModel, error) {
	db := e.db.WithContext(ctx)

	var data []*models.ExpressionModel
	err := db.Model(&models.ExpressionModel{}).Find(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e *ExpressionRepositoryImpl) Create(ctx context.Context, model *models.ExpressionModel) error {
	db := e.db.WithContext(ctx)

	return db.Create(model).Error
}

func (e *ExpressionRepositoryImpl) DeleteById(ctx context.Context, id uuid.UUID) error {
	db := e.db.WithContext(ctx)

	return db.Delete(&models.ExpressionModel{Id: id}).Error
}
