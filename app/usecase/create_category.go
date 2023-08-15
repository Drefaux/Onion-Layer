package usecase

import (
	"OnionPractice/app/domain/model"
	"OnionPractice/app/domain/repository"
	"context"
	"fmt"
)

type CreateCategoryUseCase struct {
	categoryRepository repository.CategoryRepository
}

func NewCreateCategoryUseCase(categoryRepository repository.CategoryRepository) CreateCategoryUseCase {
	return CreateCategoryUseCase{categoryRepository: categoryRepository}
}

type CreateCategoryUseCaseInput struct {
	ID        int
	CompanyID int
	Name      string
}

type CreateCategoryUseCaseOutput struct {
	Category *model.Category
}

func (u CreateCategoryUseCase) CreateCategoryUseCase(ctx context.Context, input CreateCategoryUseCaseInput) (*CreateCategoryUseCaseOutput, error) {
	category := model.NewCategory(input.ID, input.CompanyID, input.Name)

	createCategory, err := u.categoryRepository.CreateCategory(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("failed to categoryRepository.Create: %w", err)
	}

	return &CreateCategoryUseCaseOutput{Category: createCategory}, nil
}
