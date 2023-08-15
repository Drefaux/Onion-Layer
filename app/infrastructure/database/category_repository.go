package database

import (
	"OnionPractice/app/domain/model"
	"OnionPractice/app/domain/repository"
	"OnionPractice/app/infrastructure/database/dbmodel"
	"context"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func (r CategoryRepositoryImpl) GetAllCategories(ctx context.Context) ([]*model.Category, error) {
	var categories []dbmodel.Category
	if err := r.db.Find(&categories, "company_id = ?", model.LrmCompanyID).Error; err != nil {
		return nil, err
	}

	return r.dbModelToDomainModelPointerSlice(categories), nil
}

func (r CategoryRepositoryImpl) CreateCategory(ctx context.Context, category model.Category) (*model.Category, error) {
	m := r.domainModelToDBModel(category)
	if err := r.db.Create(m).Error; err != nil {
		return nil, err
	}

	return r.dbModelToDomainModelPointer(*m), nil
}

func (r CategoryRepositoryImpl) domainModelToDBModel(entity model.Category) *dbmodel.Category {
	return &dbmodel.Category{
		ID:        entity.ID(),
		CompanyID: entity.CompanyID(),
		Name:      entity.Name(),
	}
}

func (r CategoryRepositoryImpl) dbModelToDomainModel(category dbmodel.Category) model.Category {
	return model.RestoreCategory(
		category.ID,
		category.CompanyID,
		category.Name,
	)
}

func (r CategoryRepositoryImpl) dbModelToDomainModelPointerSlice(categories []dbmodel.Category) []*model.Category {
	domainCategories := make([]*model.Category, len(categories))

	for i, category := range categories {
		domainCategory := r.dbModelToDomainModel(category)
		domainCategories[i] = &domainCategory
	}

	return domainCategories
}

func (r CategoryRepositoryImpl) dbModelToDomainModelPointer(category dbmodel.Category) *model.Category {
	e := r.dbModelToDomainModel(category)

	return &e
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}
