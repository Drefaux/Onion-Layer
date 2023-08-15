package usecase

import (
	"OnionPractice/app/domain/model"
	"OnionPractice/app/domain/repository"
	"context"
	"fmt"
)

type GetAllCategoriesUseCase struct {
	categoryRepository repository.CategoryRepository
}

func NewGetAllCategoriesUseCase(categoryRepository repository.CategoryRepository) GetAllCategoriesUseCase {
	return GetAllCategoriesUseCase{categoryRepository: categoryRepository}
}

func (u GetAllCategoriesUseCase) GetAllCategoriesUseCase(ctx context.Context) ([]*model.Category, error) {
	categories, err := u.categoryRepository.GetAllCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to categoryRepository.Get: %w", err)
	}

	return categories, nil
}
