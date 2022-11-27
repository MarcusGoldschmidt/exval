package postgres

import (
	"context"
	"exval/pkg/models"
	"exval/pkg/test"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestExpressionRepositoryImplCreate(t *testing.T) {
	// Setup
	container := test.SetupPostgresContainer(t)
	defer container.Close()

	ctx := context.Background()

	db := container.GormDb()

	err := AutoMigration(db)
	if err != nil {
		t.Fatal(err)
	}
	repo := NewExpressionRepositoryImpl(db)

	model := &models.ExpressionModel{
		Id:         uuid.New(),
		Expression: "A AND B",
		CreatedAt:  time.Now(),
	}

	// Act
	err = repo.Create(ctx, model)

	// Assert

	if err != nil {
		t.Fatal(err)
	}

	var result models.ExpressionModel

	err = db.Model(&models.ExpressionModel{}).
		Where(&models.ExpressionModel{
			Id: model.Id,
		}).
		First(&result).
		Error

	if err != nil {
		t.Fatal(err)
	}

	// assert result
	if result.Id != model.Id {
		t.Fatalf("result.Id != model.Id: %v != %v", result.Id, model.Id)
	}

	if result.Expression != model.Expression {
		t.Fatalf("result.Expression != model.Expression: %v != %v", result.Expression, model.Expression)
	}

	if result.CreatedAt.Equal(model.CreatedAt) {
		t.Fatalf("result.CreatedAt != model.CreatedAt: %v != %v", result.CreatedAt, model.CreatedAt)
	}
}

func TestExpressionRepositoryImplDelete(t *testing.T) {
	// Setup
	container := test.SetupPostgresContainer(t)
	defer container.Close()

	ctx := context.Background()

	db := container.GormDb()

	err := AutoMigration(db)
	if err != nil {
		t.Fatal(err)
	}
	repo := NewExpressionRepositoryImpl(db)

	model := &models.ExpressionModel{
		Id:         uuid.New(),
		Expression: "A AND B",
		CreatedAt:  time.Now(),
	}

	err = db.Create(model).Error
	if err != nil {
		t.Fatal(err)
	}

	// Act
	err = repo.DeleteById(ctx, model.Id)

	// Assert

	if err != nil {
		t.Fatal(err)
	}

	var result *models.ExpressionModel

	err = db.Model(&models.ExpressionModel{}).
		Where(&models.ExpressionModel{
			Id: model.Id,
		}).
		First(&result).
		Error

	if err != gorm.ErrRecordNotFound {
		t.Fatal(err)
	}
}

func TestExpressionRepositoryImplGet(t *testing.T) {
	// Setup
	container := test.SetupPostgresContainer(t)
	defer container.Close()

	ctx := context.Background()

	db := container.GormDb()

	err := AutoMigration(db)
	if err != nil {
		t.Fatal(err)
	}
	repo := NewExpressionRepositoryImpl(db)

	modelList := []*models.ExpressionModel{
		{
			Id:         uuid.New(),
			Expression: "A AND B",
			CreatedAt:  time.Now(),
		},
		{
			Id:         uuid.New(),
			Expression: "A AND B",
			CreatedAt:  time.Now(),
		},
		{
			Id:         uuid.New(),
			Expression: "A AND B",
			CreatedAt:  time.Now(),
		},
		{
			Id:         uuid.New(),
			Expression: "A AND B",
			CreatedAt:  time.Now(),
		},
		{
			Id:         uuid.New(),
			Expression: "A AND B",
			CreatedAt:  time.Now(),
		},
	}

	for _, model := range modelList {
		err = db.Create(model).Error
		if err != nil {
			t.Fatal(err)
		}
	}

	// Act
	list, err := repo.List(ctx)
	if err != nil {
		t.Fatal(err)
	}

	getById, err := repo.GetById(ctx, modelList[0].Id)
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	if len(list) != len(modelList) {
		t.Fatalf("len(list) != len(modelList): %v != %v", len(list), len(modelList))
	}

	if getById.Id != modelList[0].Id {
		t.Fatalf("getById.Id != modelList[0].Id: %v != %v", getById.Id, modelList[0].Id)
	}
}
