//go:build wireinject
// +build wireinject

package di

import (
	"OnionPractice/app/infrastructure/database"
	"OnionPractice/app/usecase"
	"OnionPractice/db"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	db.GetDB,

	// repository
	database.NewTodoRepository,
	database.NewCategoryRepository,
)

// use case
func CreateTodoUseCase() (usecase.CreateTodoUseCase, error) {
	wire.Build(
		providerSet,
		usecase.NewCreateTodoUseCase,
	)
	return usecase.CreateTodoUseCase{}, nil
}

func GetTodoUseCase() (usecase.GetTodoUseCase, error) {
	wire.Build(
		providerSet,
		usecase.NewGetTodoUseCase,
	)
	return usecase.GetTodoUseCase{}, nil
}

func GetAllCategoriesUseCase() (usecase.GetAllCategoriesUseCase, error) {
	wire.Build(
		providerSet,
		usecase.NewGetAllCategoriesUseCase,
	)
	return usecase.GetAllCategoriesUseCase{}, nil
}

func CreateCategoryUseCase() (usecase.CreateCategoryUseCase, error) {
	wire.Build(
		providerSet,
		usecase.NewCreateCategoryUseCase,
	)
	return usecase.CreateCategoryUseCase{}, nil
}
