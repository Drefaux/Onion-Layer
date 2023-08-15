package repository

import (
	"OnionPractice/app/domain/model"
	"context"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category model.Category) (*model.Category, error)
	GetAllCategories(ctx context.Context) ([]*model.Category, error)
}
